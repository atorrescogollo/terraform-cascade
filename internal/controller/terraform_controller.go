package controller

import (
	"fmt"
	"os"

	"github.com/atorrescogollo/terraform-cascade/internal/terraform"
	"github.com/atorrescogollo/terraform-cascade/internal/usecases"
	"github.com/spf13/cobra"
)

type TerraformController struct {
	RunRawTerraformUseCase       usecases.RunRawTerraformUseCase
	RunRecursiveTerraformUseCase usecases.RunRecursiveTerraformUseCase
}

func NewTerraformController(runRawTerraformUseCase usecases.RunRawTerraformUseCase, runRecursiveTerraformUseCase usecases.RunRecursiveTerraformUseCase) *TerraformController {
	return &TerraformController{
		RunRawTerraformUseCase:       runRawTerraformUseCase,
		RunRecursiveTerraformUseCase: runRecursiveTerraformUseCase,
	}
}

func (c TerraformController) Handle(recursive bool, tfargs []string) error {
	cwd, _ := os.Getwd()
	if !recursive {
		/*
		* Simply run terraform in the current directory
		 */
		return c.RunRawTerraformUseCase.Execute(cwd, tfargs)
	}
	return c.RunRecursiveTerraformUseCase.Execute(cwd, tfargs)
}

func (c TerraformController) Usage(cmd *cobra.Command) error {
	terraformUsage := terraform.TerraformUsage()
	err := cmd.UsageFunc()(cmd)
	if err != nil {
		return err
	}

	fmt.Println()
	fmt.Println(terraformUsage)
	return nil
}
