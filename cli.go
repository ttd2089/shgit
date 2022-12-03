// Package shgit exports an interface for shelling out to the Git CLI.
package shgit

import (
	"errors"

	"github.com/ttd2089/shellout"
	"github.com/ttd2089/tyers"
)

// A CLI provides an interface to execute Git commands.
type CLI interface {

	// Run executes an arbitrary git command. Note that the cmd slice should not contain the
	// executable name.
	Run(cmd ...string) (string, error)
}

// NewCLI returns a new instance of CLI.
func NewCLI() CLI {
	return &cli{
		gitPath: "git",
		shell:   shellout.New(),
	}
}

type cli struct {
	gitPath string
	shell   shellout.Shell
}

func (c *cli) Run(cmd ...string) (string, error) {
	res, err := c.shell.Run(shellout.Cmd{
		Command: c.gitPath,
		Args:    cmd,
	})
	if errors.Is(err, shellout.ErrCommandNotFound) {
		return "", tyers.As(ErrGitNotFound, err)
	}
	if err != nil {
		return "", tyers.As(ErrGitCommandFailed, err)
	}
	if res.ExitCode != 0 {
		return "", &CLIError{
			ExitCode: res.ExitCode,
			Stdout:   res.Stdout.String(),
			Stderr:   res.Stderr.String(),
		}
	}
	return res.Stdout.String(), nil
}
