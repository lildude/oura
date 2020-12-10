package oura

import (
	"context"
	"net/http"
)

// UserInfo is the information for the current user
type UserInfo struct {
	Age    int    `json:"age"`
	Weight int    `json:"weight"`
	Gender string `json:"gender"`
	Email  string `json:"email"`
}

// GetUserInfo returns the user information for the current user
func (c *Client) GetUserInfo(ctx context.Context) (*UserInfo, *http.Response, error) {
	req, err := c.NewRequest("GET", "userinfo", nil)
	if err != nil {
		return nil, nil, err
	}

	var userInfo *UserInfo
	resp, err := c.Do(ctx, req, &userInfo)

	if err != nil {
		return userInfo, resp, err
	}

	return userInfo, resp, nil
}
