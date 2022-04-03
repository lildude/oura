package oura

import (
	"context"
	"fmt"
	"net/http"
	"net/url"
	"time"
)

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

type Workouts struct {
	Data      []Tag  `json:"data"`
	NextToken string `json:"next_token"`
}

func (c *Client) Workout(ctx context.Context, start_date, end_date, next_token string) (*Workouts, *http.Response, error) {
	path := "v2/usercollection/workout"
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

	var workouts *Workouts
	resp, err := c.Do(ctx, req, &workouts)
	if err != nil {
		return workouts, resp, err
	}

	return workouts, resp, nil
}
