package exec

import (
	"bytes"
	"fmt"
	"os/exec"
)

type Executor struct{}

func (e *Executor) ExecuteCommand(command string, args ...string) (string, error) {
	cmd := exec.Command(command, args...)
	var out bytes.Buffer
	cmd.Stdout = &out
	err := cmd.Run()
	if err != nil {
		return "", fmt.Errorf("error executing command: %w", err)
	}
	return out.String(), nil
}
