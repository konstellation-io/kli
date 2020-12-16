package version

func (c *versionClient) Unpublish(versionID, comment string) error {
	query := `
		mutation UnpublishVersion($input: UnpublishVersionInput!) {
			unpublishVersion(input: $input) {
				id
				status
			}
		}
	`
	vars := map[string]interface{}{
		"input": map[string]string{
			"versionId": versionID,
			"comment":   comment,
		},
	}

	var respData struct {
		Status string
	}

	return c.client.MakeRequest(query, vars, &respData)
}
