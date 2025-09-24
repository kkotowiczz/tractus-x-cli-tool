package cmd

import (
	"fmt"

	"github.com/schaeffler/tractus-x-cli-tool/utils"
	"github.com/spf13/cobra"
)

// createPolicyCmd represents the createPolicy command
var createPolicyCmd = &cobra.Command{
	Use:   "createPolicy",
	Short: "Command that creates policy with HTTP call",
	Long:  `Command that creates an policy. Invokation of this command requires --policyId argument`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("createPolicy called")
		policyId, _ := cmd.Flags().GetString("policyId")
		if policyId != "" {
			CreatePolicy(policyId)
		} else {
			panic("policyId is required")
		}
	},
}

func init() {
	rootCmd.AddCommand(createPolicyCmd)
	createPolicyCmd.Flags().String("policyId", "", "Policy ID that's being created")
}

func CreatePolicy(policyId string) {
	policyDefinitionUrl := "http://dataprovider-controlplane.tx.test/management/v3/policydefinitions"

	createPolicyDto := `
	{
  "@context": {
      "@vocab": "https://w3id.org/edc/v0.0.1/ns/"
    },
  "@type": "PolicyDefinition",
  "@id": "%s",
  "policy": {
    "@context": [
        "http://www.w3.org/ns/odrl.jsonld",
        "https://w3id.org/catenax/2025/9/policy/context.jsonld"
    ],
    "@type": "Set",
	"permission": [
      {
        "action": "access",
        "constraint": [
        ]
      }
    ]
  }
}`
	createPolicyDto = fmt.Sprintf(createPolicyDto, policyId)

	utils.SendPostRequest([]byte(createPolicyDto), policyDefinitionUrl)
}
