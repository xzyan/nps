package tool

import (
	"bytes"
	"os/exec"
)

func CallShell(sh string) string {
	var stdout bytes.Buffer
	var stderr bytes.Buffer
	cmd := exec.Command("bash", "-c", sh)
	cmd.Stdout = &stdout
	cmd.Stderr = &stderr
	cmd.Dir = "."
	if cmd.Run() != nil {
		return stderr.String()
	} else {
		return stdout.String()
	}
}
