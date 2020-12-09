package oura

import (
	"context"
	"fmt"
	"net/http"
	"net/url"
)

// Readiness represents a single readiness entry
type Readiness struct {
	SummaryDate          string `json:"summary_date"`
	PeriodID             int    `json:"period_id"`
	Score                int    `json:"score"`
	ScorePreviousNight   int    `json:"score_previous_night"`
	ScoreSleepBalance    int    `json:"score_sleep_balance"`
	ScorePreviousDay     int    `json:"score_previous_day"`
	ScoreActivityBalance int    `json:"score_activity_balance"`
	ScoreRestingHr       int    `json:"score_resting_hr"`
	ScoreHrvBalance      int    `json:"score_hrv_balance"`
	ScoreRecoveryIndex   int    `json:"score_recovery_index"`
	ScoreTemperature     int    `json:"score_temperature"`
	RestModeState        int    `json:"rest_mode_state"`
}

// ReadinessSummaries represents all readiness periods for the period requested
type ReadinessSummaries struct {
	ReadinessSummaries []Readiness `json:"readiness"`
}

// Readiness gets all of the readiness entries for a specified period of time.
// If a start and end date are not provided, ie are empty strings, we fall back to Oura which states:
// 	"If you omit the start date, it will be set to one week ago.
//	 If you omit the end date, it will be set to the current day."
func (c *Client) Readiness(ctx context.Context, start string, end string) (*ReadinessSummaries, *http.Response, error) {
	path := "readiness"
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

	req, err := c.NewRequest("GET", path, nil)
	if err != nil {
		return nil, nil, err
	}

	var readinessSummaries *ReadinessSummaries
	resp, err := c.Do(ctx, req, &readinessSummaries)
	if err != nil {
		return readinessSummaries, resp, err
	}

	return readinessSummaries, resp, nil
}
