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

// SleepPeriod represents a sleep period.
type SleepPeriod struct {
	AverageBreath       *float32          `json:"average_breath,omitempty"`
	AverageHeartRate    *float32          `json:"average_heart_rate,omitempty"`
	AverageHrv          *int              `json:"average_hrv,omitempty"`
	AwakeTime           *int              `json:"awake_time,omitempty"`
	BedtimeEnd          time.Time         `json:"bedtime_end"`
	BedtimeStart        time.Time         `json:"bedtime_start"`
	Day                 string            `json:"day"`
	DeepSleepDuration   *int              `json:"deep_sleep_duration,omitempty"`
	Efficiency          *int              `json:"efficiency,omitempty"`
	HeartRate           *timeSeriesData   `json:"heart_rate,omitempty"`
	Hrv                 *timeSeriesData   `json:"hrv,omitempty"`
	Latency             *int              `json:"latency,omitempty"`
	LightSleepDuration  *int              `json:"light_sleep_duration,omitempty"`
	LowBatteryAlert     bool              `json:"low_battery_alert"`
	LowestHeartRate     *int              `json:"lowest_heart_rate,omitempty"`
	Movement30Sec       *string           `json:"movement_30_sec,omitempty"`
	Period              int               `json:"period"`
	Readiness           *ReadinessSummary `json:"readiness,omitempty"`
	ReadinessScoreDelta *int              `json:"readiness_score_delta,omitempty"`
	RemSleepDuration    *int              `json:"rem_sleep_duration,omitempty"`
	RestlessPeriods     *int              `json:"restless_periods,omitempty"`
	SleepPhase5Min      *string           `json:"sleep_phase_5_min,omitempty"`
	SleepScoreDelta     *int              `json:"sleep_score_delta,omitempty"`
	TimeInBed           int               `json:"time_in_bed"`
	TotalSleepDuration  *int              `json:"total_sleep_duration,omitempty"`
	Type                string            `json:"type,omitempty"`
}

// SleepPeriods represents the sleep data for a given timeframe.
type SleepPeriods struct {
	Data []SleepPeriod `json:"data"`
	// Pagination token
	NextToken *string `json:"next_token,omitempty"`
}

type ReadinessSummary struct {
	Contributors              ReadinessContributors `json:"contributors"`
	Score                     *int                  `json:"score,omitempty"`
	TemperatureDeviation      *float32              `json:"temperature_deviation,omitempty"`
	TemperatureTrendDeviation *float32              `json:"temperature_trend_deviation,omitempty"`
}

// GetSleep gets all of the sleeps for a specified period of time.
// If a start and end date are not provided, ie are empty strings, we fall back to Oura which states:
//
//	"If you omit the start date, it will be set to one week ago.
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

// Sleeps gets the detailed sleep data for a specified period of time.
// If a start and end date are not provided, ie are empty strings, we fall back to Oura's defaults which are:
//
//	start_date: end_date - 1 day
//	end_date: current UTC date
func (c *Client) Sleeps(ctx context.Context, start_date, end_date, next_token string) (*SleepPeriods, *http.Response, error) {
	path := parametiseDate("v2/usercollection/sleep", start_date, end_date, next_token)
	req, err := c.NewRequest("GET", path, nil)
	if err != nil {
		return nil, nil, err
	}

	var data *SleepPeriods
	resp, err := c.do(ctx, req, &data)
	if err != nil {
		return data, resp, err
	}

	return data, resp, nil
}
