package cmd

import (
	"fmt"

	"github.com/schaeffler/tractus-x-cli-tool/utils"
	"github.com/spf13/cobra"
)

// createContractDefinitionCmd represents the createContractDefinition command
var createContractDefinitionCmd = &cobra.Command{
	Use:   "createContractDefinition",
	Short: "Command that creates contact definition with HTTP call",
	Long:  `Command that creates an asset. Invokation of this command requires --contractId, --accessPolicyId, --contractPolicyId, --criterionId argument`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("createContractDefinition called")
		contractId, _ := cmd.Flags().GetString("contractId")
		accessPolicyId, _ := cmd.Flags().GetString("accessPolicyId")
		contractPolicyId, _ := cmd.Flags().GetString("contractPolicyId")
		criterionId, _ := cmd.Flags().GetString("criterionId")

		if contractId != "" && accessPolicyId != "" && contractPolicyId != "" && criterionId != "" {
			CreateContractDefinition(contractId, accessPolicyId, contractPolicyId, criterionId)
		} else {
			panic("One or more of following arguments wasn't provided [contractId, accessPolicyId, contractPolicyId, criterionId]")
		}
	},
}

func init() {
	rootCmd.AddCommand(createContractDefinitionCmd)
	createContractDefinitionCmd.Flags().String("contractId", "", "Contract ID that's being created")
	createContractDefinitionCmd.Flags().String("accessPolicyId", "", "Access policy id to link contract with")
	createContractDefinitionCmd.Flags().String("contractPolicyId", "", "Contract policy ID to link contract with")
	createContractDefinitionCmd.Flags().String("criterionId", "", "Criterion ID to filter")
}

func CreateContractDefinition(contractId string, accessPolicyId string, contractPolicyId string, criterionId string) {
	createContractUrl := "http://dataprovider-controlplane.tx.test/management/v3/contractdefinitions"

	createContractDto := `
	{
    "@context": {},
    "@id": "%s",
    "@type": "ContractDefinition",
    "accessPolicyId": "%s",
    "contractPolicyId": "%s",
    "assetsSelector": {
      "@type": "CriterionDto",
      "operandLeft": "https://w3id.org/edc/v0.0.1/ns/id",
      "operator": "=",
      "operandRight": "%s"
    }
  }`

	createContractDto = fmt.Sprintf(createContractDto, contractId, accessPolicyId, contractPolicyId, criterionId)

	utils.SendPostRequest([]byte(createContractDto), createContractUrl, "TEST2")
}
