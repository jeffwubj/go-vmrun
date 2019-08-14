package vmrun

import (
	"bytes"
	"fmt"
	"os/exec"
)

func cmdStdout(cmd *exec.Cmd) ([]byte, error) {
	var stdout, stderr bytes.Buffer
	cmd.Stdout = &stdout
	cmd.Stderr = &stderr
	if err := cmd.Run(); err != nil {
		return nil, err
	}

	if len(stderr.Bytes()) != 0 {
		return stdout.Bytes(), fmt.Errorf(stderr.String())
	}

	return stdout.Bytes(), nil
}
