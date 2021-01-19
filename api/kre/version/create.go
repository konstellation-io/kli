package version

import (
	"os"
	"path/filepath"

	"github.com/konstellation-io/graphql"
)

func (c *versionClient) Create(runtimeID, krtFile string) (string, error) {
	query := `
		mutation CreateVersion($input: CreateVersionInput!) {
			createVersion(input: $input) {
				id
				name
			}
		}
		`
	vars := map[string]interface{}{
		"input": map[string]interface{}{
			"runtimeId": runtimeID,
			"file":      nil,
		},
	}

	var respData struct {
		Data struct {
			CreateVersion struct {
				ID   string `json:"id"`
				Name string `json:"name"`
			} `json:"createVersion"`
		} `json:"data"`
	}

	r, err := os.Open(krtFile)
	if err != nil {
		return "", err
	}

	file := graphql.File{
		Field: "variables.input.file",
		Name:  filepath.Base(krtFile),
		R:     r,
	}

	err = c.client.UploadFile(file, query, vars, &respData)
	versionId := respData.Data.CreateVersion.ID

	return versionId, err
}
