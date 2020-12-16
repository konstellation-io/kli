package version

func (c *versionClient) Stop(versionID, comment string) error {
	query := `
		mutation StopVersion($input: StopVersionInput!) {
			stopVersion(input: $input) {
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
