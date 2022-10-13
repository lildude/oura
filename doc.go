/*
Package oura provides a client for using the Oura API.

Usage:

	import "github.com/lildude/oura"

Construct a new Oura client, then call various methods on the API to access
different functions of the Oura API. For example:

	client := oura.NewClient(nil)

	// retrieve the user information for the current user using the v1 API
	user, _, err := client.GetUserInfo(ctx, nil)

All of the API calls will require you to pass in an access token. This can
be a personal access token or a full OAuth2 access token:

	ts := oauth2.StaticTokenSource(
		&oauth2.Token{AccessToken: "TOKEN"},
	)
	ctx := context.Background()
	tc := oauth2.NewClient(ctx, ts)

	client := oura.NewClient(tc)
	// retrieve the user information for the current user using the v2 API
	info, _, err := client.PersonalInfo(ctx, nil)

This library supports both v1 and v2 of the Oura API. Function names are in
the plural form, where appropriate, with the v1 API calls prefixed with `Get`.
For example, `GetActivities` queries the v1 API, and `DailyActivities` queries
the v2 API. `GetUserInfo` queries the v1 API and `PersonalInfo` queries the v2 API.

The Oura API documentation is available at https://cloud.ouraring.com/v2/docs.
*/
package oura
