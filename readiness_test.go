package oura

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"testing"

	"github.com/stretchr/testify/assert"
)

var readinessTestCases = []struct {
	name        string
	start       string
	end         string
	expectedURL string
	mock        string
}{
	{
		name:        "get readiness without specific dates",
		start:       "",
		end:         "",
		expectedURL: "/v1/readiness",
		mock: `{
			"readiness": [{
				"summary_date": "2016-09-03",
				"period_id": 0,
				"score": 62,
				"score_previous_night": 5,
				"score_readiness_balance": 75,
				"score_previous_day": 61,
				"score_activity_balance": 77,
				"score_resting_hr": 98,
				"score_hrv_balance": 90,
				"score_recovery_index": 45,
				"score_temperature": 86,
				"rest_mode_state": 0
			}]
		}`,
	},
	{
		name:        "get readiness with only start date",
		start:       "2020-01-20",
		end:         "",
		expectedURL: "/v1/readiness?start=2020-01-20",
		mock: `{
			"readiness": [{
				"summary_date": "2020-01-20",
				"score": 62
			},
			{
				"summary_date": "2020-01-21",
				"score": 62
			},
			{
				"summary_date": "2020-01-22",
				"score": 62
			},
			{
				"summary_date": "2020-01-23",
				"score": 62
			}]
		}`,
	},
	{
		name:        "get readiness with start and end dates",
		start:       "2020-01-20",
		end:         "2020-01-22",
		expectedURL: "/v1/readiness?end=2020-01-22&start=2020-01-20",
		mock: `{
			"readiness": [{
				"summary_date": "2020-01-20",
				"score": 62
			},
			{
				"summary_date": "2020-01-21",
				"score": 62
			},
			{
				"summary_date": "2020-01-22",
				"score": 62
			}]
		}`,
	},
}

func TestGetReadiness(t *testing.T) {
	for _, tc := range readinessTestCases {
		t.Run(tc.name, func(st *testing.T) {
			testGetReadiness(st, tc.start, tc.end, tc.expectedURL, tc.mock)
		})
	}
}

func testGetReadiness(t *testing.T, start, end, expectedURL, mock string) {
	client, mux, _, teardown := setup()
	defer teardown()

	mux.HandleFunc("/v1/readiness", func(w http.ResponseWriter, r *http.Request) {
		assert.Equal(t, http.MethodGet, r.Method)
		assert.Equal(t, expectedURL, r.URL.String())
		fmt.Fprint(w, mock)
	})

	got, _, err := client.GetReadiness(context.Background(), start, end)
	assert.NoError(t, err, "should not return an error")

	want := &ReadinessSummaries{}
	json.Unmarshal([]byte(mock), want) //nolint:errcheck

	assert.ObjectsAreEqual(want, got)
}
