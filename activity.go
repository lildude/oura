package oura

import (
	"context"
	"fmt"
	"net/http"
	"net/url"
	"time"
)

// Activity represents a single activity
type Activity struct {
	SummaryDate            string    `json:"summary_date"`
	DayStart               time.Time `json:"day_start"`
	DayEnd                 time.Time `json:"day_end"`
	Timezone               int       `json:"timezone"`
	Score                  int       `json:"score"`
	ScoreStayActive        int       `json:"score_stay_active"`
	ScoreMoveEveryHour     int       `json:"score_move_every_hour"`
	ScoreMeetDailyTargets  int       `json:"score_meet_daily_targets"`
	ScoreTrainingFrequency int       `json:"score_training_frequency"`
	ScoreTrainingVolume    int       `json:"score_training_volume"`
	ScoreRecoveryTime      int       `json:"score_recovery_time"`
	DailyMovement          int       `json:"daily_movement"`
	NonWear                int       `json:"non_wear"`
	Rest                   int       `json:"rest"`
	Inactive               int       `json:"inactive"`
	InactivityAlerts       int       `json:"inactivity_alerts"`
	Low                    int       `json:"low"`
	Medium                 int       `json:"medium"`
	High                   int       `json:"high"`
	Steps                  int       `json:"steps"`
	CalTotal               int       `json:"cal_total"`
	CalActive              int       `json:"cal_active"`
	MetMinInactive         int       `json:"met_min_inactive"`
	MetMinLow              int       `json:"met_min_low"`
	MetMinMediumPlus       int       `json:"met_min_medium_plus"`
	MetMinMedium           int       `json:"met_min_medium"`
	MetMinHigh             int       `json:"met_min_high"`
	AverageMet             float32   `json:"average_met"`
	Class5min              string    `json:"class_5min"`
	Met1min                []float32 `json:"met_1min"`
	RestModeState          int       `json:"rest_mode_state"`
}

// ActivitySummaries represents all activities for a the period requested
type ActivitySummaries struct {
	ActivitySummaries []Activity `json:"activity"`
}

// Activity gets all of the activities for a specified period of time.
// If a start and end date are not provided, ie are empty strings, we fall back to Oura which states:
// 	"If you omit the start date, it will be set to one week ago.
//	 If you omit the end date, it will be set to the current day."
func (c *Client) Activity(ctx context.Context, start string, end string) (*ActivitySummaries, *http.Response, error) {
	path := "activity"
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

	var activities *ActivitySummaries
	resp, err := c.Do(ctx, req, &activities)
	if err != nil {
		return activities, resp, err
	}

	return activities, resp, nil
}
