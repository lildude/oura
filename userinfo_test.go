package oura

import (
	"context"
	"fmt"
	"net/http"
	"testing"

	"github.com/stretchr/testify/assert"
)

var infoTestCases = []struct {
	name     string
	mock     string
	expected UserInfo
}{
	{
		name: "Regular info",
		mock: `{
			"age": 27,
			"weight": 80,
			"height": 180,
			"gender": "male",
			"email": "john.doe@the.domain"
		}`,
		expected: UserInfo{
			Age:    27,
			Weight: 80.0,
			Height: 180,
			Gender: "male",
			Email:  "john.doe@the.domain",
		},
	},
	{
		name: "Info w/ weight & height as float",
		mock: `{
			"age": 27,
			"weight": 80.0,
			"height": 180.0,
			"gender": "male",
			"email": "john.doe@the.domain"
		}`,
		expected: UserInfo{
			Age:    27,
			Weight: 80.0,
			Height: 180.0,
			Gender: "male",
			Email:  "john.doe@the.domain",
		},
	},
}

func TestUserInfo(t *testing.T) {
	for _, tc := range infoTestCases {
		client, mux, teardown := setup()
		defer teardown() //nolint:gocritic // We're iterating over test cases, so we can't use t.Cleanup

		mux.HandleFunc("/v1/userinfo", func(w http.ResponseWriter, r *http.Request) {
			assert.Equal(t, http.MethodGet, r.Method)
			fmt.Fprint(w, tc.mock)
		})

		got, _, err := client.GetUserInfo(context.Background())
		assert.NoError(t, err, tc.name+" should not return an error")
		assert.ObjectsAreEqual(tc.expected, got)
	}
}
