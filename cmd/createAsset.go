package cmd

import (
	"fmt"

	"github.com/schaeffler/tractus-x-cli-tool/utils"
	"github.com/spf13/cobra"
)

// getCmd represents the get command
var createAssetCmd = &cobra.Command{
	Use:   "createAsset",
	Short: "Command that creates asset with HTTP call",
	Long:  `Command that creates an asset. Invokation of this command requires --assetId argument`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("createAsset called")
		assetId, _ := cmd.Flags().GetString("assetId")
		if assetId != "" {
			CreateAsset(assetId)
		} else {
			panic("assetId is required")
		}
	},
}

func init() {
	rootCmd.AddCommand(createAssetCmd)
	createAssetCmd.Flags().String("assetId", "", "Asset ID that's being created")
}

func CreateAsset(assetId string) {
	assetsUrl := "http://dataprovider-controlplane.tx.test/management/v3/assets"

	createAssetDto := `{
    "@context": {
      "@vocab": "https://w3id.org/edc/v0.0.1/ns/",
      "edc": "https://w3id.org/edc/v0.0.1/ns/",
      "tx": "https://w3id.org/tractusx/v0.0.1/ns/",
      "tx-auth": "https://w3id.org/tractusx/auth/",
      "cx-policy": "https://w3id.org/catenax/policy/",
      "odrl": "http://www.w3.org/ns/odrl/2/"
     },
    "@id": "%s",
    "properties": {
      "description": "Product EDC Demo Asset"
    },
    "dataAddress": {
      "@type": "DataAddress",
      "type": "HttpData",
      "proxyPath": "true",
      "proxyMethod": "true",
      "proxyQueryParams": "true",
      "proxyBody": "true",
      "baseUrl": "http://dataprovider-submodelserver.tx.test"
    }
  }`
	createAssetDto = fmt.Sprintf(createAssetDto, assetId)
	d := utils.SendPostRequest([]byte(createAssetDto), assetsUrl, "TEST2")
	fmt.Println(d)
}
