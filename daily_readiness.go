package oura

import (
	"context"
	"net/http"
	"time"
)

// DailyReadiness represents the readiness data for a single day.
type DailyReadiness struct {
	Contributors              ReadinessContributors `json:"contributors"`
	Day                       string                `json:"day"`
	Score                     *int                  `json:"score,omitempty"`
	TemperatureDeviation      *float32              `json:"temperature_deviation,omitempty"`
	TemperatureTrendDeviation *float32              `json:"temperature_trend_deviation,omitempty"`
	Timestamp                 time.Time             `json:"timestamp"`
}

// DailyReadinesses represents the readiness data for a given timeframe.
type DailyReadinesses struct {
	Data      []DailyReadiness `json:"data"`
	NextToken string           `json:"next_token"`
}

// Readiness score contributors
type ReadinessContributors struct {
	// Contribution of cumulative activity balance in range [1, 100].
	ActivityBalance *int `json:"activity_balance"`

	// Contribution of body temperature in range [1, 100].
	BodyTemperature *int `json:"body_temperature"`

	// Contribution of heart rate variability balance in range [1, 100].
	HrvBalance *int `json:"hrv_balance"`

	// Contribution of previous day's activity in range [1, 100].
	PreviousDayActivity *int `json:"previous_day_activity"`

	// Contribution of previous night's sleep in range [1, 100].
	PreviousNight *int `json:"previous_night"`

	// Contribution of recovery index in range [1, 100].
	RecoveryIndex *int `json:"recovery_index"`

	// Contribution of resting heart rate in range [1, 100].
	RestingHeartRate *int `json:"resting_heart_rate"`

	// Contribution of sleep balance in range [1, 100].
	SleepBalance *int `json:"sleep_balance"`
}

func (c *Client) DailyReadinesses(ctx context.Context, start_date, end_date, next_token string) (*DailyReadinesses, *http.Response, error) {
	path := parametiseDate("/v2/usercollection/daily_readiness", start_date, end_date, next_token)
	req, err := c.NewRequest("GET", path, nil)
	if err != nil {
		return nil, nil, err
	}

	var data *DailyReadinesses
	resp, err := c.do(ctx, req, &data)
	if err != nil {
		return data, resp, err
	}

	return data, resp, nil
}
