package utils

import (
	"os/exec"
)

var (
	pasteCmdArgs = "pbpaste"
	copyCmdArgs  = "pbcopy"
)

func getCopyCommand() *exec.Cmd {
	return exec.Command(copyCmdArgs)
}

// WriteTextToClipboard sends a given string to the System clipboard.
// Currently only compatible with macOS
func WriteTextToClipboard(value string) {
	copyCmd := getCopyCommand()
	in, err := copyCmd.StdinPipe()
	if err != nil {
		return
	}

	if err := copyCmd.Start(); err != nil {
		return
	}
	if _, err := in.Write([]byte(value)); err != nil {
		return
	}
	if err := in.Close(); err != nil {
		return
	}

	copyCmd.Wait()
	return
}
