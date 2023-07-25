package orchestration

import (
	"io"
	"os"
	"os/exec"
	"testing"

	"github.com/atorrescogollo/terraform-cascade/internal/project"
	"github.com/stretchr/testify/assert"
)

func TestExecute(t *testing.T) {
	var executions []struct {
		execDir string
		args    []string
	}
	r, w, _ := os.Pipe()
	executor := NewProjectExecutor(
		func(execDir string, args []string, stdout io.Writer, sterr io.Writer, stdin io.Reader) *exec.ExitError {
			executions = append(executions, struct {
				execDir string
				args    []string
			}{execDir, args})
			return nil
		},
		w,
		w,
		r,
	)

	args := []string{"init"}
	err := executor.Execute(
		[]project.TerraformProject{
			{BaseDir: ".", RelativePath: "."},
			{BaseDir: ".", RelativePath: "base"},
			{BaseDir: ".", RelativePath: "base/network"},
			{BaseDir: ".", RelativePath: "dev"},
			{BaseDir: ".", RelativePath: "dev/network"},
			{BaseDir: ".", RelativePath: "prod"},
			{BaseDir: ".", RelativePath: "prod/network"},
		},
		args,
	)

	assert.Nil(t, err)
	assert.Equal(t, []struct {
		execDir string
		args    []string
	}{
		{".", args},
		{"base", args},
		{"base/network", args},
		{"dev", args},
		{"dev/network", args},
		{"prod", args},
		{"prod/network", args},
	}, executions)
}
