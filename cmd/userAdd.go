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

		client := actions.WazctlClientFactory()
		printers.PrintJsonFormattedOrError(client.CreateNewUserInWazuhManager(&actions.CreateNewWazuhUserOptions{
			Username: username,
			Password: password,
		}))
	},
}

func init() {
	useraddCmd.Flags().StringP("username", "u", "", "username of new user to create in wazuh")
	useraddCmd.Flags().StringP("password", "p", "", "password of new user to create in wazuh")

	useraddCmd.MarkFlagRequired("username")
	useraddCmd.MarkFlagRequired("password")
}
