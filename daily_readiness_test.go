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

var dailyReadinessCases = []struct {
	name        string
	startDate   string
	endDate     string
	nextToken   string
	expectedURL string
	mock        string
}{
	{
		name:        "get daily readiness without specific dates",
		startDate:   "",
		endDate:     "",
		nextToken:   "",
		expectedURL: "/v2/usercollection/daily_readiness",
		mock:        `testdata/v2/daily_readiness.json`,
	},
	{
		name:        "get daily readiness with only start date",
		startDate:   "2020-01-20",
		endDate:     "",
		nextToken:   "",
		expectedURL: "/v2/usercollection/daily_readiness?start_date=2020-01-20",
		mock:        `{}`, // we don't care about the response here
	},
	{
		name:        "get daily readiness with start and end dates",
		startDate:   "2020-01-20",
		endDate:     "2020-01-22",
		nextToken:   "",
		expectedURL: "/v2/usercollection/daily_readiness?end_date=2020-01-22&start_date=2020-01-20",
		mock:        `{}`, // we don't care about the response here
	},
	{
		name:        "get daily readiness with next token",
		startDate:   "",
		endDate:     "",
		nextToken:   "thisisbase64encodedjson",
		expectedURL: "/v2/usercollection/daily_readiness?next_token=thisisbase64encodedjson",
		mock:        `{}`, // We don't care about the response here
	},
	{
		name:        "get error when dates the wrong way round",
		startDate:   "2021-10-01",
		endDate:     "2021-01-01",
		nextToken:   "",
		expectedURL: "/v2/usercollection/daily_readiness?end_date=2021-01-01&start_date=2021-10-01",
		mock: `{
			"detail": "Start time is greater than end time: [start_time: 2021-10-01 01:02:03+00:00; end_date: 2021-01-01 01:02:03+00:00"
		}`,
	},
}

func TestDailyReadiness(t *testing.T) {
	for _, tc := range dailyReadinessCases {
		t.Run(tc.name, func(t *testing.T) {
			mock := tc.mock
			if strings.HasPrefix(tc.mock, "testdata/") {
				resp, _ := os.ReadFile(tc.mock)
				mock = string(resp)
			}
			testDailyReadinesses(t, tc.startDate, tc.endDate, tc.nextToken, tc.expectedURL, mock)
		})
	}
}

func testDailyReadinesses(t *testing.T, startDate, endDate, nextToken, expectedURL, mock string) {
	client, mux, teardown := setup()
	defer teardown()

	mux.HandleFunc("/v2/usercollection/daily_readiness", func(w http.ResponseWriter, r *http.Request) {
		assert.Equal(t, http.MethodGet, r.Method)
		assert.Equal(t, expectedURL, r.URL.String())
		fmt.Fprint(w, mock)
	})

	got, _, err := client.DailyReadinesses(context.Background(), startDate, endDate, nextToken)
	assert.NoError(t, err, "should not return an error")

	want := &DailyReadinesses{}
	json.Unmarshal([]byte(mock), want)

	assert.ObjectsAreEqual(want, got)
}
