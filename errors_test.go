package shgit

import (
	"fmt"
	"testing"
)

func TestCLIError(t *testing.T) {

	t.Run("Error()", func(t *testing.T) {

		t.Run("formats as the first line in Stderr", func(t *testing.T) {
			expected := "first line"
			underTest := &CLIError{
				Stderr: fmt.Sprintf("%s\nsecond line", expected),
			}
			actual := underTest.Error()
			if actual != expected {
				t.Errorf("expected '%s'; got '%s'", expected, actual)
			}
		})

	})
}
