package oura

import (
	"context"
	"net/http"
)

// UserInfo is the information for the current user
type UserInfo struct {
	Age    int     `json:"age"`
	Email  string  `json:"email"`
	Gender string  `json:"gender"`
	Height float64 `json:"height"`
	Weight float64 `json:"weight"`
}

// GetUserInfo returns the user information for the current user
func (c *Client) GetUserInfo(ctx context.Context) (*UserInfo, *http.Response, error) {
	req, err := c.NewRequest(ctx, "GET", "v1/userinfo", nil)
	if err != nil {
		return nil, nil, err
	}

	var userInfo *UserInfo
	resp, err := c.do(req, &userInfo)
	if err != nil {
		return userInfo, resp, err
	}

	return userInfo, resp, nil
}
