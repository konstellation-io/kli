package version

func (c *versionClient) Publish(versionID, comment string) error {
	query := `
		mutation PublishVersion($input: PublishVersionInput!) {
			publishVersion(input: $input) {
				id
				status
			}
		}
	`
	vars := map[string]interface{}{
		"input": map[string]string{"versionId": versionID, "comment": comment},
	}

	var respData struct {
		Status string
	}

	return c.client.MakeRequest(query, vars, &respData)
}
