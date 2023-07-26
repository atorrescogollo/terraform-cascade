package usecases

import (
	"github.com/atorrescogollo/terraform-cascade/internal/terraform"
)

// RunRawTerraformUseCase runs simple terraform command in a given directory
type RunRawTerraformUseCase interface {
	Execute(execDir string, terraformArgs []string) error
}

/*
* Mocks
 */

// MockRunRawTerraformUseCase is a mock implementation of RunRawTerraformUseCase
type MockRunRawTerraformUseCase struct {
	Executions []struct {
		ExecDir       string
		TerraformArgs []string
	}
}

// NewMockRunRawTerraformUseCase MockRunRawTerraformUseCase constructor
func NewMockRunRawTerraformUseCase() *MockRunRawTerraformUseCase {
	return &MockRunRawTerraformUseCase{
		Executions: make([]struct {
			ExecDir       string
			TerraformArgs []string
		}, 0),
	}
}

// Execute executes terraform in a given directory
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

// RunRawTerraformUseCaseImpl is the implementation of RunRawTerraformUseCase
type RunRawTerraformUseCaseImpl struct{}

// NewRunRawTerraformUseCaseImpl RunRawTerraformUseCaseImpl constructor
func NewRunRawTerraformUseCaseImpl() *RunRawTerraformUseCaseImpl {
	return &RunRawTerraformUseCaseImpl{}
}

// Execute executes terraform in a given directory
func (u RunRawTerraformUseCaseImpl) Execute(execDir string, terraformArgs []string) error {
	return terraform.ExecWithOS(execDir, terraformArgs)
}
