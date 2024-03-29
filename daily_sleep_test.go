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

var dailySleepCases = []struct {
	name        string
	startDate   string
	endDate     string
	nextToken   string
	expectedURL string
	mock        string
}{
	{
		name:        "get daily sleep without specific dates",
		startDate:   "",
		endDate:     "",
		nextToken:   "",
		expectedURL: "/v2/usercollection/daily_sleep",
		mock:        `testdata/v2/daily_sleep.json`,
	},
	{
		name:        "get daily sleep with only start date",
		startDate:   "2020-01-20",
		endDate:     "",
		nextToken:   "",
		expectedURL: "/v2/usercollection/daily_sleep?start_date=2020-01-20",
		mock:        `{}`, // we don't care about the response here
	},
	{
		name:        "get daily sleep with start and end dates",
		startDate:   "2020-01-20",
		endDate:     "2020-01-22",
		nextToken:   "",
		expectedURL: "/v2/usercollection/daily_sleep?end_date=2020-01-22&start_date=2020-01-20",
		mock:        `{}`, // we don't care about the response here
	},
	{
		name:        "get daily sleep with next token",
		startDate:   "",
		endDate:     "",
		nextToken:   "thisisbase64encodedjson",
		expectedURL: "/v2/usercollection/daily_sleep?next_token=thisisbase64encodedjson",
		mock:        `{}`, // We don't care about the response here
	},
	{
		name:        "get daily sleep with start and end dates and next token",
		startDate:   "2020-01-20",
		endDate:     "2020-01-22",
		nextToken:   "thisisbase64encodedjson",
		expectedURL: "/v2/usercollection/daily_sleep?end_date=2020-01-22&next_token=thisisbase64encodedjson&start_date=2020-01-20",
		mock:        `{}`, // We don't care about the response here
	},
	{
		name:        "get error with dates the wrong way round",
		startDate:   "2021-10-01",
		endDate:     "2021-01-01",
		nextToken:   "",
		expectedURL: "/v2/usercollection/daily_sleep?end_date=2021-01-01&start_date=2021-10-01",
		mock: `{
			"detail": "Start time is greater than end time: [start_time: 2021-10-01 01:02:03+00:00; end_date: 2021-01-01 01:02:03+00:00"
		}`,
	},
}

func TestDailySleep(t *testing.T) {
	for _, tc := range dailySleepCases {
		t.Run(tc.name, func(t *testing.T) {
			mock := tc.mock
			if strings.HasPrefix(tc.mock, "testdata/") {
				resp, _ := os.ReadFile(tc.mock)
				mock = string(resp)
			}
			testDailySleeps(t, tc.startDate, tc.endDate, tc.nextToken, tc.expectedURL, mock)
		})
	}
}

func testDailySleeps(t *testing.T, startDate, endDate, nextToken, expectedURL, mock string) {
	client, mux, teardown := setup()
	defer teardown()

	mux.HandleFunc("/v2/usercollection/daily_sleep", func(w http.ResponseWriter, r *http.Request) {
		assert.Equal(t, http.MethodGet, r.Method)
		assert.Equal(t, expectedURL, r.URL.String())
		fmt.Fprint(w, mock)
	})

	got, _, err := client.DailySleeps(context.Background(), startDate, endDate, nextToken)
	assert.NoError(t, err, "should not return an error")

	want := &DailySleeps{}
	json.Unmarshal([]byte(mock), want)

	assert.ObjectsAreEqual(want, got)
}
