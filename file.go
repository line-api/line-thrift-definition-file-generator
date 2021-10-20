package main

import (
	"fmt"
	"os/exec"
)

func execCommand(command string, args ...string) (string, error) {
	cmd := exec.Command(command, args...)
	b, err := cmd.CombinedOutput()

	if err != nil {
		return "", err
	}

	if exitCode := cmd.ProcessState.ExitCode(); exitCode != 0 {
		return "", fmt.Errorf("failed to exec command: (exit code: %d, output: %s)", exitCode, string(b))
	}

	return string(b), nil
}

func decompileApkToSmali(apkPath string) string {
	return ""
}
