package oura

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"testing"

	"github.com/stretchr/testify/assert"
)

var dailyActivityCases = []struct {
	name        string
	start_date  string
	end_date    string
	next_token  string
	expectedURL string
	mock        string
}{
	{
		name:        "get daily activity without specific dates",
		start_date:  "",
		end_date:    "",
		next_token:  "",
		expectedURL: "/v2/usercollection/daily_activity",
		mock: `{
			"data": [{
				"class_5_min": "1111222233334445555",
				"score": 89,
				"active_calories": 897,
				"average_met_minutes": 1.71875,
				"contributors": {
					"meet_daily_targets": 100,
					"move_every_hour": 100,
					"recovery_time": 69,
					"stay_active": 86,
					"training_frequency": 100,
					"training_volume": 100
				},
				"equivalent_walking_distance": 15607,
				"high_activity_met_minutes": 58,
				"high_activity_time": 480,
				"inactivity_alerts": 0,
				"low_activity_met_minutes": 144,
				"low_activity_time": 17400,
				"medium_activity_met_minutes": 525,
				"medium_activity_time": 7380,
				"met": {
					"interval": 60.0,
					"items": [
						1.2,
						10.9
					],
					"timestamp": "2022-04-02T04:00:00.000+01:00"
				},
				"meters_to_target": -11500,
				"non_wear_time": 0,
				"resting_time": 34560,
				"sedentary_met_minutes": 9,
				"sedentary_time": 26580,
				"steps": 12375,
				"target_calories": 300,
				"target_meters": 6000,
				"total_calories": 3012,
				"day": "2022-04-02",
				"timestamp": "2022-04-02T04:00:00+01:00"
			}
		],
		"next_token": null
		}`,
	},
	{
		name:        "get daily activity with only start date",
		start_date:  "2020-01-20",
		end_date:    "",
		next_token:  "",
		expectedURL: "/v2/usercollection/daily_activity?start_date=2020-01-20",
		mock:        `{}`, // we don't care about the response here
	},
	{
		name:        "get daily activity with start and end dates",
		start_date:  "2020-01-20",
		end_date:    "2020-01-22",
		next_token:  "",
		expectedURL: "/v2/usercollection/daily_activity?end_date=2020-01-22&start_date=2020-01-20",
		mock:        `{}`, // we don't care about the response here
	},
	{
		name:        "get daily activity with next token",
		start_date:  "",
		end_date:    "",
		next_token:  "thisisbase64encodedjson",
		expectedURL: "/v2/usercollection/daily_activity?next_token=thisisbase64encodedjson",
		mock:        `{}`, // We don't care about the response here
	},
}

func TestDailyActivities(t *testing.T) {
	for _, tc := range dailyActivityCases {
		t.Run(tc.name, func(t *testing.T) {
			testDailyActivities(t, tc.start_date, tc.end_date, tc.next_token, tc.expectedURL, tc.mock)
		})
	}
}

func testDailyActivities(t *testing.T, start_date, end_date, next_token, expectedURL, mock string) {
	client, mux, _, teardown := setup()
	defer teardown()

	mux.HandleFunc("/v2/usercollection/daily_activity", func(w http.ResponseWriter, r *http.Request) {
		assert.Equal(t, http.MethodGet, r.Method)
		assert.Equal(t, expectedURL, r.URL.String())
		fmt.Fprint(w, mock)
	})

	got, _, err := client.DailyActivities(context.Background(), start_date, end_date, next_token)
	assert.NoError(t, err, "should not return an error")

	want := &DailyActivities{}
	json.Unmarshal([]byte(mock), want) //nolint:errcheck

	assert.ObjectsAreEqual(want, got)
}
