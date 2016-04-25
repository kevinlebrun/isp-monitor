package main

import (
	"fmt"
	"os/exec"
)

func assertCommandExists(command string) error {
	_, err := exec.LookPath("speedtest-cli")
	if err != nil {
		return fmt.Errorf("%s is required to continue", "speedtest-cli")
	}

	return nil
}
