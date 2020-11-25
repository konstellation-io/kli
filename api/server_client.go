package api

import (
	"bytes"
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"

	"github.com/konstellation-io/kli/internal/config"

	"github.com/shurcooL/graphql"
)

var ErrResponseEmpty = errors.New("response body is empty")

// ServerClient facilitates making HTTP requests to a Konstellation API.
type ServerClient struct {
	cfg    config.ServerConfig
	Client *graphql.Client
}

// AccessTokenResponse represents a response from the sign-in API.
type AccessTokenResponse struct {
	Token string `json:"access_token"`
}

// NewServerClient initializes a ServerClient for a specific server.
func NewServerClient(server config.ServerConfig, appVersion string) (*ServerClient, error) {
	authToken, err := getAccessToken(server)
	if err != nil {
		return nil, err
	}

	var opts []ClientOption

	opts = append(opts,
		addHeader("User-Agent", "Konstellation KLI"),
		addHeader("KLI-Version", appVersion),
		addHeader("Authorization", fmt.Sprintf("Bearer %s", authToken)),
	)

	c := newHTTPClient(opts...)
	gql := graphql.NewClient(fmt.Sprintf("%s/graphql", server.URL), c)

	return &ServerClient{
		cfg:    server,
		Client: gql,
	}, nil
}

func getAccessToken(server config.ServerConfig) (string, error) {
	client := &http.Client{}
	url := fmt.Sprintf("%s/api/v1/auth/token/signin", server.URL)

	postData := bytes.NewBuffer([]byte(fmt.Sprintf(`{"apiToken":"%s"}`, server.APIToken)))

	req, err := http.NewRequestWithContext(context.Background(), "POST", url, postData)
	if err != nil {
		return "", err
	}

	req.Header.Set("Content-Type", "application/json")

	r, err := client.Do(req)
	if err != nil {
		return "", err
	}
	defer r.Body.Close()

	if r.Body == nil {
		return "", ErrResponseEmpty
	}

	var t AccessTokenResponse

	err = json.NewDecoder(r.Body).Decode(&t)
	if err != nil {
		return "", err
	}

	return t.Token, nil
}
