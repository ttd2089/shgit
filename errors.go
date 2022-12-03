package shgit

import (
	"errors"
	"strings"
)

// ErrGitNotFound is returned when the git executable is not found.
var ErrGitNotFound error = errors.New("ErrGitNotFound")

// ErrGitCommandFailed is returned when the system fails to execute git. Note that this error is
// not returned when the system successfully exectutes git but the git command itself fails.
var ErrGitCommandFailed error = errors.New("ErrGitCommandFailed")

// A CLIError is the error type returned when a git command fails.
type CLIError struct {

	// ExitCode contains the exit code from the failed command.
	ExitCode int

	// Stdout contains any content that was written to stdout during the failed command.
	Stdout string

	// Stderr contains any content that was written to stderr during the failed command.
	Stderr string
}

// Error formats a CLIError using the first line from stderr
func (e *CLIError) Error() string {
	return strings.Split(e.Stderr, "\n")[0]
}
