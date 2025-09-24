/*
Copyright © 2025 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"encoding/json"
	"fmt"

	"os/exec"

	"github.com/schaeffler/tractus-x-cli-tool/utils"
	"github.com/spf13/cobra"
)

// createContractNegotiationCmd represents the createContractNegotiation command
var createContractNegotiationCmd = &cobra.Command{
	Use:   "createContractNegotiation",
	Short: "Command that creates contact negotiation with HTTP call",
	Long:  `Creates contract negotiation by requesting catalog first and embedding the offer id into contract initiation http call`,
	Run: func(cmd *cobra.Command, args []string) {
		url := "http://dataconsumer-1-controlplane.tx.test/management/v3/catalog/request"
		catologRequest := `{
			"@context": {
			"@vocab": "https://w3id.org/edc/v0.0.1/ns/"
			},
			"@type": "CatalogRequest",
			"counterPartyAddress": "http://dataprovider-controlplane.tx.test/api/v1/dsp",
			"counterPartyId": "BPNL00000003AYRE",
			"protocol": "dataspace-protocol-http",
			"querySpec": {
			"filterExpression": {
				"operandLeft": "https://w3id.org/edc/v0.0.1/ns/id",
				"operator": "=",
				"operandRight": "200"
			},
			"offset": 0,
			"limit": 50
			}
		}`
		responseBody := utils.SendPostRequest([]byte(catologRequest), url)

		fmt.Println("createContractNegotiation called")
		catalogResponse := utils.CatalogRequestResponse{}
		var plainJson interface{}
		json.Unmarshal([]byte(responseBody), &plainJson)
		m := plainJson.(map[string]interface{})
		dataset, _ := json.Marshal(m["dcat:dataset"])

		json.Unmarshal(dataset, &catalogResponse)

		out := RunCurl(catalogResponse.ID, catalogResponse.OdrlHasPolicy.ID)
		fmt.Println(out)
	},
}

func init() {
	rootCmd.AddCommand(createContractNegotiationCmd)
	createContractDefinitionCmd.Flags().String("assetId", "", "asset ID that's being part of contract negotiation")
}

func RunCurl(assetId string, offerId string) string {
	curl := `  curl -L -X POST 'http://dataconsumer-1-controlplane.tx.test/management/v3/contractnegotiations'   -H 'Content-Type: application/json'   -H 'X-Api-Key: TEST2'   --data-raw '{
  "@context": {
    "@vocab": "https://w3id.org/edc/v0.0.1/ns/"
  },
  "@type": "ContractRequest",
  "counterPartyAddress": "http://dataprovider-controlplane.tx.test/api/v1/dsp",
  "protocol": "dataspace-protocol-http",
  "policy": {
    "@context": [
      "http://www.w3.org/ns/odrl.jsonld",
      "https://w3id.org/tractusx/edc/v0.0.1"
    ],
    "assigner": "BPNL00000003AYRE",
    "target": "%s",
    "@id": "%s",
    "@type": "odrl:Offer",
      "odrl:permission": {
        "odrl:action": {
          "@id": "use"
        }
      },
      "odrl:prohibition": [],
      "odrl:obligation": []
  },
  "callbackAddresses": []
}' | jq`
	curl = fmt.Sprintf(curl, assetId, offerId)
	out, err := exec.Command("bash", "-c", curl).Output()
	if err != nil {
		panic("some error found")
	}
	return string(out)
}
