package kap

import (
	"context"
	"strconv"
)

// Members returns the list of all KAP member companies.
func (c *Client) Members(ctx context.Context) ([]Member, error) {
	var members []Member
	if err := c.get(ctx, "/api/vyk/members", nil, &members); err != nil {
		return nil, err
	}
	return members, nil
}

// MemberSecurities returns listed companies with their securities.
func (c *Client) MemberSecurities(ctx context.Context) ([]MemberSecurities, error) {
	var ms []MemberSecurities
	if err := c.get(ctx, "/api/vyk/memberSecurities", nil, &ms); err != nil {
		return nil, err
	}
	return ms, nil
}

// MemberDetail returns general information for a specific company.
func (c *Client) MemberDetail(ctx context.Context, id int) ([]DetailField, error) {
	path := "/api/vyk/memberDetail/" + strconv.Itoa(id)

	var fields []DetailField
	if err := c.get(ctx, path, nil, &fields); err != nil {
		return nil, err
	}
	return fields, nil
}
