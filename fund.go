package kap

import (
	"context"
	"net/url"
	"strconv"
)

// Funds returns the list of all funds, optionally filtered by the provided
// parameters.
func (c *Client) Funds(ctx context.Context, params *FundListParams) ([]Fund, error) {
	var q url.Values
	if params != nil {
		q = url.Values{}
		for _, s := range params.FundState {
			q.Add("fundState", s)
		}
		for _, s := range params.FundClass {
			q.Add("fundClass", s)
		}
		for _, s := range params.FundType {
			q.Add("fundType", s)
		}
	}

	var funds []Fund
	if err := c.get(ctx, "/api/vyk/funds", q, &funds); err != nil {
		return nil, err
	}
	return funds, nil
}

// FundDetail returns general information for a specific fund.
func (c *Client) FundDetail(ctx context.Context, fundID int) ([]DetailField, error) {
	path := "/api/vyk/fundDetail/" + strconv.Itoa(fundID)

	var fields []DetailField
	if err := c.get(ctx, path, nil, &fields); err != nil {
		return nil, err
	}
	return fields, nil
}
