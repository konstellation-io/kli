package api

//go:generate mockgen -source=${GOFILE} -destination=$PWD/mocks/${GOFILE} -package=mocks

import (
	"bytes"
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"

	"github.com/machinebox/graphql"

	"github.com/konstellation-io/kli/internal/config"
)

var ErrResponseEmpty = errors.New("response body is empty")

// ServerClient facilitates making HTTP requests to a Konstellation API.
type ServerClient struct {
	cfg    config.Config
	srv    config.ServerConfig
	client *graphql.Client
}

type ServerClienter interface {
	ListRuntimes() (RuntimeList, error)
}

// accessTokenResponse represents a response from the sign-in API.
type accessTokenResponse struct {
	Token string `json:"access_token"`
}

// NewServerClient initializes a ServerClient for a specific server.
func NewServerClient(cfg *config.Config, server config.ServerConfig, appVersion string) (ServerClienter, error) {
	accessToken, err := getAccessToken(cfg, server)
	if err != nil {
		return nil, err
	}

	c := newHTTPClient([]ClientOption{
		addHeader("User-Agent", "Konstellation KLI"),
		addHeader("KLI-Version", appVersion),
		addHeader("Cache-Control", "no-cache"),
		addHeader("Authorization", fmt.Sprintf("Bearer %s", accessToken)),
	}...)

	gql := graphql.NewClient(fmt.Sprintf("%s/graphql", server.URL), graphql.WithHTTPClient(c))

	return &ServerClient{
		srv:    server,
		client: gql,
	}, nil
}

func getAccessToken(cfg *config.Config, server config.ServerConfig) (string, error) {
	client := &http.Client{}
	url := fmt.Sprintf("%s/api/v1/auth/token/signin", server.URL)

	postData := bytes.NewBuffer([]byte(fmt.Sprintf(`{"apiToken":"%s"}`, server.APIToken)))

	ctx, cancel := context.WithTimeout(context.Background(), cfg.DefaultRequestTimeout)
	defer cancel()

	req, err := http.NewRequestWithContext(ctx, "POST", url, postData)
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

	var t accessTokenResponse

	err = json.NewDecoder(r.Body).Decode(&t)
	if err != nil {
		return "", err
	}

	return t.Token, nil
}
