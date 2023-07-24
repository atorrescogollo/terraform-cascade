package controller

import (
	"fmt"

	"github.com/atorrescogollo/terraform-cascade/internal/usecases"
)

type CascadeController struct {
	RunRecursiveTerraformUseCase usecases.RunRecursiveTerraformUseCase
}

func NewCascadeController(runRecursiveTerraformUseCase usecases.RunRecursiveTerraformUseCase) *CascadeController {
	return &CascadeController{
		runRecursiveTerraformUseCase,
	}
}

func (c CascadeController) HandleCascade() error {
	// TODO: Implement cascade logic
	//cwd, _ := os.Getwd()
	//tfargs := []string{"init"}
	//chdir := cwd + "/samples/basic"
	//return c.RunRecursiveTerraformUseCase.Execute(chdir, tfargs)
	return fmt.Errorf("not implemented")
}
