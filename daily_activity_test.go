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
		mock:        `{}`,
	},
	{
		name:        "get daily activity with only start date",
		start_date:  "2020-01-20",
		end_date:    "",
		next_token:  "",
		expectedURL: "/v2/usercollection/daily_activity?start_date=2020-01-20",
		mock:        `{}`,
	},
	{
		name:        "get daily activity with start and end dates",
		start_date:  "2020-01-20",
		end_date:    "2020-01-22",
		next_token:  "",
		expectedURL: "/v2/usercollection/daily_activity?end_date=2020-01-22&start_date=2020-01-20",
		mock:        `{}`,
	},
}

func TestDailyActivity(t *testing.T) {
	for _, tc := range dailyActivityCases {
		t.Run(tc.name, func(t *testing.T) {
			testDailyActivity(t, tc.start_date, tc.end_date, tc.next_token, tc.expectedURL, tc.mock)
		})
	}
}

func testDailyActivity(t *testing.T, start_date, end_date, next_token, expectedURL, mock string) {
	client, mux, _, teardown := setup()
	defer teardown()

	mux.HandleFunc("/v2/usercollection/daily_activity", func(w http.ResponseWriter, r *http.Request) {
		assert.Equal(t, http.MethodGet, r.Method)
		assert.Equal(t, expectedURL, r.URL.String())
		fmt.Fprint(w, mock)
	})

	got, _, err := client.DailyActivity(context.Background(), start_date, end_date, next_token)
	assert.NoError(t, err, "should not return an error")

	want := &DailyActivities{}
	json.Unmarshal([]byte(mock), want)

	assert.ObjectsAreEqual(want, got)
}
