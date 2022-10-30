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

var activitiesTestCases = []struct {
	name        string
	start       string
	end         string
	expectedURL string
	mock        string
}{
	{
		name:        "get activity without specific dates",
		start:       "",
		end:         "",
		expectedURL: "/v1/activity",
		mock:        `testdata/v1/activity.json`,
	},
	{
		name:        "get activity with only start date",
		start:       "2020-01-20",
		end:         "",
		expectedURL: "/v1/activity?start=2020-01-20",
		mock: `{
			"activity": [
				{"summary_date": "2020-01-20"},
				{"summary_date": "2020-01-21"},
				{"summary_date": "2020-01-22"},
				{"summary_date": "2020-01-23"}
			]
		}`,
	},
	{
		name:        "get activity with start and end dates",
		start:       "2020-01-20",
		end:         "2020-01-22",
		expectedURL: "/v1/activity?end=2020-01-22&start=2020-01-20",
		mock: `{
			"activity": [
				{"summary_date": "2020-01-20"},
				{"summary_date": "2020-01-21"},
				{"summary_date": "2020-01-22"}
			]
		}`,
	},
}

func TestGetActivities(t *testing.T) {
	for _, tc := range activitiesTestCases {
		t.Run(tc.name, func(st *testing.T) {
			mock := tc.mock
			if strings.HasPrefix(tc.mock, "testdata/") {
				resp, _ := os.ReadFile(tc.mock)
				mock = string(resp)
			}
			testGetActivities(st, tc.start, tc.end, tc.expectedURL, mock)
		})
	}
}

func testGetActivities(t *testing.T, start, end, expectedURL, mock string) {
	client, mux, teardown := setup()
	defer teardown()

	mux.HandleFunc("/v1/activity", func(w http.ResponseWriter, r *http.Request) {
		assert.Equal(t, http.MethodGet, r.Method)
		assert.Equal(t, expectedURL, r.URL.String())
		fmt.Fprint(w, mock)
	})

	got, _, err := client.GetActivities(context.Background(), start, end)
	assert.NoError(t, err, "should not return an error")

	want := &Activities{}
	json.Unmarshal([]byte(mock), want)

	assert.ObjectsAreEqual(want, got)
}
