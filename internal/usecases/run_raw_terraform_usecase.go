package usecases

import (
	"github.com/atorrescogollo/terraform-cascade/internal/terraform"
)

type RunRawTerraformUseCase struct{}

func NewRunRawTerraformUseCase() *RunRawTerraformUseCase {
	return &RunRawTerraformUseCase{}
}

func (u RunRawTerraformUseCase) Execute(execDir string, terraformArgs []string) error {
	return terraform.TerraformExecWithOS(execDir, terraformArgs)
}
