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
	"log"

	"github.com/EpykLab/wazctl/internal/instance/local/docker"
	"github.com/spf13/cobra"
)

// dockerCmd represents the docker command
var dockerCmd = &cobra.Command{
	Use:   "docker",
	Short: "create a local wazuh instance in docker",
	Long:  ``,
	Run: func(cmd *cobra.Command, args []string) {
		start := cmd.Flag("start").Changed
		stop := cmd.Flag("stop").Changed
		clean := cmd.Flag("clean").Changed
		instance, err := docker.NewWazuhDockerManager()
		if err != nil {
			log.Println(err)
		}

		if start {
			if err := instance.Start(); err != nil {
				log.Println(err)
			}
		}
		if stop {
			if err := instance.Stop(); err != nil {
				log.Println(err)
			}
		}
		if clean {
			if err := instance.Clean(); err != nil {
				log.Println(err)
			}
		}
	},
}

func init() {
	dockerCmd.Flags().Bool("start", false, "spin up wazuh instance")
	dockerCmd.Flags().Bool("stop", false, "stop wazuh instance")
	dockerCmd.Flags().Bool("clean", false, "remove wazuh instance")
}
