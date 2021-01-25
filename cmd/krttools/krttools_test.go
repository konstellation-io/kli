package krttools

import (
	"io/ioutil"
	"os"
	"path/filepath"
	"testing"

	"github.com/MakeNowJust/heredoc"
	"github.com/stretchr/testify/require"
)

func createKrtFile(t *testing.T, root, content, filename string) string {
	t.Helper()

	path := filepath.Join(root, filename)
	err := ioutil.WriteFile(path, []byte(content), 0600)
	require.NoError(t, err)

	return path
}

func TestNewKrtTools(t *testing.T) {
	krt := NewKrtTools().(*KrtTools)
	require.NotEmpty(t, krt.validator)
	require.NotEmpty(t, krt.builder)
}

func TestKrtTools_Validate(t *testing.T) {
	sampleKrtString := heredoc.Doc(`
    version: test-v1
    description: Version for testing.
    entrypoint:
      proto: public_input.proto
      image: konstellation/kre-entrypoint:latest
    config:
      variables:
        - SOME_CONFIG_VAR
      files:
        - SOME_FILE
    nodes:
      - name: py-test
        image: konstellation/kre-py:latest
        src: src/py-test/main.py
        gpu: true
    workflows:
      - name: py-test
        entrypoint: PyTest
        sequential:
          - py-test
	`)

	tempDir, err := ioutil.TempDir("", "TestKrtTools_Validate")
	require.NoError(t, err)

	defer os.RemoveAll(tempDir)

	yamlFile := createKrtFile(t, tempDir, sampleKrtString, ".krt.yaml")
	krt := NewKrtTools()
	err = krt.Validate(yamlFile)
	require.NoError(t, err)
}
