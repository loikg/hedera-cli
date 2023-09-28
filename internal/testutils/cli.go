// Package testutils contain test helper function
package testutils

import (
	"errors"
	"os"
	"os/exec"
	"path/filepath"
	"slices"
	"testing"

	"github.com/stretchr/testify/require"
)

// Testdata is a test helper function to read a golden file.
func Testdata(t *testing.T, filename string) []byte {
	t.Helper()

	b, err := os.ReadFile(filepath.Join("testdata", filename))
	require.NoError(t, err)

	return b
}

// RunCLI is a test help function to execute the cli binary with the given args arguments.
func RunCLI(t *testing.T, args ...string) []byte {
	t.Helper()

	args = slices.Insert(args, 0, "--network")
	args = slices.Insert(args, 1, "local")

	binPath := filepath.Join("..", "..", "hedera-cli")
	output, err := exec.Command(binPath, args...).Output()
	if err != nil {
		var exitError *exec.ExitError
		if errors.As(err, &exitError) {
			t.Fatalf("failed to execute cli binary: %v: %s", exitError, string(exitError.Stderr))
		}
		t.Fatalf("failed to execute cli binary: %v", err)
	}

	return output
}
