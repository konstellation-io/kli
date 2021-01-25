package version

func (c *versionClient) Start(versionName, comment string) error {
	query := `
		mutation StartVersion($input: StartVersionInput!) {
			startVersion(input: $input) {
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
