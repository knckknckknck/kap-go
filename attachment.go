package kap

import (
	"context"
	"io"
)

// DownloadAttachment downloads a disclosure attachment by its ID. It returns
// the response body as an io.ReadCloser, the Content-Disposition header
// value, and any error. The caller must close the returned ReadCloser.
func (c *Client) DownloadAttachment(ctx context.Context, id string) (io.ReadCloser, string, error) {
	path := "/api/vyk/downloadAttachment/" + id
	return c.getRaw(ctx, path, nil)
}
