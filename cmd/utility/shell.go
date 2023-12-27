package utility

import (
	"os"
	"os/exec"
)

func ExecuteCommand(commandString ...string) (cmd *exec.Cmd) {
	cmd = exec.Command(commandString[0],commandString[1:]...)
	cmd.Stdout = os.Stdout // Enroll standard out to system's
	cmd.Stderr = os.Stderr // Enroll standard error to system's
	cmd.Stdin = os.Stdin // Enroll standard input to system's
	return
}