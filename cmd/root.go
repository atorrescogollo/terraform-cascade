package cmd

/*
Copyright © 2023 Álvaro Torres Cogollo <atorrescogollo@gmail.com>

This program is free software: you can redistribute it and/or modify
it under the terms of the GNU General Public License as published by
the Free Software Foundation, either version 3 of the License, or
(at your option) any later version.

This program is distributed in the hope that it will be useful,
but WITHOUT ANY WARRANTY; without even the implied warranty of
MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
GNU General Public License for more details.

You should have received a copy of the GNU General Public License
along with this program. If not, see <http://www.gnu.org/licenses/>.
*/

import (
	"os"
	"strings"

	"github.com/spf13/cobra"
	"golang.org/x/exp/slices"
)

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "terraform-cascade",
	Short: "An opinionated terraform project orchestrator",
	Long: `
Terraform Cascade is a terraform-like tool that allows you to manage multiple terraform projects.

It's made to be fully compatible with terraform, so you can use it as a drop-in replacement. However, it requires the terraform binary to be available in the PATH.

== Design ==
It works with a very opinionated design:

* Every project is inside a deep directory structure.
* To define a project, you only need to place a backend.tf file in that directory.
    * In each layer, will be executed in the following order:
    * Current directory (only when it has a backend.tf file)
    * Whole base directory (with its layer)
    * Other directories (with its layer)
`,
	// Uncomment the following line if your bare application
	// has an action associated with it:
	//Run: func(cmd *cobra.Command, args []string) {},
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	if len(os.Args) > 1 && !strings.HasPrefix(os.Args[1], "-") && !slices.Contains([]string{
		"terraform",
		"cascade",
		"help",
		"completion",
	}, os.Args[1]) {
		// Hijack as a terraform command
		os.Args = append([]string{os.Args[0], "terraform"}, os.Args[1:]...)
	}

	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {
}
