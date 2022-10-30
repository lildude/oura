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

var sleepTestCases = []struct {
	name        string
	start       string
	end         string
	expectedURL string
	mock        string
}{
	{
		name:        "get sleep without specific dates",
		start:       "",
		end:         "",
		expectedURL: "/v1/sleep",
		mock:        `testdata/v1/sleep.json`,
	},
	{
		name:        "get sleep with only start date",
		start:       "2020-01-20",
		end:         "",
		expectedURL: "/v1/sleep?start=2020-01-20",
		mock: `{
			"sleep": [
				{
					"summary_date": "2020-01-20",
					"duration": 21540
				},
				{
					"summary_date": "2020-01-21",
					"duration": 21541
				},
				{
					"summary_date": "2020-01-22",
					"duration": 21541
				},
				{
					"summary_date": "2020-01-23",
					"duration": 21541
				}
			]
		}`,
	},
	{
		name:        "get sleep with start and end dates",
		start:       "2020-01-20",
		end:         "2020-01-22",
		expectedURL: "/v1/sleep?end=2020-01-22&start=2020-01-20",
		mock: `{
			"sleep": [
				{
					"summary_date": "2020-01-20",
					"duration": 21540
				},
				{
					"summary_date": "2020-01-21",
					"duration": 21541
				},
				{
					"summary_date": "2020-01-22",
					"duration": 21541
				}
			]
		}`,
	},
}

func TestGetSleep(t *testing.T) {
	for _, tc := range sleepTestCases {
		t.Run(tc.name, func(st *testing.T) {
			mock := tc.mock
			if strings.HasPrefix(tc.mock, "testdata/") {
				resp, _ := os.ReadFile(tc.mock)
				mock = string(resp)
			}
			testGetSleep(st, tc.start, tc.end, tc.expectedURL, mock)
		})
	}
}

func testGetSleep(t *testing.T, start, end, expectedURL, mock string) {
	client, mux, teardown := setup()
	defer teardown()

	mux.HandleFunc("/v1/sleep", func(w http.ResponseWriter, r *http.Request) {
		assert.Equal(t, http.MethodGet, r.Method)
		assert.Equal(t, expectedURL, r.URL.String())
		fmt.Fprint(w, mock)
	})

	got, _, err := client.GetSleep(context.Background(), start, end)
	assert.NoError(t, err, "should not return an error")

	want := &Sleeps{}
	json.Unmarshal([]byte(mock), want)

	assert.ObjectsAreEqual(want, got)
}

var sleepCases = []struct {
	name        string
	startDate   string
	endDate     string
	nextToken   string
	expectedURL string
	mock        string
}{
	{
		name:        "get sleep without specific dates",
		startDate:   "",
		endDate:     "",
		nextToken:   "",
		expectedURL: "/v2/usercollection/sleep",
		mock:        "testdata/v2/sleep.json",
	},
	{
		name:        "get sleep with only start date",
		startDate:   "2020-01-20",
		endDate:     "",
		nextToken:   "",
		expectedURL: "/v2/usercollection/sleep?start_date=2020-01-20",
		mock:        `{}`, // We don't care about the response here
	},
	{
		name:        "get sleep with start and end dates",
		startDate:   "2020-01-20",
		endDate:     "2020-01-22",
		nextToken:   "",
		expectedURL: "/v2/usercollection/sleep?end_date=2020-01-22&start_date=2020-01-20",
		mock:        `{}`, // We don't care about the response here
	},
	{
		name:        "get sleep with next token",
		startDate:   "",
		endDate:     "",
		nextToken:   "thisisbase64encodedjson",
		expectedURL: "/v2/usercollection/sleep?next_token=thisisbase64encodedjson",
		mock:        `{}`, // We don't care about the response here
	},
	{
		name:        "get error when dates the wrong way round",
		startDate:   "2021-10-01",
		endDate:     "2021-01-01",
		nextToken:   "",
		expectedURL: "/v2/usercollection/sleep?end_date=2021-01-01&start_date=2021-10-01",
		mock: `{
			"detail": "Start time is greater than end time: [start_time: 2021-10-01 01:02:03+00:00; end_date: 2021-01-01 01:02:03+00:00"
		}`,
	},
}

func TestSleeps(t *testing.T) {
	for _, tc := range sleepCases {
		t.Run(tc.name, func(st *testing.T) {
			mock := tc.mock
			if strings.HasPrefix(tc.mock, "testdata/") {
				resp, _ := os.ReadFile(tc.mock)
				mock = string(resp)
			}
			testSleeps(st, tc.startDate, tc.endDate, tc.nextToken, tc.expectedURL, mock)
		})
	}
}

func testSleeps(t *testing.T, startDate, endDate, nextToken, expectedURL, mock string) {
	client, mux, teardown := setup()
	defer teardown()

	mux.HandleFunc("/v2/usercollection/sleep", func(w http.ResponseWriter, r *http.Request) {
		assert.Equal(t, http.MethodGet, r.Method)
		assert.Equal(t, expectedURL, r.URL.String())
		fmt.Fprint(w, mock)
	})

	got, _, err := client.Sleeps(context.Background(), startDate, endDate, nextToken)
	assert.NoError(t, err, "should not return an error")

	want := &Sleeps{}
	json.Unmarshal([]byte(mock), want)

	assert.ObjectsAreEqual(want, got)
}
