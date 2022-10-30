package oura

import (
	"context"
	"net/http"
	"time"
)

// Session represents the data returned from the Oura API for a single session.
type Session struct {
	// The date when the session occurred
	Day string `json:"day"`

	// The end datetime when the session occurred
	EndDatetime time.Time `json:"end_datetime"`

	// Timeseries data represented by an array of numbers; this data is available for sessions longer than 5 minutes
	HeartRate *timeSeriesData `json:"heart_rate,omitempty"`

	// Timeseries data represented by an array of numbers; this data is available for sessions longer than 3 minutes
	HeartRateVariability *timeSeriesData `json:"heart_rate_variability,omitempty"`

	// The user's selected mood after the session:
	// * `bad`
	// * `worse`
	// * `same`
	// * `good`
	// * `great`
	Mood *string `json:"mood,omitempty"`

	// Timeseries data represented by an array of numbers
	MotionCount *timeSeriesData `json:"motion_count,omitempty"`

	// The start datetime when the session occurred
	StartDatetime time.Time `json:"start_datetime"`

	// The session type:
	// * `breathing`
	// * `meditation`
	// * `nap`
	// * `relaxation`
	// * `rest`
	// * `body_status`
	Type string `json:"type"`
}

// Sessions represents the data returned from the Oura API for a list of sessions.
type Sessions struct {
	Data []Session `json:"data"`
	// Pagination token
	NextToken *string `json:"next_token,omitempty"`
}

// Sessions gets the session data within a given timeframe.
// If a start and end date are not provided, ie are empty strings, we fall back to Oura's defaults which are:
//
//	startDate: endDate - 1 day
//	endDate: current UTC date
func (c *Client) Sessions(ctx context.Context, startDate, endDate, nextToken string) (*Sessions, *http.Response, error) {
	path := parametiseDate("v2/usercollection/session", startDate, endDate, nextToken)
	req, err := c.NewRequest(ctx, "GET", path, nil)
	if err != nil {
		return nil, nil, err
	}

	var data *Sessions
	resp, err := c.do(req, &data)
	if err != nil {
		return data, resp, err
	}

	return data, resp, nil
}
