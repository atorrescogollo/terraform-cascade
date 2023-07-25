package usecases

import (
	"github.com/atorrescogollo/terraform-cascade/internal/terraform"
)

type RunRawTerraformUseCase interface {
	Execute(execDir string, terraformArgs []string) error
}

/*
* Mock
 */
type MockRunRawTerraformUseCase struct {
	Executions []struct {
		ExecDir       string
		TerraformArgs []string
	}
}

func NewMockRunRawTerraformUseCase() *MockRunRawTerraformUseCase {
	return &MockRunRawTerraformUseCase{
		Executions: make([]struct {
			ExecDir       string
			TerraformArgs []string
		}, 0),
	}
}
func (u *MockRunRawTerraformUseCase) Execute(execDir string, terraformArgs []string) error {
	u.Executions = append(u.Executions, struct {
		ExecDir       string
		TerraformArgs []string
	}{execDir, terraformArgs})
	return nil
}

/*
* Implementation
 */
type RunRawTerraformUseCaseImpl struct{}

func NewRunRawTerraformUseCaseImpl() *RunRawTerraformUseCaseImpl {
	return &RunRawTerraformUseCaseImpl{}
}

func (u RunRawTerraformUseCaseImpl) Execute(execDir string, terraformArgs []string) error {
	return terraform.TerraformExecWithOS(execDir, terraformArgs)
}
