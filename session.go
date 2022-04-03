package oura

import (
	"context"
	"fmt"
	"net/http"
	"net/url"
	"time"
)

type Session struct {
	Day                  string         `json:"day"`
	StartDatetime        time.Time      `json:"start_date"`
	EndDatetime          time.Time      `json:"end_date"`
	Type                 string         `json:"type"`
	Heartrate            TimeSeriesData `json:"heart_rate"`
	HeartrateVariability TimeSeriesData `json:"heart_rate_variability"`
	Mood                 string         `json:"mood"`
	MotionCount          TimeSeriesData `json:"motion_count"`
}

type Sessions struct {
	Data      []Session `json:"data"`
	NextToken string    `json:"next_token"`
}

type TimeSeriesData struct {
	Interval  float32   `json:"interval"`
	Items     []float32 `json:"items"`
	Timestamp time.Time `json:"timestamp"`
}

func (c *Client) Session(ctx context.Context, start_date, end_date, next_token string) (*Sessions, *http.Response, error) {
	path := "v2/usercollection/session"
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

	var sessions *Sessions
	resp, err := c.Do(ctx, req, &sessions)
	if err != nil {
		return sessions, resp, err
	}

	return sessions, resp, nil
}
