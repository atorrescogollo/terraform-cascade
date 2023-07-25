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

type ProjectExecutor struct {
	terraformExecFn terraformExecutorFn
	stdout          io.Writer
	stderr          io.Writer
	stdin           io.Reader
}

func NewTerraformProjectExecutorWithOS() *ProjectExecutor {
	return NewProjectExecutor(terraform.TerraformExec, os.Stdout, os.Stderr, os.Stdin)
}

func NewTerraformProjectExecutor(stdout io.Writer, stderr io.Writer, stdin io.Reader) *ProjectExecutor {
	return NewProjectExecutor(terraform.TerraformExec, stdout, stderr, stdin)
}

func NewProjectExecutor(terraformExecFn terraformExecutorFn, stdout io.Writer, stderr io.Writer, stdin io.Reader) *ProjectExecutor {
	return &ProjectExecutor{
		terraformExecFn,
		stdout,
		stderr,
		stdin,
	}
}

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
