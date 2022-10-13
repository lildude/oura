package oura

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"testing"

	"github.com/stretchr/testify/assert"
)

var heartrateCases = []struct {
	name        string
	startDate   string
	endDate     string
	nextToken   string
	expectedURL string
	mock        string
}{
	{
		name:        "get heartrates without specific dates",
		startDate:   "",
		endDate:     "",
		nextToken:   "",
		expectedURL: "/v2/usercollection/heartrate",
		mock: `{
			"data": [
				{
					"bpm": 79,
					"source": "awake",
					"timestamp": "2022-04-02T12:38:38+00:00"
				},
				{
					"bpm": 79,
					"source": "awake",
					"timestamp": "2022-04-02T12:38:43+00:00"
				},
				{
					"bpm": 79,
					"source": "awake",
					"timestamp": "2022-04-02T12:38:47+00:00"
				}],
			"nextToken": "123456"
		}`,
	},
	{
		name:        "get heartrates with only start date",
		startDate:   "2020-01-20T00:00:00+00:00",
		endDate:     "",
		nextToken:   "",
		expectedURL: "/v2/usercollection/heartrate?start_datetime=2020-01-20T00%3A00%3A00%2B00%3A00",
		mock:        `{}`, // we don't care about the response
	},
	{
		name:        "get heartrates with start and end dates",
		startDate:   "2020-01-20T00:00:00+00:00",
		endDate:     "2020-01-22T00:00:00+00:00",
		nextToken:   "",
		expectedURL: "/v2/usercollection/heartrate?end_datetime=2020-01-22T00%3A00%3A00%2B00%3A00&start_datetime=2020-01-20T00%3A00%3A00%2B00%3A00",
		mock:        `{}`,
	},
	{
		name:        "get heartrates with next token",
		startDate:   "",
		endDate:     "",
		nextToken:   "thisisbase64encodedjson",
		expectedURL: "/v2/usercollection/heartrate?next_token=thisisbase64encodedjson",
		mock:        `{}`, // we don't care about the response
	},
}

func TestHeartrates(t *testing.T) {
	for _, tc := range heartrateCases {
		t.Run(tc.name, func(t *testing.T) {
			testHeartrates(t, tc.startDate, tc.endDate, tc.nextToken, tc.expectedURL, tc.mock)
		})
	}
}

func testHeartrates(t *testing.T, startDate, endDate, nextToken, expectedURL, mock string) {
	client, mux, teardown := setup()
	defer teardown()

	mux.HandleFunc("/v2/usercollection/heartrate", func(w http.ResponseWriter, r *http.Request) {
		assert.Equal(t, http.MethodGet, r.Method)
		assert.Equal(t, expectedURL, r.URL.String())
		fmt.Fprint(w, mock)
	})

	got, _, err := client.Heartrates(context.Background(), startDate, endDate, nextToken)
	assert.NoError(t, err, "should not return an error")

	want := &Heartrate{}
	json.Unmarshal([]byte(mock), want)

	assert.ObjectsAreEqual(want, got)
}
