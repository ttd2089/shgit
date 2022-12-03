package shgit

import (
	"bytes"
	"errors"
	"testing"

	"github.com/ttd2089/shellout"
	"github.com/ttd2089/tyers"
)

func TestRun(t *testing.T) {

	t.Run("Returns ErrGitNotFound when command is not found", func(t *testing.T) {
		underTest := &cli{}
		underTest.shell = &mockShell{
			err: tyers.New(shellout.ErrCommandNotFound, ""),
		}
		_, err := underTest.Run("return", "mock", "result")
		if !errors.Is(err, ErrGitNotFound) {
			t.Errorf("expected error '%v' to be an instance of %v\n", err, ErrGitNotFound)
		}
	})

	t.Run("Returns ErrGitCommandFailed when command process fails", func(t *testing.T) {
		underTest := &cli{}
		underTest.shell = &mockShell{
			err: tyers.New(shellout.ErrCommandProcessFailed, ""),
		}
		_, err := underTest.Run("return", "mock", "result")
		if !errors.Is(err, ErrGitCommandFailed) {
			t.Errorf("expected error '%v' to be an instance of %v\n", err, ErrGitCommandFailed)
		}
	})

	t.Run("Returns CLIError when command exits with non-zero status", func(t *testing.T) {
		expectedExitCode := 17
		expectedStdout := "expected stdout"
		expectedStderr := "expected stderr"
		expectedErr := &CLIError{
			ExitCode: expectedExitCode,
			Stdout:   expectedStdout,
			Stderr:   expectedStderr,
		}
		underTest := &cli{}
		underTest.shell = &mockShell{
			res: shellout.Result{
				ExitCode: expectedExitCode,
				Stdout:   bytes.NewBuffer([]byte(expectedStdout)),
				Stderr:   bytes.NewBuffer([]byte(expectedStderr)),
			},
		}
		_, err := underTest.Run("return", "mock", "result")
		var actualErr *CLIError
		if !errors.As(err, &actualErr) {
			t.Errorf("expected error '%v' to be an instance of CLIError\n", err)
		}
		if *expectedErr != *actualErr {
			t.Errorf("expected '%+v'; got '%+v'\n", *expectedErr, *actualErr)
		}
	})
}

type mockShell struct {
	res shellout.Result
	err error
}

func (m *mockShell) Run(cmd shellout.Cmd) (shellout.Result, error) {
	return m.res, m.err
}
