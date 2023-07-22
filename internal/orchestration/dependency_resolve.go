package orchestration

import (
	"fmt"
	"os"
	"os/exec"
	"strings"

	"github.com/atorrescogollo/terraform-cascade/internal/terraform"
)

type OrchestrateDir struct {
	IsProject   bool             `json:"is_project"`
	ProjectPath string           `json:"project_path"`
	Base        []OrchestrateDir `json:"base"`
	Others      []OrchestrateDir `json:"others"`
}

func NewOrchestrateDir(projectPath string) *OrchestrateDir {
	return &OrchestrateDir{
		IsProject:   false,
		ProjectPath: projectPath,
		Base:        make([]OrchestrateDir, 0),
		Others:      make([]OrchestrateDir, 0),
	}
}

func OrchestrateProjectDirectory(baseDir string) (*OrchestrateDir, error) {
	entries, err := os.ReadDir(baseDir)
	if err != nil {
		return nil, err
	}

	result := NewOrchestrateDir(baseDir)

	for _, entry := range entries {
		if strings.HasPrefix(entry.Name(), ".") {
			continue
		}
		if entry.IsDir() {
			orchestrateDir, err := OrchestrateProjectDirectory(baseDir + "/" + entry.Name())
			if err != nil {
				return nil, err
			}
			if entry.Name() == "base" {
				result.Base = append(result.Base, *orchestrateDir)
			} else {
				result.Others = append(result.Others, *orchestrateDir)
			}
		} else if entry.Name() == "backend.tf" {
			result.IsProject = true
		}
	}
	return result, nil
}

func RunTerraformRecursively(orchestrateDir OrchestrateDir, args []string, recursiveLevel int) *exec.ExitError {
	if orchestrateDir.IsProject {
		//utils.WaitForConfirmation("Run terraform in "+orchestrateDir.ProjectPath+"? (y/n)")
		fmt.Println(`


=============== Running terraform in ` + orchestrateDir.ProjectPath + ` ===============

		`)
		exitErr := terraform.TerraformExecWithOS(orchestrateDir.ProjectPath, args)
		if exitErr != nil && exitErr.ExitCode() != 0 {
			fmt.Println("Error: Terraform exited with code", exitErr.ExitCode())
			return exitErr
		}
	}
	for _, base := range orchestrateDir.Base {
		exitErr := RunTerraformRecursively(base, args, recursiveLevel+1)
		if exitErr != nil && exitErr.ExitCode() != 0 {
			return exitErr
		}
	}
	for _, other := range orchestrateDir.Others {
		exitErr := RunTerraformRecursively(other, args, recursiveLevel+1)
		if exitErr != nil && exitErr.ExitCode() != 0 {
			return exitErr
		}
	}
	return nil
}
