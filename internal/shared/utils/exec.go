package utils

import (
	"fmt"
	"os"
	"os/exec"
)

// ExitWithErr exits with OS exit code that matches the exit code of the error
func ExitWithErr(err error) {
	var exitErr *exec.ExitError
	if err != nil {
		errCast, ok := err.(*exec.ExitError)
		if !ok {
			fmt.Println("Unexpected error:", err)
			os.Exit(1)
		}
		exitErr = errCast

	}
	if exitErr != nil && exitErr.ExitCode() != 0 {
		os.Exit(exitErr.ExitCode())
	}
	os.Exit(0)
}
