package oura

import (
	"context"
	"fmt"
	"net/http"
	"net/url"
	"time"
)

type Tag struct {
	Day       string    `json:"day"`
	Text      string    `json:"text"`
	Timestamp time.Time `json:"timestamp"`
	Tags      []string  `json:"tags"`
}

type Tags struct {
	Data      []Tag  `json:"data"`
	NextToken string `json:"next_token"`
}

func (c *Client) Tag(ctx context.Context, start_date, end_date, next_token string) (*Tags, *http.Response, error) {
	path := "v2/usercollection/tag"
	params := url.Values{}

	if start_date != "" {
		params.Add("start_date", start_date)
	}
	if end_date != "" {
		params.Add("end_date", end_date)
	}
	if next_token != "" {
		params.Add("next_token", next_token)
	}
	if len(params) > 0 {
		path += fmt.Sprintf("?%s", params.Encode())
	}

	req, err := c.NewRequest("GET", path, nil)
	if err != nil {
		return nil, nil, err
	}

	var tags *Tags
	resp, err := c.Do(ctx, req, &tags)
	if err != nil {
		return tags, resp, err
	}

	return tags, resp, nil
}
