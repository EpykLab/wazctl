/*
Copyright © 2025 EpykLab

Permission is hereby granted, free of charge, to any person obtaining a copy
of this software and associated documentation files (the "Software"), to deal
in the Software without restriction, including without limitation the rights
to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
copies of the Software, and to permit persons to whom the Software is
furnished to do so, subject to the following conditions:

The above copyright notice and this permission notice shall be included in
all copies or substantial portions of the Software.

THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN
THE SOFTWARE.
*/
package cmd

import (
	"fmt"
	"os"

	"github.com/EpykLab/wazctl/internal/printers"
	"github.com/EpykLab/wazctl/pkg/actions"
	"github.com/spf13/cobra"
)

// useraddCmd represents the useradd command
var useraddCmd = &cobra.Command{
	Use:   "add",
	Short: "Create a new user in wazuh",
	Run: func(cmd *cobra.Command, args []string) {
		username := cmd.Flag("username").Value.String()
		password := cmd.Flag("password").Value.String()
		components := cmd.Flag("component").Value.String()
		role := cmd.Flag("role").Value.String()

		switch components {
		case "wazuh":
			client := actions.WazctlClientFactory()
			printers.PrintJsonFormattedOrError(client.CreateNewUserInWazuhManager(&actions.CreateNewWazuhUserOptions{
				Username: username,
				Password: password,
			}))
		case "indexer":
			if role == "" {
				fmt.Printf("no role provided. use the [-r --role] flag to define the role the new user should have")
				os.Exit(1)
			}

			client := actions.IndexerClientFactory()
			printers.PrintJsonFormattedOrError(client.CreateNewUserInOSIndexer(&actions.CreateNewIndexerUserOptions{
				CreateIndexerUserPayload: actions.CreateIndexerUserPayload{
					Password: password,
				},
				IndexerRoleMappingPayload: actions.IndexerRoleMappingPayload{
					// TODO: will want to handle multiple users being passed in
					Users:        []string{username},
					BackendRoles: []string{role},
				},
			}))
		default:
			fmt.Print("component option not recognized. Must be one of [wazuh, indexer]")
			os.Exit(1)
		}
	},
}

func init() {
	// Configure flags
	useraddCmd.Flags().StringP("username", "u", "", "username of new user to create in wazuh")
	useraddCmd.Flags().StringP("password", "p", "", "password of new user to create in wazuh")
	useraddCmd.Flags().StringP("component", "c", "", "create a user in either the wazuh or indexer components")
	useraddCmd.Flags().StringP("role", "r", "", "role to assign user in the indexer")

	// set required flags
	useraddCmd.MarkFlagRequired("username")
	useraddCmd.MarkFlagRequired("password")
	useraddCmd.MarkFlagRequired("component")
}
