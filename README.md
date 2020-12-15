# Oura

![Tests Status Badge](https://github.com/lildude/oura/workflows/Tests/badge.svg)

An unofficial Go client for the [Oura Cloud API](https://cloud.ouraring.com/docs/).

## Installation

Use Go to fetch the latest version of the package.

```shell
go get -u 'github.com/lildude/oura'
```

## Usage

Depending on your requirements, you will need an access token to query the API. This can be a personal access token or a full OAuth2 authenticated access token.

See the section on Authentication in the [Oura Cloud API Docs](https://cloud.ouraring.com/docs) for more information the authentication methods.

You will need to provide an application name as a string when initialising the client. This is used in the user agent when querying the API and is used to identify applications that are accessing the API and enable Oura to contact the application author if there are problems. So pick a name that stands out!

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

  cl := oura.NewClient(tc, "My Cool App/3.2.1")

  userInfo, _, _err_ := cl.UserInfo(ctx)
  if err != nil {
    fmt.Println(err)
  }
  fmt.Println(userInfo.Age, userInfo.Gender, userInfo.Weight, userInfo.Email)
}
```
