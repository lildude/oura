package oura

import (
	"context"
	"net/http"
	"time"
)

// Tag represents the data returned from the Oura API for a single tag.
type Tag struct {
	Day       string    `json:"day"`
	Text      string    `json:"text"`
	Timestamp time.Time `json:"timestamp"`
	Tags      []string  `json:"tags"`
}

// Tags represents the tag data returned from the Oura API within a given timeframe.
type Tags struct {
	Data      []Tag  `json:"data"`
	NextToken string `json:"next_token"`
}

// Tags gets the tag data within a given timeframe.
// If a start and end date are not provided, ie are empty strings, we fall back to Oura's defaults which are:
// 	start_date: end_date - 1 day
//	end_date: current UTC date
func (c *Client) Tags(ctx context.Context, start_date, end_date, next_token string) (*Tags, *http.Response, error) {
	path := parametiseDate("v2/usercollection/tag", start_date, end_date, next_token)
	req, err := c.NewRequest("GET", path, nil)
	if err != nil {
		return nil, nil, err
	}

	var data *Tags
	resp, err := c.Do(ctx, req, &data)
	if err != nil {
		return data, resp, err
	}

	return data, resp, nil
}
