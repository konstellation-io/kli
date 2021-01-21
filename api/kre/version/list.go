package version

// Version represents a Version entity in KRE.
type Version struct {
	ID     string
	Name   string
	Status string
}

// List contains a list of  Version.
type List []Version

// List calls to KRE API and returns a list of Version entities.
func (c *versionClient) List() (List, error) {
	query := `
		query GetVersions() {
			versions() {
				id
				name
				status
			}
		}
	`

	var respData struct {
		Versions List
	}

	err := c.client.MakeRequest(query, nil, &respData)

	return respData.Versions, err
}
