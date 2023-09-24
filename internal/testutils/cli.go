// Package testutils contain test helper function
package testutils

import (
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
func RunCLI(t *testing.T, args ...string) ([]byte, error) {
	t.Helper()

	args = slices.Insert(args, 0, "--network")
	args = slices.Insert(args, 1, "local")

	binPath := filepath.Join("../", "hedera-cli")
	return exec.Command(binPath, args...).Output()
}
