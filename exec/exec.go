package exec

import (
	"context"
	"os/exec"
)

func CommandContext(ctx context.Context, name string, arg ...string) error {
	cmd := exec.CommandContext(ctx, name, arg...)
	stdout := &writer{
		prefix: name + "_out",
	}
	stderr := &writer{
		prefix: name + "_err",
	}
	cmd.Stdout = stdout
	cmd.Stderr = stderr
	err := cmd.Run()
	stdout.flush()
	stderr.flush()
	return err
}
