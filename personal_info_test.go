package oura

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"testing"

	"github.com/stretchr/testify/assert"
)

var personalInfoTestCases = []struct {
	name     string
	mock     string
	expected PersonalInfo
}{
	{
		name: "Regular info",
		mock: `{
			"age": 27,
			"weight": 80,
			"height": 180,
			"biological_sex": "male",
			"email": "john.doe@the.domain"
		}`,
	},
	{
		name: "Info w/ weight & height as float",
		mock: `{
			"age": 27,
			"weight": 80.0,
			"height": 180.0,
			"biological_sex": "male",
			"email": "john.doe@the.domain"
		}`,
	},
}

func TestPersonalInfo(t *testing.T) {
	for _, tc := range personalInfoTestCases {
		client, mux, teardown := setup()
		defer teardown() //nolint:gocritic // We're iterating over test cases, so we can't use t.Cleanup

		mux.HandleFunc("/v2/usercollection/personal_info", func(w http.ResponseWriter, r *http.Request) {
			assert.Equal(t, http.MethodGet, r.Method)
			fmt.Fprint(w, tc.mock)
		})

		got, _, err := client.PersonalInfo(context.Background())
		assert.NoError(t, err, tc.name+" should not return an error")

		want := &PersonalInfo{}
		json.Unmarshal([]byte(tc.mock), want)

		assert.ObjectsAreEqual(tc.expected, got)
	}
}
