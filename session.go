package oura

import (
	"context"
	"fmt"
	"net/http"
	"net/url"
	"time"
)

// Session represents the data returned from the Oura API for a single session.
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

// Sessions represents the data returned from the Oura API for a list of sessions.
type Sessions struct {
	Data      []Session `json:"data"`
	NextToken string    `json:"next_token"`
}

// Sessions gets the session data within a given timeframe.
// If a start and end date are not provided, ie are empty strings, we fall back to Oura's defaults which are:
// 	start_date: end_date - 1 day
//	end_date: current UTC date
func (c *Client) Sessions(ctx context.Context, start_date, end_date, next_token string) (*Sessions, *http.Response, error) {
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
