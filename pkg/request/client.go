package request

import (
	"net/http"
	"time"
)

// Option ...
type Option func(*http.Client)

// DefaultHTTPClient ...
func DefaultHTTPClient(opts ...Option) *http.Client {
	c := &http.Client{
		Timeout: 10 * time.Second,
	}

	for _, opt := range opts {
		opt(c)
	}

	return c
}

func WithTransport(t http.RoundTripper) func(r *http.Client) {
	return func(r *http.Client) {
		r.Transport = t
	}
}
