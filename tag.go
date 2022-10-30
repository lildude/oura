package oura

import (
	"context"
	"net/http"
	"time"
)

// Tag represents the data returned from the Oura API for a single tag.
type Tag struct {
	// The `YYYY-MM-DD` formatted local date indicating when the tag was collected
	Day string `json:"day"`

	// A list of tags selected by the user. A translation of tag values can be found [here](https://cloud.ouraring.com/edu/tag-translations).
	Tags []string `json:"tags"`

	// Custom annotations associated with the tag, as provided by the user
	Text *string `json:"text,omitempty"`

	// ISO 8601 formatted local timestamp indicating when the tag was collected
	Timestamp time.Time `json:"timestamp"`
}

// Tags represents the tag data returned from the Oura API within a given timeframe.
type Tags struct {
	Data []Tag `json:"data"`
	// Pagination token
	NextToken *string `json:"next_token,omitempty"`
}

// Tags gets the tag data within a given timeframe.
// If a start and end date are not provided, ie are empty strings, we fall back to Oura's defaults which are:
//
//	startDate: endDate - 1 day
//	endDate: current UTC date
func (c *Client) Tags(ctx context.Context, startDate, endDate, nextToken string) (*Tags, *http.Response, error) {
	path := parametiseDate("v2/usercollection/tag", startDate, endDate, nextToken)
	req, err := c.NewRequest(ctx, "GET", path, nil)
	if err != nil {
		return nil, nil, err
	}

	var data *Tags
	resp, err := c.do(req, &data)
	if err != nil {
		return data, resp, err
	}

	return data, resp, nil
}
