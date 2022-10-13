package oura

import (
	"context"
	"net/http"
)

// DailyActivity represents the data returned from the Oura API for a single activity.
type DailyActivity struct {
	// Active calories expended (in kilocalories)
	ActiveCalories int `json:"active_calories"`

	// Average metabolic equivalent (MET) in minutes
	AverageMetMinutes float32 `json:"average_met_minutes"`

	// 5-minute activity classification for the activity period:
	// * `0` non wear
	// * `1` rest
	// * `2` inactive
	// * `3` low activity
	// * `4` medium activity
	// * `5` high activity
	Class5Min *string `json:"class_5_min,omitempty"`

	// Activity score contributors
	Contributors ActivityContributors `json:"contributors"`

	// The `YYYY-MM-DD` formatted local date indicating when the daily activity occurred
	Day string `json:"day"`

	// Equivalent walking distance (in meters) of energy expenditure
	EquivalentWalkingDistance int `json:"equivalent_walking_distance"`

	// High activity metabolic equivalent (MET) in minutes
	HighActivityMetMinutes int `json:"high_activity_met_minutes"`

	// High activity metabolic equivalent (MET) in seconds
	HighActivityTime int `json:"high_activity_time"`

	// Number of inactivity alerts received
	InactivityAlerts int `json:"inactivity_alerts"`

	// Low activity metabolic equivalent (MET) in minutes
	LowActivityMetMinutes int `json:"low_activity_met_minutes"`

	// Low activity metabolic equivalent (MET) in seconds
	LowActivityTime int `json:"low_activity_time"`

	// Medium activity metabolic equivalent (MET) in minutes
	MediumActivityMetMinutes int `json:"medium_activity_met_minutes"`

	// Medium activity metabolic equivalent (MET) in seconds
	MediumActivityTime int `json:"medium_activity_time"`

	// Metabolic equivalent (MET) timeseries data represented by an array of numbers
	Met timeSeriesData `json:"met"`

	// Remaining meters to target (from `target_meters`)
	MetersToTarget int `json:"meters_to_target"`

	// The time (in seconds) in which the ring was not worn
	NonWearTime int `json:"non_wear_time"`

	// Resting time (in seconds)
	RestingTime int `json:"resting_time"`

	// Activity score in range `[1, 100]`
	Score *int `json:"score,omitempty"`

	// Sedentary metabolic equivalent (MET) in minutes
	SedentaryMetMinutes int `json:"sedentary_met_minutes"`

	// Sedentary metabolic equivalent (MET) in seconds
	SedentaryTime int `json:"sedentary_time"`

	// Total number of steps taken
	Steps int `json:"steps"`

	// Daily activity target (in kilocalories)
	TargetCalories int `json:"target_calories"`

	// Daily activity target (in meters)
	TargetMeters int `json:"target_meters"`

	// ISO 8601 formatted local timestamp indicating the start datetime of when the daily activity occurred
	Timestamp string `json:"timestamp"`

	// Total calories expended (in kilocalories)
	TotalCalories int `json:"total_calories"`
}

// DailyActivities represents the data returned from the Oura API for a list of daily activity summaries.
type DailyActivities struct {
	Data []DailyActivity `json:"data"`
	// Pagination token
	NextToken *string `json:"next_token,omitempty"`
}

// Contributors is an alias for ActivityContributors provided for compatibility with earlier users of this library.
type Contributors = ActivityContributors

// Activity score contributors
type ActivityContributors struct {
	// Contribution of meeting previous 7-day daily activity targets in range `[1, 100]`
	MeetDailyTargets *int `json:"meet_daily_targets,omitempty"`

	// Contribution of previous 24-hour inactivity alerts in range `[1, 100]`
	MoveEveryHour *int `json:"move_every_hour,omitempty"`

	// Contribution of previous 7-day recovery time in range `[1, 100]`
	RecoveryTime *int `json:"recovery_time,omitempty"`

	// Contribution of previous 24-hour activity in range `[1, 100]`
	StayActive *int `json:"stay_active,omitempty"`

	// Contribution of previous 7-day exercise frequency in range `[1, 100]`
	TrainingFrequency *int `json:"training_frequency,omitempty"`

	// Contribution of previous 7-day exercise volume in range `[1, 100]`
	TrainingVolume *int `json:"training_volume,omitempty"`
}

// DailyActivities gets the daily activity summary values and detailed activity levels for a specified period of time.
// If a start and end date are not provided, ie are empty strings, we fall back to Oura's defaults which are:
//
//	start_date: end_date - 1 day
//	end_date: current UTC date
func (c *Client) DailyActivities(ctx context.Context, start_date, end_date, next_token string) (*DailyActivities, *http.Response, error) {
	path := parametiseDate("v2/usercollection/daily_activity", start_date, end_date, next_token)
	req, err := c.NewRequest("GET", path, nil)
	if err != nil {
		return nil, nil, err
	}

	var data *DailyActivities
	resp, err := c.do(ctx, req, &data)
	if err != nil {
		return data, resp, err
	}

	return data, resp, nil
}
