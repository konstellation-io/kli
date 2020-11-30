package http

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/konstellation-io/kli/internal/config"
)

// accessTokenResponse represents a response from the sign-in API.
type accessTokenResponse struct {
	Token string `json:"access_token"`
}

func GetAccessToken(cfg *config.Config, server *config.ServerConfig) (string, error) {
	client := &http.Client{}
	url := fmt.Sprintf("%s/api/v1/auth/token/signin", server.URL)

	postData := bytes.NewBuffer([]byte(fmt.Sprintf(`{"apiToken":"%s"}`, server.APIToken)))

	ctx, cancel := context.WithTimeout(context.Background(), cfg.DefaultRequestTimeout)
	defer cancel()

	req, err := http.NewRequestWithContext(ctx, "POST", url, postData)
	if err != nil {
		return "", fmt.Errorf("error creating request call: %s", err)
	}

	req.Header.Set("Content-Type", "application/json")

	r, err := client.Do(req)
	if err != nil {
		return "", fmt.Errorf("error calling access token URL: %s", err)
	}
	defer r.Body.Close()

	if r.Body == nil {
		return "", ErrResponseEmpty
	}

	var t accessTokenResponse

	err = json.NewDecoder(r.Body).Decode(&t)
	if err != nil {
		return "", fmt.Errorf("error decoding access token response: %s", err)
	}

	return t.Token, nil
}
