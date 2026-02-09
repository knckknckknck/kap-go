package kap

import (
	"context"
	"net/url"
)

// CAEventStatus returns the status of a corporate action event identified
// by the given process reference ID.
func (c *Client) CAEventStatus(ctx context.Context, processRefID string) (*CAEventStatus, error) {
	q := url.Values{}
	q.Set("processRefId", processRefID)

	var status CAEventStatus
	if err := c.get(ctx, "/api/vyk/caEventStatus", q, &status); err != nil {
		return nil, err
	}
	return &status, nil
}
