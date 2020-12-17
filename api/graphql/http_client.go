package graphql

import (
	"errors"
	"net/http"
)

var (
	// ErrResponseEmpty used when response body is empty.
	ErrResponseEmpty = errors.New("response body is empty")
)

// option represents an argument to NewClient.
type option = func(http.RoundTripper) http.RoundTripper

// newHTTPClient initializes an http.Client with options.
func newHTTPClient(opts ...option) *http.Client {
	tr := http.DefaultTransport
	for _, opt := range opts {
		tr = opt(tr)
	}

	return &http.Client{Transport: tr}
}

// addHeader turns a RoundTripper into one that adds a request header.
func addHeader(name, value string) option {
	return func(tr http.RoundTripper) http.RoundTripper {
		return &funcTripper{
			roundTrip: func(req *http.Request) (*http.Response, error) {
				if req.Header.Get(name) == "" {
					req.Header.Add(name, value)
				}
				return tr.RoundTrip(req)
			},
		}
	}
}

type funcTripper struct {
	roundTrip func(*http.Request) (*http.Response, error)
}

func (tr funcTripper) RoundTrip(req *http.Request) (*http.Response, error) {
	return tr.roundTrip(req)
}
