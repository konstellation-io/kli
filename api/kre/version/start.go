package version

func (c *versionClient) Start(versionID, comment string) error {
	query := `
		mutation StartVersion($input: StartVersionInput!) {
			startVersion(input: $input) {
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
