package oura

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"testing"

	"github.com/stretchr/testify/assert"
)

var dailyReadinessCases = []struct {
	name        string
	start_date  string
	end_date    string
	next_token  string
	expectedURL string
	mock        string
}{
	{
		name:        "get daily readiness without specific dates",
		start_date:  "",
		end_date:    "",
		next_token:  "",
		expectedURL: "/v2/usercollection/daily_readiness",
		mock: `{
			"data": [{
				"contributors": {
					"activity_balance": 56,
					"body_temperature": 98,
					"hrv_balance": 75,
					"previous_day_activity": null,
					"previous_night": 35,
					"recovery_index": 47,
					"resting_heart_rate": 94,
					"sleep_balance": 73
				},
				"day": "2021-10-27",
				"score": 66,
				"temperature_deviation": -0.2,
				"temperature_trend_deviation": 0.1,
				"timestamp": "2021-10-27T00:00:00+00:00"
			}],
			"next_token": null
		}`,
	},
	{
		name:        "get daily readiness with only start date",
		start_date:  "2020-01-20",
		end_date:    "",
		next_token:  "",
		expectedURL: "/v2/usercollection/daily_readiness?start_date=2020-01-20",
		mock:        `{}`, // we don't care about the response here
	},
	{
		name:        "get daily readiness with start and end dates",
		start_date:  "2020-01-20",
		end_date:    "2020-01-22",
		next_token:  "",
		expectedURL: "/v2/usercollection/daily_readiness?end_date=2020-01-22&start_date=2020-01-20",
		mock:        `{}`, // we don't care about the response here
	},
	{
		name:        "get daily readiness with next token",
		start_date:  "",
		end_date:    "",
		next_token:  "thisisbase64encodedjson",
		expectedURL: "/v2/usercollection/daily_readiness?next_token=thisisbase64encodedjson",
		mock:        `{}`, // We don't care about the response here
	},
	{
		name:        "get error when dates the wrong way round",
		start_date:  "2021-10-01",
		end_date:    "2021-01-01",
		next_token:  "",
		expectedURL: "/v2/usercollection/daily_readiness?end_date=2021-01-01&start_date=2021-10-01",
		mock: `{
			"detail": "Start time is greater than end time: [start_time: 2021-10-01 01:02:03+00:00; end_date: 2021-01-01 01:02:03+00:00"
		}`,
	},
}

func TestDailyReadiness(t *testing.T) {
	for _, tc := range dailyReadinessCases {
		t.Run(tc.name, func(t *testing.T) {
			testDailyReadinesses(t, tc.start_date, tc.end_date, tc.next_token, tc.expectedURL, tc.mock)
		})
	}
}

func testDailyReadinesses(t *testing.T, start_date, end_date, next_token, expectedURL, mock string) {
	client, mux, _, teardown := setup()
	defer teardown()

	mux.HandleFunc("/v2/usercollection/daily_readiness", func(w http.ResponseWriter, r *http.Request) {
		assert.Equal(t, http.MethodGet, r.Method)
		assert.Equal(t, expectedURL, r.URL.String())
		fmt.Fprint(w, mock)
	})

	got, _, err := client.DailyReadinesses(context.Background(), start_date, end_date, next_token)
	assert.NoError(t, err, "should not return an error")

	want := &DailyReadinesses{}
	json.Unmarshal([]byte(mock), want) //nolint:errcheck

	assert.ObjectsAreEqual(want, got)
}
