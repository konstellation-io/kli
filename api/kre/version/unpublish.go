package version

func (c *versionClient) Unpublish(versionName, comment string) error {
	query := `
		mutation UnpublishVersion($input: UnpublishVersionInput!) {
			unpublishVersion(input: $input) {
				status
			}
		}
	`
	vars := map[string]interface{}{
		"input": map[string]string{
			"versionName": versionName,
			"comment":     comment,
		},
	}

	var respData struct {
		Status string
	}

	return c.client.MakeRequest(query, vars, &respData)
}
