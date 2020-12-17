package runtime

// Runtime represents a Runtime entity in KRE.
type Runtime struct {
	ID     string
	Name   string
	Status string
}

// List contains a list of  Runtime.
type List []Runtime

// List calls to KRE API and returns a list of Runtime entities.
func (s *runtimeClient) List() (List, error) {
	query := `
		query GetRuntimes {
			runtimes {
				id
				name
				status
			}
		}
  `

	var respData struct {
		Runtimes List
	}

	err := s.client.MakeRequest(query, nil, &respData)

	return respData.Runtimes, err
}
