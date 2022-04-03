package oura

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"testing"

	"github.com/stretchr/testify/assert"
)

var workoutCases = []struct {
	name        string
	start_date  string
	end_date    string
	next_token  string
	expectedURL string
	mock        string
}{
	{
		name:        "get workouts without specific dates",
		start_date:  "",
		end_date:    "",
		next_token:  "",
		expectedURL: "/v2/usercollection/workout",
		mock: `{
			"data": [
				{
					"activity": "walking",
					"calories": 106.206,
					"day": "2022-04-02",
					"distance": "2.3",
					"end_datetime": "2022-04-02T15:12:00+01:00",
					"intensity": "moderate",
					"label": "foo",
					"source": "confirmed",
					"start_datetime": "2022-04-02T14:41:00+01:00"
				},
				{
					"activity": "cycling",
					"calories": 350.784,
					"day": "2022-04-02",
					"distance": "50.2",
					"end_datetime": "2022-04-02T20:36:00+01:00",
					"intensity": "moderate",
					"label": "bar",
					"source": "confirmed",
					"start_datetime": "2022-04-02T19:48:00+01:00"
				}
			],
			"next_token": "12345"
		}`,
	},
	{
		name:        "get workouts with only start date",
		start_date:  "2020-01-20",
		end_date:    "",
		next_token:  "",
		expectedURL: "/v2/usercollection/workout?start_date=2020-01-20",
		mock: `{
			"data": [
				{"timestamp": "2020-01-20T04:00:00+00:00"},
				{"timestamp": "2020-01-21T04:00:00+00:00"},
				{"timestamp": "2020-01-22T04:00:00+00:00"},
				{"timestamp": "2020-01-23T04:00:00+00:00"},
				{"timestamp": "2020-01-24T04:00:00+00:00"},
				{"timestamp": "2020-01-25T04:00:00+00:00"},
				{"timestamp": "2020-01-26T04:00:00+00:00"}
			],
			"next_token": "12345"
		}`,
	},
	{
		name:        "get workouts with start and end dates",
		start_date:  "2020-01-20",
		end_date:    "2020-01-22",
		next_token:  "",
		expectedURL: "/v2/usercollection/workout?end_date=2020-01-22&start_date=2020-01-20",
		mock: `{
			"data": [
				{"timestamp": "2020-01-20T04:00:00+00:00"},
				{"timestamp": "2020-01-21T04:00:00+00:00"},
				{"timestamp": "2020-01-22T04:00:00+00:00"}
			],
			"next_token": "12345"
		}`,
	},
}

func TestWorkout(t *testing.T) {
	for _, tc := range workoutCases {
		t.Run(tc.name, func(t *testing.T) {
			testWorkout(t, tc.start_date, tc.end_date, tc.next_token, tc.expectedURL, tc.mock)
		})
	}
}

func testWorkout(t *testing.T, start_date, end_date, next_token, expectedURL, mock string) {
	client, mux, _, teardown := setup()
	defer teardown()

	mux.HandleFunc("/v2/usercollection/workout", func(w http.ResponseWriter, r *http.Request) {
		assert.Equal(t, http.MethodGet, r.Method)
		assert.Equal(t, expectedURL, r.URL.String())
		fmt.Fprint(w, mock)
	})

	got, _, err := client.Workout(context.Background(), start_date, end_date, next_token)
	assert.NoError(t, err, "should not return an error")

	want := &Sessions{}
	json.Unmarshal([]byte(mock), want)

	assert.ObjectsAreEqual(want, got)
}
