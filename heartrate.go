package oura

import (
	"context"
	"fmt"
	"net/http"
	"net/url"
	"time"
)

type Heartrate struct {
	Bpm       int       `json:"bpm"`
	Source    string    `json:"source"`
	Timestamp time.Time `json:"timestamp"`
}

type Heartrates struct {
	Data      []Heartrate `json:"data"`
	NextToken string      `json:"next_token"`
}

func (c *Client) Heartrate(ctx context.Context, start_datetime, end_datetime, next_token string) (*Heartrates, *http.Response, error) {
	path := "v2/usercollection/heartrate"
	params := url.Values{}

	if start_datetime != "" {
		params.Add("start_datetime", start_datetime)
	}
	if end_datetime != "" {
		params.Add("end_datetime", end_datetime)
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

	var heartrates *Heartrates
	resp, err := c.Do(ctx, req, &heartrates)
	if err != nil {
		return heartrates, resp, err
	}

	return heartrates, resp, nil
}
