package oura

import (
	"context"
	"net/http"
)

// Heartrate represents the data returned from the Oura API for a single heart rate measurement.
type Heartrate struct {
	// Beats per minute
	Bpm int `json:"bpm"`

	// The supported heart rate types:
	// * `awake`
	// * `rest`
	// * `sleep`
	// * `session`
	// * `live`
	Source string `json:"source"`

	// ISO 8601 formatted local timestamp indicating when the heart rate data was collected
	Timestamp string `json:"timestamp"`
}

// Heartrates represents the data returned from the Oura API for a list of heart rate measurements.
type Heartrates struct {
	Data []Heartrate `json:"data"`
	// Pagination token
	NextToken *string `json:"next_token,omitempty"`
}

// Heartrates gets the heart rate data for a specified Oura user within a given timeframe.
// If a start and end date are not provided, ie are empty strings, we fall back to Oura's defaults which are:
//
//	startDatetime: endDatetime - 1 day
//	endDatetime: current UTC date
func (c *Client) Heartrates(ctx context.Context, startDatetime, endDatetime, nextToken string) (*Heartrates, *http.Response, error) {
	path := parametiseDatetime("v2/usercollection/heartrate", startDatetime, endDatetime, nextToken)
	req, err := c.NewRequest(ctx, "GET", path, nil)
	if err != nil {
		return nil, nil, err
	}

	var data *Heartrates
	resp, err := c.do(req, &data)
	if err != nil {
		return data, resp, err
	}

	return data, resp, nil
}
