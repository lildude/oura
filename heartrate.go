package oura

import (
	"context"
	"fmt"
	"net/http"
	"net/url"
	"time"
)

// Heartrate represents the data returned from the Oura API for a single heart rate measurement.
type Heartrate struct {
	Bpm       int       `json:"bpm"`
	Source    string    `json:"source"`
	Timestamp time.Time `json:"timestamp"`
}

// Heartrates represents the data returned from the Oura API for a list of heart rate measurements.
type Heartrates struct {
	Data      []Heartrate `json:"data"`
	NextToken string      `json:"next_token"`
}

// Heartrate gets the heart rate data for a specified Oura user within a given timeframe.
// If a start and end date are not provided, ie are empty strings, we fall back to Oura's defaults which are:
// 	start_datetime: end_datetime - 1 day
//	end_datetime: current UTC date
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
