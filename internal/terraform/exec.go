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

// ExecWithOS executes terraform cli with default OS streams
func ExecWithOS(execDir string, args []string) *exec.ExitError {
	return Exec(execDir, args, os.Stdout, os.Stderr, os.Stdin)
}

// Exec executes terraform cli
func Exec(execDir string, args []string, stdout io.Writer, sterr io.Writer, stdin io.Reader) *exec.ExitError {
	cmd := exec.Command("terraform", args...)
	cmd.Dir = execDir
	cmd.Stdout = stdout
	cmd.Stderr = sterr
	cmd.Stdin = stdin

	fmt.Printf("%s\nRunning terraform %s\n", execDir, args)

	if err := cmd.Run(); err != nil {
		if exiterr, ok := err.(*exec.ExitError); ok {
			return exiterr
		}
		log.Panic(err)
	}
	return nil
}

// Usage outputs terraform usage
func Usage() string {
	var buffer bytes.Buffer
	bufferWriter := bufio.NewWriter(&buffer)

	cwd, err := os.Getwd()
	if err != nil {
		log.Fatal(err)
	}
	exitErr := Exec(cwd, []string{"--help"}, bufferWriter, bufferWriter, bytes.NewReader(nil))
	if exitErr != nil {
		log.Fatal(exitErr)
	}
	err = bufferWriter.Flush()
	if err != nil {
		log.Fatal(err)
	}
	return buffer.String()
}
