package oura

import (
	"context"
	"fmt"
	"net/http"
	"net/url"
	"time"
)

// Sleep represents a single sleep entry
type Sleep struct {
	Awake                     int       `json:"awake"`
	BedtimeEnd                time.Time `json:"bedtime_end"`
	BedtimeEndDelta           int       `json:"bedtime_end_delta"`
	BedtimeStart              time.Time `json:"bedtime_start"`
	BedtimeStartDelta         int       `json:"bedtime_start_delta"`
	BreathAverage             float32   `json:"breath_average"`
	Deep                      int       `json:"deep"`
	Duration                  int       `json:"duration"`
	Efficiency                int       `json:"efficiency"`
	Hr5min                    []int     `json:"hr_5min"`
	HrAverage                 float32   `json:"hr_average"`
	HrLowest                  float32   `json:"hr_lowest"`
	Hypnogram5Min             string    `json:"hypnogram_5min"`
	IsLongest                 int       `json:"is_longest"`
	Light                     int       `json:"light"`
	MidpointAtDelta           int       `json:"midpoint_at_delta"`
	MidpointTime              int       `json:"midpoint_time"`
	OnsetLatency              int       `json:"onset_latency"`
	PeriodID                  int       `json:"period_id"`
	Rem                       int       `json:"rem"`
	Restless                  int       `json:"restless"`
	Rmssd                     int       `json:"rmssd"`
	Rmssd5min                 []int     `json:"rmssd_5min"`
	Score                     int       `json:"score"`
	ScoreAlignment            int       `json:"score_alignment"`
	ScoreDeep                 int       `json:"score_deep"`
	ScoreDisturbances         int       `json:"score_disturbances"`
	ScoreEfficiency           int       `json:"score_efficiency"`
	ScoreLatency              int       `json:"score_latency"`
	ScoreRem                  int       `json:"score_rem"`
	ScoreTotal                int       `json:"score_total"`
	SummaryDate               string    `json:"summary_date"`
	TemperatureDelta          float32   `json:"temperature_delta"`
	TemperatureDeviation      float32   `json:"temperature_deviation"`
	TemperatureTrendDeviation float32   `json:"temperature_trend_deviation"`
	Timezone                  int       `json:"timezone"`
	Total                     int       `json:"total"`
}

// Sleeps represents all sleep periods for the period requested
type Sleeps struct {
	Sleeps []Sleep `json:"sleep"`
}

// GetSleep gets all of the sleeps for a specified period of time.
// If a start and end date are not provided, ie are empty strings, we fall back to Oura which states:
// 	"If you omit the start date, it will be set to one week ago.
//	 If you omit the end date, it will be set to the current day."
func (c *Client) GetSleep(ctx context.Context, start string, end string) (*Sleeps, *http.Response, error) {
	path := "v1/sleep"
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

	var sleepSummaries *Sleeps
	resp, err := c.do(ctx, req, &sleepSummaries)
	if err != nil {
		return sleepSummaries, resp, err
	}

	return sleepSummaries, resp, nil
}
