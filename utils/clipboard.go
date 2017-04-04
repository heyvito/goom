package utils

import (
	"fmt"
	"os"
	"os/exec"
	"runtime"
)

var copyCmdArgs string

func init() {
	if runtime.GOOS == "darwin" {
		copyCmdArgs = "pbcopy"
	} else {
		if err := exec.Command("which", "xclip").Run(); err != nil {
			fmt.Println("On non-macOS environments, goom requires xclip, which could not be found.")
			os.Exit(1)
		} else {
			copyCmdArgs = "xclip"
		}
	}
}

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
