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

func Example() {
	godotenv.Load(".env")
	ts := oauth2.StaticTokenSource(&oauth2.Token{AccessToken: os.Getenv("OURA_ACCESS_TOKEN")})
	ctx := context.Background()
	tc := oauth2.NewClient(ctx, ts)

	cl := oura.NewClient(tc)

	userInfo, _, _ := cl.UserInfo(ctx)
	fmt.Println(userInfo.Age, userInfo.Gender, userInfo.Weight, userInfo.Email)
}