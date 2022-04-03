package oura

import (
	"context"
	"fmt"
	"net/http"
	"net/url"
	"time"
)

// DailyActivity represents the daily activity data for a given date range
type DailyActivity struct {
	Class5min                 string       `json:"class_5_min"`
	Score                     int          `json:"score"`
	ActiveCalories            int          `json:"active_calories"`
	AverageMetMinutes         float32      `json:"average_met_minutes"`
	Contributors              Contributors `json:"contributors"`
	EquivalentWalkingDistance int          `json:"equivalent_walking_distance"`
	HighActivityMetMinutes    int          `json:"high_activity_met_minutes"`
	HighActivityTime          int          `json:"high_activity_time"`
	InactivityAlerts          int          `json:"inactivity_alerts"`
	LowActivityMetMinutes     int          `json:"low_activity_met_minutes"`
	LowActivityTime           int          `json:"low_activity_time"`
	MediumActivityMetMinutes  int          `json:"medium_activity_met_minutes"`
	MediumActivityTime        int          `json:"medium_activity_time"`
	Met                       Met          `json:"met"`
	MetersToTarget            int          `json:"meters_to_target"`
	NonWearTime               int          `json:"non_wear_time"`
	RestingTime               int          `json:"resting_time"`
	SedentaryMetMinutes       int          `json:"sedentary_met_minutes"`
	SedentaryTime             int          `json:"sedentary_time"`
	Steps                     int          `json:"steps"`
	TargetCalories            int          `json:"target_calories"`
	TargetMeters              int          `json:"target_meters"`
	TotalCalories             int          `json:"total_calories"`
	Day                       string       `json:"day"`
	Timestamp                 time.Time    `json:"timestamp"`
}

type DailyActivities struct {
	Data      []DailyActivity `json:"data"`
	NextToken string          `json:"next_token"`
}

// Activity score contributors
type Contributors struct {
	MeetDailyTargets  int `json:"meet_daily_targets"`
	MoveEveryHour     int `json:"move_every_hour"`
	RecoveryTime      int `json:"recovery_time"`
	StayActive        int `json:"stay_active"`
	TrainingFrequency int `json:"training_frequency"`
	TrainingVolume    int `json:"training_volume"`
}

// MET (Metabolic equivalent) timeseries data
type Met struct {
	Interval  float32   `json:"interval"`
	Items     []float32 `json:"items"`
	Timestamp time.Time `json:"timestamp"`
}

// DailyActivity gets the daily activity summary values and detailed activity levels for a specified period of time.
// If a start and end date are not provided, ie are empty strings, we fall back to Oura's defaults which are:
// 	start_date: end_date - 1 day
//	end_date: current UTC date
func (c *Client) DailyActivity(ctx context.Context, start_date, end_date, next_token string) (*DailyActivities, *http.Response, error) {
	path := "v2/usercollection/daily_activity"
	params := url.Values{}

	if start_date != "" {
		params.Add("start_date", start_date)
	}
	if end_date != "" {
		params.Add("end_date", end_date)
	}
	if next_token != "" {
		params.Add("next_token", next_token)
	}
	if len(params) > 0 {
		path += fmt.Sprintf("?%s", params.Encode())
	}

	req, err := c.NewRequest("GET", path, nil)
	if err != nil {
		return nil, nil, err
	}

	var dailyActivities *DailyActivities
	resp, err := c.Do(ctx, req, &dailyActivities)
	if err != nil {
		return dailyActivities, resp, err
	}

	return dailyActivities, resp, nil
}
