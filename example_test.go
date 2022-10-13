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

func Example_getUserInfo() { //nolint:testableexamples // This is an example of how to use the library, not a test
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

func Example_getSleep() { //nolint:testableexamples // This is an example of how to use the library, not a test
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

func Example_getActivities() { //nolint:testableexamples // This is an example of how to use the library, not a test
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

func Example_dailyActivity() { //nolint:testableexamples // This is an example of how to use the library, not a test
	godotenv.Load(".env")
	ts := oauth2.StaticTokenSource(&oauth2.Token{AccessToken: os.Getenv("OURA_ACCESS_TOKEN")})
	ctx := context.Background()
	tc := oauth2.NewClient(ctx, ts)

	cl := oura.NewClient(tc)
	dailyActivity, httpResp, err := cl.DailyActivities(ctx, "2022-03-20", "2022-03-22", "")
	if err != nil {
		fmt.Println(err)
		fmt.Println(httpResp)
	} else {
		fmt.Println(dailyActivity)
	}
	// Output
}

func Example_dailyReadiness() { //nolint:testableexamples // This is an example of how to use the library, not a test
	godotenv.Load(".env")
	ts := oauth2.StaticTokenSource(&oauth2.Token{AccessToken: os.Getenv("OURA_ACCESS_TOKEN")})
	ctx := context.Background()
	tc := oauth2.NewClient(ctx, ts)

	cl := oura.NewClient(tc)
	dailyReadiness, httpResp, err := cl.DailyReadinesses(ctx, "2022-03-20", "2022-03-22", "")
	if err != nil {
		fmt.Println(err)
		fmt.Println(httpResp)
	} else {
		fmt.Println(dailyReadiness)
	}
	// Output
}

func Example_dailySleep() { //nolint:testableexamples // This is an example of how to use the library, not a test
	godotenv.Load(".env")
	ts := oauth2.StaticTokenSource(&oauth2.Token{AccessToken: os.Getenv("OURA_ACCESS_TOKEN")})
	ctx := context.Background()
	tc := oauth2.NewClient(ctx, ts)

	cl := oura.NewClient(tc)
	dailySleeps, httpResp, err := cl.DailySleeps(ctx, "2022-03-20", "2022-03-22", "")
	if err != nil {
		fmt.Println(err)
		fmt.Println(httpResp)
	} else {
		fmt.Println(dailySleeps)
	}
	// Output
}

func Example_heartrate() { //nolint:testableexamples // This is an example of how to use the library, not a test
	godotenv.Load(".env")
	ts := oauth2.StaticTokenSource(&oauth2.Token{AccessToken: os.Getenv("OURA_ACCESS_TOKEN")})
	ctx := context.Background()
	tc := oauth2.NewClient(ctx, ts)

	cl := oura.NewClient(tc)
	heartrate, httpResp, err := cl.Heartrates(ctx, "2022-03-20T00:00:00+00:00", "2022-03-22T00:00:00+00:00", "")
	if err != nil {
		fmt.Println(err)
		fmt.Println(httpResp)
	} else {
		fmt.Println(heartrate)
	}
	// Output
}

func Example_personalInfo() { //nolint:testableexamples // This is an example of how to use the library, not a test
	godotenv.Load(".env")
	ts := oauth2.StaticTokenSource(&oauth2.Token{AccessToken: os.Getenv("OURA_ACCESS_TOKEN")})
	ctx := context.Background()
	tc := oauth2.NewClient(ctx, ts)

	cl := oura.NewClient(tc)
	personalInfo, httpResp, err := cl.PersonalInfo(ctx)
	if err != nil {
		fmt.Println(err)
		fmt.Println(httpResp)
	} else {
		fmt.Println(personalInfo)
	}
	// Output
}

func Example_session() { //nolint:testableexamples // This is an example of how to use the library, not a test
	godotenv.Load(".env")
	ts := oauth2.StaticTokenSource(&oauth2.Token{AccessToken: os.Getenv("OURA_ACCESS_TOKEN")})
	ctx := context.Background()
	tc := oauth2.NewClient(ctx, ts)

	cl := oura.NewClient(tc)
	session, httpResp, err := cl.Sessions(ctx, "2022-03-20", "2022-03-22", "")
	if err != nil {
		fmt.Println(err)
		fmt.Println(httpResp)
	} else {
		fmt.Println(session)
	}
	// Output
}

func Example_sleep() { //nolint:testableexamples // This is an example of how to use the library, not a test
	godotenv.Load(".env")
	ts := oauth2.StaticTokenSource(&oauth2.Token{AccessToken: os.Getenv("OURA_ACCESS_TOKEN")})
	ctx := context.Background()
	tc := oauth2.NewClient(ctx, ts)

	cl := oura.NewClient(tc)
	sleeps, httpResp, err := cl.Sleeps(ctx, "2022-03-20", "2022-03-22", "")
	if err != nil {
		fmt.Println(err)
		fmt.Println(httpResp)
	} else {
		fmt.Println(sleeps)
	}
	// Output
}

func Example_tag() { //nolint:testableexamples // This is an example of how to use the library, not a test
	godotenv.Load(".env")
	ts := oauth2.StaticTokenSource(&oauth2.Token{AccessToken: os.Getenv("OURA_ACCESS_TOKEN")})
	ctx := context.Background()
	tc := oauth2.NewClient(ctx, ts)

	cl := oura.NewClient(tc)
	tags, httpResp, err := cl.Tags(ctx, "2022-03-20", "2022-03-22", "")
	if err != nil {
		fmt.Println(err)
		fmt.Println(httpResp)
	} else {
		fmt.Println(tags)
	}
	// Output
}

func Example_workout() { //nolint:testableexamples // This is an example of how to use the library, not a test
	godotenv.Load(".env")
	ts := oauth2.StaticTokenSource(&oauth2.Token{AccessToken: os.Getenv("OURA_ACCESS_TOKEN")})
	ctx := context.Background()
	tc := oauth2.NewClient(ctx, ts)

	cl := oura.NewClient(tc)
	tags, httpResp, err := cl.Workouts(ctx, "2022-03-20", "2022-03-22", "")
	if err != nil {
		fmt.Println(err)
		fmt.Println(httpResp)
	} else {
		fmt.Println(tags)
	}
	// Output
}
