package oura

import (
	"context"
	"net/http"
	"time"
)

// Workout represents the data returned from the Oura API for a single workout.
type Workout struct {
	Activity      string    `json:"activity"`
	Calories      float32   `json:"calories"`
	Day           string    `json:"day"`
	Distance      float32   `json:"distance"`
	EndDatetime   time.Time `json:"end_datetime"`
	Intensity     string    `json:"intensity"`
	Label         string    `json:"label"`
	Source        string    `json:"source"`
	StartDatetime time.Time `json:"start_datetime"`
}

// Workouts represents the workout data within a given timeframe.
type Workouts struct {
	Data      []Tag  `json:"data"`
	NextToken string `json:"next_token"`
}

// Workout gets the workout data within a given timeframe.
// If a start and end date are not provided, ie are empty strings, we fall back to Oura's defaults which are:
// 	start_date: end_date - 1 day
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
