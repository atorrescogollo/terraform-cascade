package terraform

import (
	"bufio"
	"bytes"
	"io"
	"log"
	"os"
	"os/exec"
)

func TerraformExecWithOS(args []string) *exec.ExitError {
	return TerraformExec(args, os.Stdout, os.Stderr, os.Stdin)
}

func TerraformExec(args []string, stdout io.Writer, sterr io.Writer, stdin io.Reader) *exec.ExitError {
	cmd := exec.Command("terraform", args...)
	cmd.Stdout = stdout
	cmd.Stderr = sterr
	cmd.Stdin = stdin
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

	exitErr := TerraformExec([]string{"--help"}, bufferWriter, bufferWriter, bytes.NewReader(nil))
	if exitErr != nil {
		log.Fatal(exitErr)
	}
	err := bufferWriter.Flush()
	if err != nil {
		log.Fatal(err)
	}
	return buffer.String()
}
