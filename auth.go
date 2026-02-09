package kap

import (
	"context"
	"net/url"
)

// GenerateToken requests a new bearer token from the KAP API using the
// client's API key. The token is stored on the client and used for all
// subsequent requests. It returns the token string.
func (c *Client) GenerateToken(ctx context.Context) (string, error) {
	params := url.Values{}
	params.Set("apiKey", c.apiKey)

	var resp TokenResponse
	if err := c.get(ctx, "/auth/generateToken", params, &resp); err != nil {
		return "", err
	}

	c.mu.Lock()
	c.token = resp.Token
	c.mu.Unlock()

	return resp.Token, nil
}
