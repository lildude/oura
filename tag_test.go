package oura

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"testing"

	"github.com/stretchr/testify/assert"
)

var tagCases = []struct {
	name        string
	start_date  string
	end_date    string
	next_token  string
	expectedURL string
	mock        string
}{
	{
		name:        "get tags without specific dates",
		start_date:  "",
		end_date:    "",
		next_token:  "",
		expectedURL: "/v2/usercollection/tag",
		mock: `{
			"data": [{
					"day": "2021-01-01",
					"text": "Need coffee",
					"timestamp": "2021-01-01T01:02:03-08:00",
					"tags":	[
						"tag_generic_nocaffeine"
					]
        }
    	],
   	 "next_token": "12345"
		}`,
	},
	{
		name:        "get tags with only start date",
		start_date:  "2020-01-20",
		end_date:    "",
		next_token:  "",
		expectedURL: "/v2/usercollection/tag?start_date=2020-01-20",
		mock:        `{}`, // we don't care about the response
	},
	{
		name:        "get tags with start and end dates",
		start_date:  "2020-01-20",
		end_date:    "2020-01-22",
		next_token:  "",
		expectedURL: "/v2/usercollection/tag?end_date=2020-01-22&start_date=2020-01-20",
		mock:        `{}`, // we don't care about the response
	},
	{
		name:        "get tags with next token",
		start_date:  "",
		end_date:    "",
		next_token:  "thisisbase64encodedjson",
		expectedURL: "/v2/usercollection/tag?next_token=thisisbase64encodedjson",
		mock:        `{}`, // we don't care about the response
	},
}

func TestTags(t *testing.T) {
	for _, tc := range tagCases {
		t.Run(tc.name, func(t *testing.T) {
			testTags(t, tc.start_date, tc.end_date, tc.next_token, tc.expectedURL, tc.mock)
		})
	}
}

func testTags(t *testing.T, start_date, end_date, next_token, expectedURL, mock string) {
	client, mux, _, teardown := setup()
	defer teardown()

	mux.HandleFunc("/v2/usercollection/tag", func(w http.ResponseWriter, r *http.Request) {
		assert.Equal(t, http.MethodGet, r.Method)
		assert.Equal(t, expectedURL, r.URL.String())
		fmt.Fprint(w, mock)
	})

	got, _, err := client.Tags(context.Background(), start_date, end_date, next_token)
	assert.NoError(t, err, "should not return an error")

	want := &Tags{}
	json.Unmarshal([]byte(mock), want) //nolint:errcheck

	assert.ObjectsAreEqual(want, got)
}
