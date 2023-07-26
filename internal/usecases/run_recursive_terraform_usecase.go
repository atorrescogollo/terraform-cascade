package usecases

import (
	"github.com/atorrescogollo/terraform-cascade/internal/orchestration"
)

// RunRecursiveTerraformUseCase runs terraform recursively in a given directory
// in the correct order
type RunRecursiveTerraformUseCase interface {
	Execute(execDir string, terraformArgs []string) error
}

/*
* Mock
 */

// MockRunRecursiveTerraformUseCase is a mock implementation of RunRecursiveTerraformUseCase
type MockRunRecursiveTerraformUseCase struct {
	Executions []struct {
		ExecDir       string
		TerraformArgs []string
	}
}

// NewMockRunRecursiveTerraformUseCase MockRunRecursiveTerraformUseCase constructor
func NewMockRunRecursiveTerraformUseCase() *MockRunRecursiveTerraformUseCase {
	return &MockRunRecursiveTerraformUseCase{
		Executions: make([]struct {
			ExecDir       string
			TerraformArgs []string
		}, 0),
	}
}

// Execute executes terraform in a given directory
func (u *MockRunRecursiveTerraformUseCase) Execute(execDir string, terraformArgs []string) error {
	u.Executions = append(u.Executions, struct {
		ExecDir       string
		TerraformArgs []string
	}{execDir, terraformArgs})
	return nil
}

/*
* Implementation
 */

// RunRecursiveTerraformUseCaseImpl is the implementation of RunRecursiveTerraformUseCase
type RunRecursiveTerraformUseCaseImpl struct {
	projectExecutor       orchestration.ProjectExecutor
	projectObjectResolver orchestration.ProjectOrderResolver
}

// NewRunRecursiveTerraformUseCaseImpl RunRecursiveTerraformUseCaseImpl constructor
func NewRunRecursiveTerraformUseCaseImpl(projectExecutor orchestration.ProjectExecutor, projectObjectResolver orchestration.ProjectOrderResolver) *RunRecursiveTerraformUseCaseImpl {
	return &RunRecursiveTerraformUseCaseImpl{
		projectExecutor,
		projectObjectResolver,
	}
}

// Execute executes terraform in a given directory
func (u RunRecursiveTerraformUseCaseImpl) Execute(execDir string, terraformArgs []string) error {
	projects, err := u.projectObjectResolver.Resolve(execDir)
	if err != nil {
		return err
	}
	return u.projectExecutor.Execute(projects, terraformArgs)
}
