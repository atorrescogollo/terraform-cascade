package utils

import (
	"fmt"
	"os"
	"os/exec"
)

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
