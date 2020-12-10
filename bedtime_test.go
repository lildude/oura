package oura

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestBedtime(t *testing.T) {
	mock := `{
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
	}`

	t.Run("get bedtime without specific dates", func(tc *testing.T) {
		client, mux, _, teardown := setup()
		defer teardown()

		mux.HandleFunc("/bedtime", func(w http.ResponseWriter, r *http.Request) {
			assert.Equal(tc, http.MethodGet, r.Method)
			fmt.Fprint(w, mock)
		})

		got, _, err := client.GetBedtime(context.Background(), "", "")
		assert.NoError(t, err, "should not return an error")

		want := &IdealBedtimes{}
		json.Unmarshal([]byte(mock), want)

		assert.ObjectsAreEqual(want, got)
	})

	t.Run("get bedtime with only start date", func(tc *testing.T) {
		client, mux, _, teardown := setup()
		defer teardown()

		mux.HandleFunc("/bedtime", func(w http.ResponseWriter, r *http.Request) {
			assert.Equal(tc, http.MethodGet, r.Method)
			assert.Equal(tc, "/bedtime?start=2020-01-20", r.URL.String())
			fmt.Fprint(w, mock)
		})

		_, _, err := client.GetBedtime(context.Background(), "2020-01-20", "")
		assert.NoError(t, err, "should not return an error")
	})

	t.Run("get bedtime with start and end dates", func(tc *testing.T) {
		client, mux, _, teardown := setup()
		defer teardown()

		mux.HandleFunc("/bedtime", func(w http.ResponseWriter, r *http.Request) {
			assert.Equal(tc, http.MethodGet, r.Method)
			assert.Equal(tc, "/bedtime?end=2020-01-25&start=2020-01-20", r.URL.String())
			fmt.Fprint(w, mock)
		})

		_, _, err := client.GetBedtime(context.Background(), "2020-01-20", "2020-01-25")
		assert.NoError(t, err, "should not return an error")
	})
}
