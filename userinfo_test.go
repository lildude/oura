package oura

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestUserInfo(t *testing.T) {
	mock := `{
		"age": 27,
		"weight": 80,
		"gender": "male",
		"email": "john.doe@the.domain"
	}`

	client, mux, _, teardown := setup()
	defer teardown()

	mux.HandleFunc("/userinfo", func(w http.ResponseWriter, r *http.Request) {
		assert.Equal(t, http.MethodGet, r.Method)
		fmt.Fprint(w, mock)
	})

	got, _, err := client.UserInfo(context.Background())
	assert.NoError(t, err, "should not return an error")

	want := &UserInfo{}
	json.Unmarshal([]byte(mock), want)

	assert.ObjectsAreEqual(want, got)
}