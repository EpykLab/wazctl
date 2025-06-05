/*
Copyright © 2025 stllr

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
	"log"

	"github.com/EpykLab/wazctl/internal/files"
	"github.com/EpykLab/wazctl/internal/templates/rules"
	"github.com/spf13/cobra"
)

// ruleCmd represents the rule command
var ruleCmd = &cobra.Command{
	Use:   "rule",
	Short: "create wazctl rule file",
	Run: func(cmd *cobra.Command, args []string) {
		name := cmd.Flag("name").Value.String()

		err := files.FileCreateWithSpecifiedNameAndContent(
			fmt.Sprintf("%s.yaml", name),
			rules.ScaffoldFromTempl())

		if err != nil {
			log.Println(err)
		}

	},
}

func init() {
	rootCmd.AddCommand(ruleCmd)

	ruleCmd.Flags().StringP("name", "n", "", "name of new rule file")

	ruleCmd.MarkFlagRequired("name")
}
