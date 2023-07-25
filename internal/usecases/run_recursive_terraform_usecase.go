package usecases

import (
	"github.com/atorrescogollo/terraform-cascade/internal/orchestration"
)

type RunRecursiveTerraformUseCase interface {
	Execute(execDir string, terraformArgs []string) error
}

/*
* Mock
 */
type MockRunRecursiveTerraformUseCase struct {
	Executions []struct {
		ExecDir       string
		TerraformArgs []string
	}
}

func NewMockRunRecursiveTerraformUseCase() *MockRunRecursiveTerraformUseCase {
	return &MockRunRecursiveTerraformUseCase{
		Executions: make([]struct {
			ExecDir       string
			TerraformArgs []string
		}, 0),
	}
}
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
type RunRecursiveTerraformUseCaseImpl struct {
	projectExecutor       orchestration.ProjectExecutor
	projectObjectResolver orchestration.ProjectOrderResolver
}

func NewRunRecursiveTerraformUseCaseImpl(projectExecutor orchestration.ProjectExecutor, projectObjectResolver orchestration.ProjectOrderResolver) *RunRecursiveTerraformUseCaseImpl {
	return &RunRecursiveTerraformUseCaseImpl{
		projectExecutor,
		projectObjectResolver,
	}
}

func (u RunRecursiveTerraformUseCaseImpl) Execute(execDir string, terraformArgs []string) error {
	projects, err := u.projectObjectResolver.Resolve(execDir)
	if err != nil {
		return err
	}
	return u.projectExecutor.Execute(projects, terraformArgs)
}
