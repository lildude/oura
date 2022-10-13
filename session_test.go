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

var sessionCases = []struct {
	name        string
	startDate   string
	endDate     string
	nextToken   string
	expectedURL string
	mock        string
}{
	{
		name:        "get sessions without specific dates",
		startDate:   "",
		endDate:     "",
		nextToken:   "",
		expectedURL: "/v2/usercollection/session",
		mock:        `testdata/v2/session.json`,
	},
	{
		name:        "get sessions with only start date",
		startDate:   "2020-01-20",
		endDate:     "",
		nextToken:   "",
		expectedURL: "/v2/usercollection/session?start_date=2020-01-20",
		mock:        `{}`, // we don't care about the response
	},
	{
		name:        "get sessions with start and end dates",
		startDate:   "2020-01-20",
		endDate:     "2020-01-22",
		nextToken:   "",
		expectedURL: "/v2/usercollection/session?end_date=2020-01-22&start_date=2020-01-20",
		mock:        `{}`, // we don't care about the response
	},
	{
		name:        "get sessions with next token",
		startDate:   "",
		endDate:     "",
		nextToken:   "thisisbase64encodedjson",
		expectedURL: "/v2/usercollection/session?next_token=thisisbase64encodedjson",
		mock:        `{}`, // we don't care about the response
	},
}

func TestSessions(t *testing.T) {
	for _, tc := range sessionCases {
		t.Run(tc.name, func(t *testing.T) {
			mock := tc.mock
			if strings.HasPrefix(tc.mock, "testdata/") {
				resp, _ := os.ReadFile(tc.mock)
				mock = string(resp)
			}
			testSessions(t, tc.startDate, tc.endDate, tc.nextToken, tc.expectedURL, mock)
		})
	}
}

func testSessions(t *testing.T, startDate, endDate, nextToken, expectedURL, mock string) {
	client, mux, teardown := setup()
	defer teardown()

	mux.HandleFunc("/v2/usercollection/session", func(w http.ResponseWriter, r *http.Request) {
		assert.Equal(t, http.MethodGet, r.Method)
		assert.Equal(t, expectedURL, r.URL.String())
		fmt.Fprint(w, mock)
	})

	got, _, err := client.Sessions(context.Background(), startDate, endDate, nextToken)
	assert.NoError(t, err, "should not return an error")

	want := &Sessions{}
	json.Unmarshal([]byte(mock), want)

	assert.ObjectsAreEqual(want, got)
}
