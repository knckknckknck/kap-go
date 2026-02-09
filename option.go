package kap

import (
	"net/http"
	"time"
)

// Option configures a Client. Use the With* functions to create Options.
type Option func(*Client)

// WithBaseURL sets the API base URL. This is useful for targeting the test
// environment or a proxy.
func WithBaseURL(url string) Option {
	return func(c *Client) {
		c.baseURL = url
	}
}

// WithTimeout sets the HTTP client timeout.
func WithTimeout(d time.Duration) Option {
	return func(c *Client) {
		c.httpClient.Timeout = d
	}
}

// WithHTTPClient replaces the default HTTP client entirely.
func WithHTTPClient(hc *http.Client) Option {
	return func(c *Client) {
		if hc != nil {
			c.httpClient = hc
		}
	}
}

// WithToken sets a pre-existing bearer token, skipping the need to call
// GenerateToken.
func WithToken(token string) Option {
	return func(c *Client) {
		c.token = token
	}
}

// basicAuth holds credentials for the test environment.
type basicAuth struct {
	Username string
	Password string
}

// WithBasicAuth configures basic authentication for the test environment.
// When set, bearer token authentication is not used.
func WithBasicAuth(username, password string) Option {
	return func(c *Client) {
		c.basicAuth = &basicAuth{
			Username: username,
			Password: password,
		}
	}
}
