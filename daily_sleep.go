package oura

import (
	"context"
	"net/http"
	"time"
)

// DailySleep represents the sleep data for a single day.
type DailySleep struct {
	Contributors SleepContributors `json:"contributors"`
	Day          string            `json:"day"`
	Score        *int              `json:"score,omitempty"`
	Timestamp    time.Time         `json:"timestamp"`
}

// DailySleeps represents the sleep data for a given timeframe.
type DailySleeps struct {
	Data      []DailySleep `json:"data"`
	NextToken *string      `json:"next_token,omitempty"`
}

// SleepContributors represents all the contributors to the sleep score.
type SleepContributors struct {
	// Contribution of deep sleep in range `[1, 100]`.
	DeepSleep *int `json:"deep_sleep"`

	// Contribution of sleep efficiency in range `[1, 100]`.
	Efficiency *int `json:"efficiency"`

	// Contribution of sleep latency in range `[1, 100]`.
	Latency *int `json:"latency"`

	// Contribution of REM sleep in range `[1, 100]`.
	RemSleep *int `json:"rem_sleep"`

	// Contribution of sleep restfulness in range `[1, 100]`.
	Restfulness *int `json:"restfulness"`

	// Contribution of sleep timing in range `[1, 100]`.
	Timing *int `json:"timing"`

	// Contribution of total sleep in range `[1, 100]`.
	TotalSleep *int `json:"total_sleep"`
}

// DailySleeps gets the daily sleep data for a specified period of time.
// If a start and end date are not provided, ie are empty strings, we fall back to Oura's defaults which are:
//
//	startDate: endDate - 1 day
//	endDate: current UTC date
func (c *Client) DailySleeps(ctx context.Context, startDate, endDate, nextToken string) (*DailySleeps, *http.Response, error) {
	path := parametiseDate("/v2/usercollection/daily_sleep", startDate, endDate, nextToken)
	req, err := c.NewRequest(ctx, "GET", path, nil)
	if err != nil {
		return nil, nil, err
	}

	var data *DailySleeps
	resp, err := c.do(req, &data)
	if err != nil {
		return data, resp, err
	}

	return data, resp, nil
}
