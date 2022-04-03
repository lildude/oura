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
	AverageMet             float32   `json:"average_met"`
	CalActive              int       `json:"cal_active"`
	CalTotal               int       `json:"cal_total"`
	Class5min              string    `json:"class_5min"`
	DailyMovement          int       `json:"daily_movement"`
	DayEnd                 time.Time `json:"day_end"`
	DayStart               time.Time `json:"day_start"`
	High                   int       `json:"high"`
	Inactive               int       `json:"inactive"`
	InactivityAlerts       int       `json:"inactivity_alerts"`
	Low                    int       `json:"low"`
	Medium                 int       `json:"medium"`
	Met1min                []float32 `json:"met_1min"`
	MetMinHigh             int       `json:"met_min_high"`
	MetMinInactive         int       `json:"met_min_inactive"`
	MetMinLow              int       `json:"met_min_low"`
	MetMinMedium           int       `json:"met_min_medium"`
	MetMinMediumPlus       int       `json:"met_min_medium_plus"`
	NonWear                int       `json:"non_wear"`
	Rest                   int       `json:"rest"`
	RestModeState          int       `json:"rest_mode_state"`
	Score                  int       `json:"score"`
	ScoreMeetDailyTargets  int       `json:"score_meet_daily_targets"`
	ScoreMoveEveryHour     int       `json:"score_move_every_hour"`
	ScoreRecoveryTime      int       `json:"score_recovery_time"`
	ScoreStayActive        int       `json:"score_stay_active"`
	ScoreTrainingFrequency int       `json:"score_training_frequency"`
	ScoreTrainingVolume    int       `json:"score_training_volume"`
	Steps                  int       `json:"steps"`
	SummaryDate            string    `json:"summary_date"`
	TargetCalories         int       `json:"target_calories"`
	TargetKm               float32   `json:"target_km"`
	TargetMiles            float32   `json:"target_miles"`
	Timezone               int       `json:"timezone"`
	ToTargetKm             float32   `json:"to_target_km"`
	ToTargetMiles          float32   `json:"to_target_miles"`
	Total                  int       `json:"total"`
}

// Activities represents all activities for a the period requested
type Activities struct {
	Activities []Activity `json:"activity"`
}

// GetActivities gets all of the activities for a specified period of time.
// If a start and end date are not provided, ie are empty strings, we fall back to Oura which states:
// 	"If you omit the start date, it will be set to one week ago.
//	 If you omit the end date, it will be set to the current day."
func (c *Client) GetActivities(ctx context.Context, start string, end string) (*Activities, *http.Response, error) {
	path := "v1/activity"
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

	var activities *Activities
	resp, err := c.Do(ctx, req, &activities)
	if err != nil {
		return activities, resp, err
	}

	return activities, resp, nil
}
