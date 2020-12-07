/*
Package oura provides a client for using the Oura API.

Usage:

	import "github.com/lildude/oura"

Construct a new Oura client, then call various methods on the API to access
different functions of the Oura API. For example:

	client := oura.NewClient(nil)

	// retrieve the user information for the current user
	user, _, err := client.UserInfo(ctx, nil)

All of the API calls will require you to pass in an access token:

	ts := oauth2.StaticTokenSource(
		&oauth2.Token{AccessToken: "TOKEN"},
	)
	ctx := context.Background()
	tc := oauth2.NewClient(ctx, ts)

	client := oura.NewClient(tc)

The Oura API documentation is available at https://cloud.ouraring.com/docs.

*/
package oura