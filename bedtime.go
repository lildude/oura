package oura

import (
	"context"
	"fmt"
	"net/http"
	"net/url"
)

// Bedtime represents a single bedtime recommendation.
type Bedtime struct {
	BedtimeWindow struct {
		Start int `json:"start"`
		End   int `json:"end"`
	} `json:"bedtime_window"`
	Date   string `json:"date"`
	Status string `json:"status"`
}

// IdealBedtimes represents all ideal bedtimes for the period requested.
type IdealBedtimes struct {
	IdealBedtimes []Bedtime `json:"ideal_bedtimes"`
}

// GetBedtime gets all of the ideal bedtimes for a specified period of time.
// If a start and end date are not provided, ie are empty strings, we fall back to Oura which states:
//
//	"If you omit the start date, it will be set to one week ago.
//	 If you omit the end date, it will be set to the current day."
func (c *Client) GetBedtime(ctx context.Context, start, end string) (*IdealBedtimes, *http.Response, error) {
	path := "v1/bedtime"
	params := url.Values{}

	if start != "" {
		params.Add("start", start)
	}
	if end != "" {
		params.Add("end", end)
	}
	if len(params) > 0 {
		path += fmt.Sprintf("?%s", params.Encode())
	}

	req, err := c.NewRequest(ctx, "GET", path, nil)
	if err != nil {
		return nil, nil, err
	}

	var bedtimes *IdealBedtimes
	resp, err := c.do(req, &bedtimes)
	if err != nil {
		return bedtimes, resp, err
	}

	return bedtimes, resp, nil
}
