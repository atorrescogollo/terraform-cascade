package orchestration

import (
	"fmt"
	"io"
	"os"
	"os/exec"
	"path/filepath"

	"github.com/atorrescogollo/terraform-cascade/internal/project"
	"github.com/atorrescogollo/terraform-cascade/internal/terraform"
)

type terraformExecutorFn = func(execDir string, args []string, stdout io.Writer, sterr io.Writer, stdin io.Reader) *exec.ExitError

// ProjectExecutor is the executor for terraform commands
// given a list of projects
type ProjectExecutor struct {
	terraformExecFn terraformExecutorFn
	stdout          io.Writer
	stderr          io.Writer
	stdin           io.Reader
}

// NewTerraformProjectExecutorWithOS ProjectExecutor constructor with terraform and default OS streams
func NewTerraformProjectExecutorWithOS() *ProjectExecutor {
	return NewProjectExecutor(terraform.Exec, os.Stdout, os.Stderr, os.Stdin)
}

// NewTerraformProjectExecutor ProjectExecutor constructor with terraform
func NewTerraformProjectExecutor(stdout io.Writer, stderr io.Writer, stdin io.Reader) *ProjectExecutor {
	return NewProjectExecutor(terraform.Exec, stdout, stderr, stdin)
}

// NewProjectExecutor ProjectExecutor constructor
func NewProjectExecutor(terraformExecFn terraformExecutorFn, stdout io.Writer, stderr io.Writer, stdin io.Reader) *ProjectExecutor {
	return &ProjectExecutor{
		terraformExecFn,
		stdout,
		stderr,
		stdin,
	}
}

// Execute executes terraform in each project
func (p ProjectExecutor) Execute(projects []project.TerraformProject, args []string) error {
	for _, project := range projects {
		p.stdout.Write([]byte(`


===============[ Running terraform in ` + project.RelativePath + ` ]===============

		`))
		exitErr := p.terraformExecFn(
			filepath.Join(project.BaseDir, project.RelativePath),
			args,
			p.stdout,
			p.stderr,
			p.stdin,
		)
		if exitErr != nil && exitErr.ExitCode() != 0 {
			fmt.Println("Error running terraform in", project.BaseDir, ":", exitErr)
			return exitErr
		}
	}
	return nil
}
