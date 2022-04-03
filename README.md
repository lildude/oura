# Oura

![Tests Status Badge](https://github.com/lildude/oura/workflows/Tests/badge.svg)

An unofficial Go client for the [Oura Cloud API v1](https://cloud.ouraring.com/docs/) and [Oura Cloud API v2](https://cloud.ouraring.com/v2/docs/).

## Installation

Use Go to fetch the latest version of the package.

```shell
go get -u 'github.com/lildude/oura'
```

## Usage

Depending on your requirements, you will need an access token to query the API. This can be a personal access token or a full OAuth2 authenticated access token.

See the section on Authentication in the [Oura Cloud API Docs](https://cloud.ouraring.com/v2/docs) for more information the authentication methods.

The simplest approach for accessing your own data is to use a personal access token like this:

```go
package main

import (
  "context"
  "fmt"
  "os"

  "github.com/joho/godotenv"
  "github.com/lildude/oura"
  "golang.org/x/oauth2"
)

func main() {
  godotenv.Load(".env")
  ts := oauth2.StaticTokenSource(&oauth2.Token{AccessToken: os.Getenv("OURA_ACCESS_TOKEN")})
  ctx := context.Background()
  tc := oauth2.NewClient(ctx, ts)

  cl := oura.NewClient(tc)

  info, _, err := cl.PersonalInfo(ctx)
  if err != nil {
    fmt.Println(err)
  }
  fmt.Println(info.Age, info.Gender, info.Weight, info.Email)
}
```

This library supports both v1 and v2 of the Oura API. Function names are in the plural form, where appropriate, with the v1 API calls prefixed with `Get`. For example, `GetActivities` queries the v1 API, and `DailyActivities` queries the v2 API. `GetUserInfo` queries the v1 API and `PersonalInfo` queries the v2 API.

## Releasing

This project uses [GoReleaser](https://goreleaser.com) via GitHub Actions to make the releases quick and easy. When I'm ready for a new release, I push a new tag and the workflow takes care of things.

