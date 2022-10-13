package oura

import (
	"context"
	"net/http"
)

// PersonalInfo represents the personal information of the user returned from the Oura API.
type PersonalInfo struct {
	// The user's age. This field is only present if the user consented to the personal access scope or provided this data to Oura. Otherwise, this field will be `null`.
	Age *int `json:"age,omitempty"`

	// The user's biological sex. This field is only present if the user consented to the personal access scope or provided this data to Oura. Otherwise, this field will be `null`.
	BiologicalSex *string `json:"biological_sex,omitempty"`

	// The user's e-mail. This field is only present if the user consented to the email access scope. Otherwise, this field will be `null`.
	Email *string `json:"email,omitempty"`

	// The user's height (in meters). This field is only present if the user consented to the personal access scope or provided this data to Oura. Otherwise, this field will be `null`.
	Height *float32 `json:"height,omitempty"`

	// The user's weight (in kilograms). This field is only present if the user consented to the personal access scope or provided this data to Oura. Otherwise, this field will be `null`.
	Weight *float32 `json:"weight,omitempty"`
}

func (c *Client) PersonalInfo(ctx context.Context) (*PersonalInfo, *http.Response, error) {
	req, err := c.NewRequest("GET", "v2/usercollection/personal_info", nil)
	if err != nil {
		return nil, nil, err
	}

	var data *PersonalInfo
	resp, err := c.do(ctx, req, &data)
	if err != nil {
		return data, resp, err
	}

	return data, resp, nil
}
