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
package cmd

import (
	"fmt"
	"os"
	"strings"

	"github.com/atorrescogollo/terraform-cascade/internal/shared/utils"
	"github.com/atorrescogollo/terraform-cascade/internal/terraform"
	"github.com/spf13/cobra"
)

// terraformCmd represents the terraform command
var terraformCmd = &cobra.Command{
	Use:   "terraform",
	Short: "Terraform commands",
	Long:  `Terraform commands through the cascade cli.`,
	Run: func(cmd *cobra.Command, args []string) {
		cmd.Flags().Parse(args)
		help := false
		if len(args) == 0 ||
			(len(args) == 1 && (args[0] == "-h" ||
				args[0] == "--help" ||
				args[0] == "-help" ||
				args[0] == "help")) {
			help = true
		}

		terraformArgs := utils.ExtractUnknownArgs(cmd.Flags(), args)
		if !help {
			/*
			* Help command is actually there since cobra adds it behind the scenes.
			* We add it when it's not the first argument so that it executes
			* help as a terraform command instead of the cascade help command.
			*
			* For example:
			*
			*   cascade terraform --help        # Executes cascade help + terraform help
			*
			*   cascade terraform plan --help   # Executes terraform-plan help
			*
			 */
			trailingHelp, _ := cmd.Flags().GetBool("help")
			if trailingHelp {
				terraformArgs = append(terraformArgs, "--help")
			}
		}
		showUsage := false
		for _, tfArg := range terraformArgs {
			if strings.HasPrefix(tfArg, "--cascade-") {
				// This is not a terraform flag. We need to handle
				// it here since UnknownFlags is whitelisted
				fmt.Println("Error: unknown flag:", tfArg)
				showUsage = true
				continue
			}
		}
		if help || showUsage {
			cmd.Usage()
			return
		}
		exitErr := terraform.TerraformExecWithOS(terraformArgs)
		if exitErr != nil {
			os.Exit(exitErr.ExitCode())
			return
		}
	},
}

func init() {
	rootCmd.AddCommand(terraformCmd)
	terraformCmd.FParseErrWhitelist.UnknownFlags = true
	terraformCmd.SilenceErrors = true
	terraformCmd.SilenceUsage = true
	terraformCmd.DisableFlagParsing = true

	terraformCmd.SetUsageTemplate(`
Usage:
{{.CommandPath}} [flags] [terraform args]

{{- if .HasAvailableLocalFlags}}

Flags:
{{.LocalFlags.FlagUsages | trimTrailingWhitespaces}}
{{- end}}

{{- if .HasAvailableInheritedFlags}}

Global Flags:
{{.InheritedFlags.FlagUsages | trimTrailingWhitespaces}}
{{- end}}


Terraform Args:
`)
	defaultUsageFunc := terraformCmd.UsageFunc()
	terraformCmd.SetUsageFunc(func(cmd *cobra.Command) error {
		terraformUsage := terraform.TerraformUsage()
		err := defaultUsageFunc(cmd)
		if err != nil {
			return err
		}

		fmt.Println()
		fmt.Println(terraformUsage)
		return nil
	})

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// terraformCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	terraformCmd.Flags().Bool("cascade-recursive", false, "Execute terraform projects recursively in order")
}
