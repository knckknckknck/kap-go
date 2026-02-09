package kap

import (
	"net/http"
	"sync"
	"time"
)

const (
	// DefaultBaseURL is the KAP test environment base URL.
	DefaultBaseURL = "https://apigwdev.mkk.com.tr"

	// DefaultTimeout is the default HTTP client timeout.
	DefaultTimeout = 30 * time.Second
)

// Client is a KAP API client. It is safe for concurrent use.
type Client struct {
	baseURL    string
	apiKey     string
	token      string
	httpClient *http.Client
	basicAuth  *basicAuth

	mu sync.RWMutex // protects token
}

// NewClient creates a new KAP API client. The apiKey is required for
// production environments; it can be empty when using WithBasicAuth for
// the test environment.
func NewClient(apiKey string, opts ...Option) *Client {
	c := &Client{
		baseURL: DefaultBaseURL,
		apiKey:  apiKey,
		httpClient: &http.Client{
			Timeout: DefaultTimeout,
		},
	}
	for _, opt := range opts {
		opt(c)
	}
	return c
}
