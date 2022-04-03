package oura

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"testing"

	"github.com/stretchr/testify/assert"
)

var sessionCases = []struct {
	name        string
	start_date  string
	end_date    string
	next_token  string
	expectedURL string
	mock        string
}{
	{
		name:        "get sessions without specific dates",
		start_date:  "",
		end_date:    "",
		next_token:  "",
		expectedURL: "/v2/usercollection/session",
		mock: `{
			"data": [
				{
					"day": "2022-01-13",
					"start_datetime": "2022-01-13T17:43:00+00:00",
					"end_datetime": "2022-01-13T17:49:10+00:00",
					"type": "meditation",
					"heart_rate": {
						"interval": 5.0,
						"items": [
							72.6,
							72.7
						],
						"timestamp": "2022-01-13T17:43:00.000+00:00"
					},
					"heart_rate_variability": {
						"interval": 5.0,
						"items": [
							11.0,
							9.0
						],
						"timestamp": "2022-01-13T17:43:00.000+00:00"
					},
					"mood": "great",
					"motion_count": {
						"interval": 5.0,
						"items": [
							0.0,
							43.0,
							44.0
						],
						"timestamp": "2022-01-13T17:43:00.000+00:00"
					}
				}
			],
			"next_token": "12345"
		}`,
	},
	{
		name:        "get sessions with only start date",
		start_date:  "2020-01-20",
		end_date:    "",
		next_token:  "",
		expectedURL: "/v2/usercollection/session?start_date=2020-01-20",
		mock: `{
			"data": [
				{"day": "2020-01-20"},
				{"day": "2020-01-21"},
				{"day": "2020-01-22"},
				{"day": "2020-01-23"}
			],
			"next_token": "12345"
		}`,
	},
	{
		name:        "get sessions with start and end dates",
		start_date:  "2020-01-20",
		end_date:    "2020-01-22",
		next_token:  "",
		expectedURL: "/v2/usercollection/session?end_date=2020-01-22&start_date=2020-01-20",
		mock: `{
			"data": [
				{"day": "2020-01-20"},
				{"day": "2020-01-21"},
				{"day": "2020-01-22"}
			],
			"next_token": "12345"
		}`,
	},
}

func TestSession(t *testing.T) {
	for _, tc := range sessionCases {
		t.Run(tc.name, func(t *testing.T) {
			testSession(t, tc.start_date, tc.end_date, tc.next_token, tc.expectedURL, tc.mock)
		})
	}
}

func testSession(t *testing.T, start_date, end_date, next_token, expectedURL, mock string) {
	client, mux, _, teardown := setup()
	defer teardown()

	mux.HandleFunc("/v2/usercollection/session", func(w http.ResponseWriter, r *http.Request) {
		assert.Equal(t, http.MethodGet, r.Method)
		assert.Equal(t, expectedURL, r.URL.String())
		fmt.Fprint(w, mock)
	})

	got, _, err := client.Session(context.Background(), start_date, end_date, next_token)
	assert.NoError(t, err, "should not return an error")

	want := &Sessions{}
	json.Unmarshal([]byte(mock), want)

	assert.ObjectsAreEqual(want, got)
}
