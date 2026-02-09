package kap

import (
	"context"
	"encoding/json"
	"net/url"
	"strconv"
)

// Disclosures returns up to 50 disclosures starting from the given index.
// Optional filters can be provided via params.
func (c *Client) Disclosures(ctx context.Context, disclosureIndex int, params *DisclosureListParams) ([]Disclosure, error) {
	q := url.Values{}
	q.Set("disclosureIndex", strconv.Itoa(disclosureIndex))

	if params != nil {
		if params.DisclosureClass != "" {
			q.Set("disclosureClass", params.DisclosureClass)
		}
		if params.DisclosureType != "" {
			q.Set("disclosureTypes", params.DisclosureType)
		}
		if params.CompanyID != "" {
			q.Set("companyId", params.CompanyID)
		}
	}

	var disclosures []Disclosure
	if err := c.get(ctx, "/api/vyk/disclosures", q, &disclosures); err != nil {
		return nil, err
	}
	return disclosures, nil
}

// DisclosureDetail returns full details for a disclosure at the given index.
// fileType must be "html" or "data". subReportList is optional; when empty,
// all sub-reports are returned.
func (c *Client) DisclosureDetail(ctx context.Context, disclosureIndex int, fileType string, subReportList string) (*DisclosureDetail, error) {
	q := url.Values{}
	q.Set("fileType", fileType)
	if subReportList != "" {
		q.Set("subReportList", subReportList)
	}

	path := "/api/vyk/disclosureDetail/" + strconv.Itoa(disclosureIndex)

	var detail DisclosureDetail
	if err := c.get(ctx, path, q, &detail); err != nil {
		return nil, err
	}
	return &detail, nil
}

// LastDisclosureIndex returns the index of the most recently published
// disclosure.
func (c *Client) LastDisclosureIndex(ctx context.Context) (string, error) {
	var resp LastDisclosureIndexResponse
	if err := c.get(ctx, "/api/vyk/lastDisclosureIndex", nil, &resp); err != nil {
		return "", err
	}
	return resp.LastDisclosureIndex, nil
}

// BlockedDisclosures returns the list of blocked disclosures and attachments.
// The response schema is not fully documented, so the raw JSON is returned.
func (c *Client) BlockedDisclosures(ctx context.Context) (json.RawMessage, error) {
	var raw json.RawMessage
	if err := c.get(ctx, "/api/vyk/blockedDisclosures", nil, &raw); err != nil {
		return nil, err
	}
	return raw, nil
}
