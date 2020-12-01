package config

import (
	"io/ioutil"
	"os"
	"testing"

	"github.com/stretchr/testify/require"
)

func setupConfigDir(t *testing.T) string {
	t.Helper()

	dir := makeTempConfigDir(t, "kli-test")

	err := os.Setenv("XDG_CONFIG_HOME", dir)
	if err != nil {
		t.Fatal(err)
	}

	return dir
}

func makeTempConfigDir(t *testing.T, pattern string) string {
	t.Helper()

	d, err := ioutil.TempDir("", pattern)
	if err != nil {
		t.Fatal(err)
	}

	return d
}

func cleanConfigDir(t *testing.T, dir string) {
	t.Helper()

	err := os.RemoveAll(dir)
	require.NoError(t, err)
}
