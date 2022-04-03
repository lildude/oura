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
	start_date  string
	end_date    string
	next_token  string
	expectedURL string
	mock        string
}{
	{
		name:        "get heartrates without specific dates",
		start_date:  "",
		end_date:    "",
		next_token:  "",
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
			"next_token": "123456"
		}`,
	},
	{
		name:        "get heartrates with only start date",
		start_date:  "2020-01-20T00:00:00+00:00",
		end_date:    "",
		next_token:  "",
		expectedURL: "/v2/usercollection/heartrate?start_datetime=2020-01-20T00%3A00%3A00%2B00%3A00",
		mock: `{
			"data": [
				{"timestamp": "2020-01-20T12:38:38+00:00"},
				{"timestamp": "2020-01-21T12:38:43+00:00"},
				{"timestamp": "2020-01-22T12:38:47+00:00"}
			],
			"next_token": "123456"
		}`,
	},
	{
		name:        "get heartrates with start and end dates",
		start_date:  "2020-01-20T00:00:00+00:00",
		end_date:    "2020-01-22T00:00:00+00:00",
		next_token:  "",
		expectedURL: "/v2/usercollection/heartrate?end_datetime=2020-01-22T00%3A00%3A00%2B00%3A00&start_datetime=2020-01-20T00%3A00%3A00%2B00%3A00",
		mock: `{
			"data": [
				{"timestamp": "2020-01-20T12:38:38+00:00"},
				{"timestamp": "2020-01-21T12:38:43+00:00"},
				{"timestamp": "2020-01-22T12:38:47+00:00"}
			],
			"next_token": "123456"
		}`,
	},
}

func TestHeartrate(t *testing.T) {
	for _, tc := range heartrateCases {
		t.Run(tc.name, func(t *testing.T) {
			testHeartrate(t, tc.start_date, tc.end_date, tc.next_token, tc.expectedURL, tc.mock)
		})
	}
}

func testHeartrate(t *testing.T, start_date, end_date, next_token, expectedURL, mock string) {
	client, mux, _, teardown := setup()
	defer teardown()

	mux.HandleFunc("/v2/usercollection/heartrate", func(w http.ResponseWriter, r *http.Request) {
		assert.Equal(t, http.MethodGet, r.Method)
		assert.Equal(t, expectedURL, r.URL.String())
		fmt.Fprint(w, mock)
	})

	got, _, err := client.Heartrate(context.Background(), start_date, end_date, next_token)
	assert.NoError(t, err, "should not return an error")

	want := &Heartrate{}
	json.Unmarshal([]byte(mock), want) //nolint:errcheck

	assert.ObjectsAreEqual(want, got)
}
