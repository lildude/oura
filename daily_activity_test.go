package oura

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"os"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

var dailyActivityCases = []struct {
	name        string
	startDate   string
	endDate     string
	nextToken   string
	expectedURL string
	mock        string
}{
	{
		name:        "get daily activity without specific dates",
		startDate:   "",
		endDate:     "",
		nextToken:   "",
		expectedURL: "/v2/usercollection/daily_activity",
		mock:        `testdata/v2/daily_activity.json`,
	},
	{
		name:        "get daily activity with only start date",
		startDate:   "2020-01-20",
		endDate:     "",
		nextToken:   "",
		expectedURL: "/v2/usercollection/daily_activity?start_date=2020-01-20",
		mock:        `{}`, // we don't care about the response here
	},
	{
		name:        "get daily activity with start and end dates",
		startDate:   "2020-01-20",
		endDate:     "2020-01-22",
		nextToken:   "",
		expectedURL: "/v2/usercollection/daily_activity?end_date=2020-01-22&start_date=2020-01-20",
		mock:        `{}`, // we don't care about the response here
	},
	{
		name:        "get daily activity with next token",
		startDate:   "",
		endDate:     "",
		nextToken:   "thisisbase64encodedjson",
		expectedURL: "/v2/usercollection/daily_activity?next_token=thisisbase64encodedjson",
		mock:        `{}`, // We don't care about the response here
	},
	{
		name:        "get error with dates the wrong way round",
		startDate:   "2020-01-25",
		endDate:     "2020-01-22",
		nextToken:   "",
		expectedURL: "/v2/usercollection/daily_activity?end_date=2020-01-22&start_date=2020-01-25",
		mock: `{
			"detail": "Start date is greater than end date: [start_date: 2020-01-25; end_date: 2020-01-22]"
		}`,
	},
}

func TestDailyActivities(t *testing.T) {
	for _, tc := range dailyActivityCases {
		t.Run(tc.name, func(t *testing.T) {
			mock := tc.mock
			if strings.HasPrefix(tc.mock, "testdata/") {
				resp, _ := os.ReadFile(tc.mock)
				mock = string(resp)
			}
			testDailyActivities(t, tc.startDate, tc.endDate, tc.nextToken, tc.expectedURL, mock)
		})
	}
}

func testDailyActivities(t *testing.T, startDate, endDate, nextToken, expectedURL, mock string) {
	client, mux, teardown := setup()
	defer teardown()

	mux.HandleFunc("/v2/usercollection/daily_activity", func(w http.ResponseWriter, r *http.Request) {
		assert.Equal(t, http.MethodGet, r.Method)
		assert.Equal(t, expectedURL, r.URL.String())
		fmt.Fprint(w, mock)
	})

	got, _, err := client.DailyActivities(context.Background(), startDate, endDate, nextToken)
	assert.NoError(t, err, "should not return an error")

	want := &DailyActivities{}
	json.Unmarshal([]byte(mock), want)

	assert.ObjectsAreEqual(want, got)
}
