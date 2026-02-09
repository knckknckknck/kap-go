package kap

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"
)

// get performs an authenticated GET request and decodes the JSON response
// into dest.
func (c *Client) get(ctx context.Context, path string, params url.Values, dest interface{}) error {
	resp, err := c.doRequest(ctx, path, params)
	if err != nil {
		return err
	}
	defer resp.Body.Close() //nolint:errcheck // response body close error is not actionable

	if resp.StatusCode < 200 || resp.StatusCode >= 300 {
		return c.handleErrorResponse(resp, path)
	}

	if err := json.NewDecoder(resp.Body).Decode(dest); err != nil {
		return &RequestError{Method: http.MethodGet, Path: path, Err: fmt.Errorf("decoding response: %w", err)}
	}
	return nil
}

// getRaw performs an authenticated GET request and returns the raw response
// body along with the Content-Disposition header value. The caller is
// responsible for closing the returned ReadCloser.
func (c *Client) getRaw(ctx context.Context, path string, params url.Values) (io.ReadCloser, string, error) {
	resp, err := c.doRequest(ctx, path, params)
	if err != nil {
		return nil, "", err
	}

	if resp.StatusCode < 200 || resp.StatusCode >= 300 {
		defer resp.Body.Close() //nolint:errcheck // response body close error is not actionable
		return nil, "", c.handleErrorResponse(resp, path)
	}

	contentDisposition := resp.Header.Get("Content-Disposition")
	return resp.Body, contentDisposition, nil
}

// doRequest builds and executes an authenticated HTTP GET request.
func (c *Client) doRequest(ctx context.Context, path string, params url.Values) (*http.Response, error) {
	reqURL := c.baseURL + path
	if len(params) > 0 {
		reqURL += "?" + params.Encode()
	}

	req, err := http.NewRequestWithContext(ctx, http.MethodGet, reqURL, nil)
	if err != nil {
		return nil, &RequestError{Method: http.MethodGet, Path: path, Err: err}
	}

	req.Header.Set("Content-Type", "application/json")

	if c.basicAuth != nil {
		req.SetBasicAuth(c.basicAuth.Username, c.basicAuth.Password)
	} else {
		c.mu.RLock()
		token := c.token
		c.mu.RUnlock()
		if token != "" {
			req.Header.Set("Authorization", token)
		}
	}

	resp, err := c.httpClient.Do(req)
	if err != nil {
		return nil, &RequestError{Method: http.MethodGet, Path: path, Err: err}
	}
	return resp, nil
}

// handleErrorResponse reads the response body and returns an *APIError.
func (c *Client) handleErrorResponse(resp *http.Response, path string) error {
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return &RequestError{Method: http.MethodGet, Path: path, Err: fmt.Errorf("reading error response: %w", err)}
	}

	var apiErr APIError
	if err := json.Unmarshal(body, &apiErr); err != nil {
		return &RequestError{
			Method: http.MethodGet,
			Path:   path,
			Err:    fmt.Errorf("HTTP %d: %s", resp.StatusCode, string(body)),
		}
	}
	apiErr.HTTPStatus = resp.StatusCode
	return &apiErr
}
