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

var workoutCases = []struct {
	name        string
	startDate   string
	endDate     string
	nextToken   string
	expectedURL string
	mock        string
}{
	{
		name:        "get workouts without specific dates",
		startDate:   "",
		endDate:     "",
		nextToken:   "",
		expectedURL: "/v2/usercollection/workout",
		mock:        `testdata/v2/workout.json`,
	},
	{
		name:        "get workouts with only start date",
		startDate:   "2020-01-20",
		endDate:     "",
		nextToken:   "",
		expectedURL: "/v2/usercollection/workout?start_date=2020-01-20",
		mock:        `{}`, // we don't care about the response
	},
	{
		name:        "get workouts with start and end dates",
		startDate:   "2020-01-20",
		endDate:     "2020-01-22",
		nextToken:   "",
		expectedURL: "/v2/usercollection/workout?end_date=2020-01-22&start_date=2020-01-20",
		mock:        `{}`, // we don't care about the response
	},
	{
		name:        "get workouts with next token",
		startDate:   "",
		endDate:     "",
		nextToken:   "thisisbase64encodedjson",
		expectedURL: "/v2/usercollection/workout?next_token=thisisbase64encodedjson",
		mock:        `{}`, // we don't care about the response
	},
}

func TestWorkouts(t *testing.T) {
	for _, tc := range workoutCases {
		t.Run(tc.name, func(t *testing.T) {
			mock := tc.mock
			if strings.HasPrefix(tc.mock, "testdata/") {
				resp, _ := os.ReadFile(tc.mock)
				mock = string(resp)
			}
			testWorkouts(t, tc.startDate, tc.endDate, tc.nextToken, tc.expectedURL, mock)
		})
	}
}

func testWorkouts(t *testing.T, startDate, endDate, nextToken, expectedURL, mock string) {
	client, mux, teardown := setup()
	defer teardown()

	mux.HandleFunc("/v2/usercollection/workout", func(w http.ResponseWriter, r *http.Request) {
		assert.Equal(t, http.MethodGet, r.Method)
		assert.Equal(t, expectedURL, r.URL.String())
		fmt.Fprint(w, mock)
	})

	got, _, err := client.Workouts(context.Background(), startDate, endDate, nextToken)
	assert.NoError(t, err, "should not return an error")

	want := &Workouts{}
	json.Unmarshal([]byte(mock), want)
	assert.ObjectsAreEqual(want, got)
}
