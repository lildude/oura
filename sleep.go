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
	SummaryDate				string		`json:"summary_date"`
  PeriodID					int				`json:"period_id"`
  IsLongest					int				`json:"is_longest"`
  Timezone					int				`json:"timezone"`
  BedtimeStart			time.Time	`json:"bedtime_start"`
  BedtimeEnd				time.Time	`json:"bedtime_end"`
  Score							int 			`json:"score"`
  ScoreTotal				int				`json:"score_total"`
  ScoreDisturbances	int				`json:"score_disturbances"`
  ScoreEfficiency		int				`json:"score_efficiency"`
  ScoreLatency			int				`json:"score_latency"`
  ScoreRem					int				`json:"score_rem"`
  ScoreDeep					int				`json:"score_deep"`
  ScoreAlignment		int				`json:"score_alignment"`
  Total							int				`json:"total"`
  Duration					int				`json:"duration"`
  Awake							int				`json:"awake"`
  Light							int				`json:"light"`
  Rem								int				`json:"rem"`
  Deep							int				`json:"deep"`
  OnsetLatency			int				`json:"onset_latency"`
  Restless					int				`json:"restless"`
  Efficiency				int				`json:"efficiency"`
  MidpointTime			int				`json:"midpoint_time"`
  HrLowest					int				`json:"hr_lowest"`
  HrAverage					float32		`json:"hr_average"`
  Rmssd							int				`json:"rmssd"`
  BreathAverage			int				`json:"breath_average"`
  TemperatureDelta	float32		`json:"temperature_delta"`
  Hypnogram5Min			string		`json:"hypnogram_5min"`
  Hr5min						[]int			`json:"hr_5min"`
  Rmssd5min					[]int			`json:"rmssd_5min"`
}

// Sleeps represents all sleep periods for the period requested
type Sleeps struct {
	Sleeps						[]Sleep		`json:"sleep"`
}

// Sleep gets all of the sleeps for a specified period of time.
// If a start and end date are not provided, ie are empty strings, we fall back to Oura which states:
// 	"If you omit the start date, it will be set to one week ago.
//	 If you omit the end date, it will be set to the current day."
func (c *Client) Sleep(ctx context.Context, start string, end string) (*Sleeps, *http.Response, error) {
	path := "sleep"
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

	var sleeps *Sleeps
	resp, err := c.Do(ctx, req, &sleeps)
	if err != nil {
		return sleeps, resp, err
	}

	return sleeps, resp, nil
}