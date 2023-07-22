package terraform

import (
	"bufio"
	"bytes"
	"fmt"
	"io"
	"log"
	"os"
	"os/exec"
)

func TerraformExecWithOS(execDir string, args []string) *exec.ExitError {
	return TerraformExec(execDir, args, os.Stdout, os.Stderr, os.Stdin)
}

func TerraformExec(execDir string, args []string, stdout io.Writer, sterr io.Writer, stdin io.Reader) *exec.ExitError {
	cmd := exec.Command("terraform", args...)
	cmd.Dir = execDir
	cmd.Stdout = stdout
	cmd.Stderr = sterr
	cmd.Stdin = stdin

	fmt.Println("Running terraform:", cmd.String(), "[dir=", execDir, "]")
	if err := cmd.Run(); err != nil {
		if exiterr, ok := err.(*exec.ExitError); ok {
			return exiterr
		}
		log.Panic(err)
	}
	return nil
}

func TerraformUsage() string {
	var buffer bytes.Buffer
	bufferWriter := bufio.NewWriter(&buffer)

	cwd, err := os.Getwd()
	if err != nil {
		log.Fatal(err)
	}
	exitErr := TerraformExec(cwd, []string{"--help"}, bufferWriter, bufferWriter, bytes.NewReader(nil))
	if exitErr != nil {
		log.Fatal(exitErr)
	}
	err = bufferWriter.Flush()
	if err != nil {
		log.Fatal(err)
	}
	return buffer.String()
}
