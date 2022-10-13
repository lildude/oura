package oura

import (
	"context"
	"net/http"
	"time"
)

// Workout represents the data returned from the Oura API for a single workout.
type Workout struct {
	// The workout activity type
	Activity string `json:"activity"`

	// The number of calories (kcal) burned during the workout
	Calories *float32 `json:"calories,omitempty"`

	// The `YYYY-MM-DD` formatted local date indicating when the workout was recorded
	Day string `json:"day"`

	// The distance (measured in meters) traveled during the workout
	Distance *float32 `json:"distance,omitempty"`

	// ISO 8601 formatted local timestamp indicating when the workout ended
	EndDatetime time.Time `json:"end_datetime"`

	// The workout intensity:
	// * `easy`
	// * `moderate`
	// * `hard`
	Intensity string `json:"intensity"`

	// User-defined label for the workout
	Label *string `json:"label,omitempty"`

	// The data source where the Workout data was collected from:
	// * `manual` Workouts which were manually entered by the user
	// * `autodetected` Workouts autodetected by Oura
	// * `confirmed` Workouts autodetected by Oura and confirmed by the user
	// * `workout_heart_rate` Workouts recorded with the Workout HR feature
	Source string `json:"source"`

	// ISO 8601 formatted local timestamp indicating when the workout started
	StartDatetime time.Time `json:"start_datetime"`
}

// Workouts represents the workout data within a given timeframe.
type Workouts struct {
	Data []Workout `json:"data"`
	// Pagination token
	NextToken *string `json:"next_token,omitempty"`
}

// Workout gets the workout data within a given timeframe.
// If a start and end date are not provided, ie are empty strings, we fall back to Oura's defaults which are:
//
//	start_date: end_date - 1 day
//	end_date: current UTC date
func (c *Client) Workouts(ctx context.Context, start_date, end_date, next_token string) (*Workouts, *http.Response, error) {
	path := parametiseDate("v2/usercollection/workout", start_date, end_date, next_token)
	req, err := c.NewRequest("GET", path, nil)
	if err != nil {
		return nil, nil, err
	}

	var data *Workouts
	resp, err := c.do(ctx, req, &data)
	if err != nil {
		return data, resp, err
	}

	return data, resp, nil
}
