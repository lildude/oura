package oura_test

import (
	"context"
	"fmt"
	"os"

	"github.com/joho/godotenv"
	"github.com/lildude/oura"
	"golang.org/x/oauth2"
)

// Note: the examples listed here are compiled but not executed while testing.
// See the documentation on [Testing](https://golang.org/pkg/testing/#hdr-Examples)
// for further details.

func Example_getUserInfo() {
	godotenv.Load(".env")
	ts := oauth2.StaticTokenSource(&oauth2.Token{AccessToken: os.Getenv("OURA_ACCESS_TOKEN")})
	ctx := context.Background()
	tc := oauth2.NewClient(ctx, ts)

	cl := oura.NewClient(tc)

	userInfo, httpResp, err := cl.GetUserInfo(ctx)
	if err != nil {
		fmt.Println(err)
		fmt.Println(httpResp)
	} else {
		fmt.Println(userInfo.Age, userInfo.Gender, userInfo.Weight, userInfo.Email)
	}
	// Output
}

func Example_getSleep() {
	godotenv.Load(".env")
	ts := oauth2.StaticTokenSource(&oauth2.Token{AccessToken: os.Getenv("OURA_ACCESS_TOKEN")})
	ctx := context.Background()
	tc := oauth2.NewClient(ctx, ts)

	cl := oura.NewClient(tc)

	sleepInfo, httpResp, err := cl.GetSleep(ctx, "2021-12-02", "2021-12-03")
	if err != nil {
		fmt.Println(err)
		fmt.Println(httpResp)
	} else {
		fmt.Println(sleepInfo)
	}
	// Output
}

func Example_getActivities() {
	godotenv.Load(".env")
	ts := oauth2.StaticTokenSource(&oauth2.Token{AccessToken: os.Getenv("OURA_ACCESS_TOKEN")})
	ctx := context.Background()
	tc := oauth2.NewClient(ctx, ts)

	cl := oura.NewClient(tc)

	activityInfo, httpResp, err := cl.GetActivities(ctx, "2021-12-02", "2021-12-03")
	if err != nil {
		fmt.Println(err)
		fmt.Println(httpResp)
	} else {
		fmt.Println(activityInfo)
	}
	// Output
}
