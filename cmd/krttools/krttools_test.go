package krttools

import (
	"io/ioutil"
	"os"
	"path/filepath"
	"testing"

	"github.com/konstellation-io/kre/libs/krt-utils/pkg/validator"

	"github.com/MakeNowJust/heredoc"
	"github.com/stretchr/testify/require"
)

// nolint:gochecknoglobals
var (
	sampleKrtString = heredoc.Doc(`
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
)

func createKrtFile(t *testing.T, root, content, filename string) string {
	t.Helper()

	path := filepath.Join(root, filename)
	err := ioutil.WriteFile(path, []byte(content), 0600)
	require.NoError(t, err)

	return path
}

func createTestKrtContent(t *testing.T, root string, files ...string) {
	t.Helper()

	for _, name := range files {
		name = filepath.Join(root, name)
		path := filepath.Dir(name)

		_, err := os.Stat(path)
		if os.IsNotExist(err) {
			err = os.MkdirAll(path, 0755)
			require.NoError(t, err)
		}

		f, err := os.Create(name)

		defer func() {
			_ = f.Close()
		}()

		require.NoError(t, err)
	}
}

func TestNewKrtTools(t *testing.T) {
	krt := NewKrtTools().(*KrtTools)
	require.NotEmpty(t, krt.validator)
	require.NotEmpty(t, krt.builder)
}

func TestKrtTools_Validate(t *testing.T) {
	tempDir, err := ioutil.TempDir("", "TestKrtTools_Validate")
	require.NoError(t, err)

	defer os.RemoveAll(tempDir)

	yamlFile := createKrtFile(t, tempDir, sampleKrtString, ".krt.yaml")
	krt := NewKrtTools()
	err = krt.Validate(yamlFile)
	require.NoError(t, err)
}

func TestKrtTools_Build(t *testing.T) {
	files := []string{
		"src/py-test/main.py",
		"src/go-test/go.go",
		"docs/README.md",
		"metrics/dashboards/models.json",
		"metrics/dashboards/application.json",
	}

	// Create test dir structure
	tempDir, err := ioutil.TempDir("", "TestKrtTools_Build")
	require.NoError(t, err)

	defer os.RemoveAll(tempDir)

	createTestKrtContent(t, tempDir, files...)
	createKrtFile(t, tempDir, sampleKrtString, "krt.yaml")

	krt := NewKrtTools()
	target := "test.krt"
	err = krt.Build(tempDir, target, "testversion")
	require.NoError(t, err)

	defer os.Remove(target)

	yamlFile := filepath.Join(tempDir, "krt.yaml")
	f, err := os.Open(yamlFile)
	require.NoError(t, err)

	v := validator.New()
	yamlFields, err := v.Parse(f)
	require.NoError(t, err)
	require.Equal(t, "testversion", yamlFields.Version)
}
