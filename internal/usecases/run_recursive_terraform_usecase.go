package usecases

import (
	"github.com/atorrescogollo/terraform-cascade/internal/orchestration"
)

type RunRecursiveTerraformUseCase struct {
	projectExecutor       orchestration.ProjectExecutor
	projectObjectResolver orchestration.ProjectOrderResolver
}

func NewRunRecursiveTerraformUseCase(projectExecutor orchestration.ProjectExecutor, projectObjectResolver orchestration.ProjectOrderResolver) *RunRecursiveTerraformUseCase {
	return &RunRecursiveTerraformUseCase{
		projectExecutor,
		projectObjectResolver,
	}
}

func (u RunRecursiveTerraformUseCase) Execute(execDir string, terraformArgs []string) error {
	projects, err := u.projectObjectResolver.Resolve(execDir)
	if err != nil {
		return err
	}
	return u.projectExecutor.Execute(projects, terraformArgs)
}
