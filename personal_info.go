package oura

import (
	"context"
	"net/http"
)

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

	var personalInfo *PersonalInfo
	resp, err := c.Do(ctx, req, &personalInfo)
	if err != nil {
		return personalInfo, resp, err
	}

	return personalInfo, resp, nil
}
