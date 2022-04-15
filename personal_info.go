package oura

import (
	"context"
	"net/http"
)

// PersonalInfo represents the personal information of the user returned from the Oura API.
type PersonalInfo struct {
	Age           int     `json:"age"`
	Weight        float32 `json:"weight"`
	Height        float32 `json:"height"`
	BiologicalSex string  `json:"biological_sex"`
	Email         string  `json:"email"`
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
