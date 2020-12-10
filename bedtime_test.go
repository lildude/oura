package oura

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"testing"

	"github.com/stretchr/testify/assert"
)

var bedtimeTestCases = []struct {
	name        string
	start       string
	end         string
	expectedURL string
	mock        string
}{
	{
		name:        "get bedtime without specific dates",
		start:       "",
		end:         "",
		expectedURL: "",
		mock: `{
			"ideal_bedtimes": [
				{
					"date": "2020-03-17",
					"bedtime_window": {
						"start": -3600,
						"end": 0
					},
					"status": "IDEAL_BEDTIME_AVAILABLE"
				},
				{
					"date": "2020-03-18",
					"bedtime_window": {
						"start": null,
						"end": null
					},
					"status": "LOW_SLEEP_SCORES"
				}
			]
		}`,
	},
	{
		name:        "get bedtime with only start date",
		start:       "2020-01-20",
		end:         "",
		expectedURL: "/bedtime?start=2020-01-20",
		mock: `{
			"ideal_bedtimes": [
				{
					"date": "2020-01-20",
					"status": "IDEAL_BEDTIME_AVAILABLE"
				},
				{
					"date": "2020-01-21",
					"status": "LOW_SLEEP_SCORES"
				},
				{
					"date": "2020-01-22",
					"status": "IDEAL_BEDTIME_AVAILABLE"
				},
				{
					"date": "2020-01-23",
					"status": "LOW_SLEEP_SCORES"
				},
			]
		}`,
	},
	{
		name:        "get bedtime with start and end dates",
		start:       "2020-01-20",
		end:         "2020-01-22",
		expectedURL: "/bedtime?end=2020-01-22&start=2020-01-20",
		mock: `{
			"ideal_bedtimes": [
				{
					"date": "2020-01-20",
					"status": "IDEAL_BEDTIME_AVAILABLE"
				},
				{
					"date": "2020-01-21",
					"status": "LOW_SLEEP_SCORES"
				},
				{
					"date": "2020-01-22",
					"status": "IDEAL_BEDTIME_AVAILABLE"
				}
			]
		}`,
	},
}

func testGetBedtime(t *testing.T, start, end, expectedURL, mock string) {
	client, mux, _, teardown := setup()
	defer teardown()

	mux.HandleFunc("/bedtime", func(w http.ResponseWriter, r *http.Request) {
		assert.Equal(t, http.MethodGet, r.Method)
		assert.Equal(t, expectedURL, r.URL.String())
		fmt.Fprint(w, mock)
	})

	got, _, err := client.GetBedtime(context.Background(), start, end)
	assert.NoError(t, err, "should not return an error")

	want := &IdealBedtimes{}
	json.Unmarshal([]byte(mock), want)

	assert.ObjectsAreEqual(want, got)
}
