package oura

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"net/url"
	"strings"

	"github.com/pkg/errors"
)

var (
	// BaseURLV1 is Oura's v1 API endpoint
	BaseURLV1 = "https://api.ouraring.com/v1/"
	version = "dev"
	userAgent = fmt.Sprintf("go-oura/%s", version)
)

// Client holds configuration items for the Oura client and provides methods that interact with the Oura API.
type Client struct {
	baseURL *url.URL

	userAgent string
	client    *http.Client
}

// NewClient returns a new Oura API client. If a nil httpClient is
// provided, http.DefaultClient will be used. To use API methods which require
// authentication, provide an http.Client that will perform the authentication
// for you (such as that provided by the golang.org/x/oauth2 library).
// Inspiration: https://github.com/google/go-github/blob/master/github/github.go
func NewClient(cc *http.Client, appName string) *Client {
	if cc == nil {
		cc = http.DefaultClient
	}
	baseURL, _ := url.Parse(BaseURLV1)
	ua := fmt.Sprintf("%s (%s)", appName, userAgent)

	c := &Client{baseURL: baseURL, userAgent: ua, client: cc}
	return c
}

// NewRequest creates an HTTP Request. The client baseURL is checked to confirm that it has a trailing
// slash. A relative URL should be provided without the leading slash. If a non-nil body is provided
// it will be JSON encoded and included in the request.
// Inspiration: https://github.com/google/go-github/blob/master/github/github.go
func (c *Client) NewRequest(method, urlStr string, body interface{}) (*http.Request, error) {
	if !strings.HasSuffix(c.baseURL.Path, "/") {
		return nil, fmt.Errorf("client baseURL does not have a trailing slash: %q", c.baseURL)
	}

	u, err := c.baseURL.Parse(urlStr)
	if err != nil {
		return nil, err
	}

	var buf io.ReadWriter
	if body != nil {
		buf = new(bytes.Buffer)
		enc := json.NewEncoder(buf)
		enc.SetEscapeHTML(false)
		err := enc.Encode(body)
		if err != nil {
			return nil, err
		}
	}

	req, err := http.NewRequest(method, u.String(), buf)
	if err != nil {
		return nil, err
	}

	if body != nil {
		req.Header.Set("Content-Type", "application/json")
	}
	req.Header.Set("Accept", "application/json")
	if c.userAgent != "" {
		req.Header.Set("User-Agent", c.userAgent)
	}
	return req, nil
}

// Do sends a request and returns the response. An error is returned if the request cannot
// be sent or if the API returns an error. If a response is received, the body response body
// is decoded and stored in the value pointed to by v.
// Inspiration: https://github.com/google/go-github/blob/master/github/github.go
func (c *Client) Do(ctx context.Context, req *http.Request, v interface{}) (*http.Response, error) {
	req = req.WithContext(ctx)
	resp, err := c.client.Do(req)

	if err != nil {
		select {
		case <-ctx.Done():
			return nil, errors.Wrap(err, ctx.Err().Error())
		default:
			return nil, err
		}
	}

	data, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, errors.Wrap(err, "unable to read body")
	}
	resp.Body.Close()

	// Anything other than a HTTP 2xx response code is treated as an error. But the structure of error
	// responses differs depending on the API being called. Some APIs return validation errors as part
	// of the standard response. Others respond with a standardised error structure.
	if c := resp.StatusCode; c >= 300 {

		// Handle auth errors
		if resp.StatusCode == http.StatusUnauthorized || resp.StatusCode == http.StatusForbidden {
			err := AuthError(http.StatusText(resp.StatusCode))
			return resp, err
		}

		// Try parsing the response using the standard error schema. If this fails we wrap the parsing
		// error and return. Otherwise return the errors included in the API response payload.
		var e = Errors{}
		err := json.Unmarshal(data, &e)
		if err != nil {
			err = errors.Wrap(err, http.StatusText(resp.StatusCode))
			return resp, errors.Wrap(err, "unable to parse API error response")
		}

		if len(e) != 0 {
			return resp, errors.Wrap(e, http.StatusText(resp.StatusCode))
		}

		// In some cases, the error response is returned as part of the
		// requested resource. In these cases we attempt to decode the
		// resource and return the error.
		err = json.Unmarshal(data, v)
		if err != nil {
			err = errors.Wrap(err, http.StatusText(resp.StatusCode))
			return resp, errors.Wrap(err, "unable to parse API response")
		}

		err = errors.New("no additional error information available")
		return resp, errors.Wrap(err, http.StatusText(resp.StatusCode))
	}

	if v != nil && len(data) != 0 {
		err = json.Unmarshal(data, v)

		switch err {
		case nil:
		case io.EOF:
			err = nil
		default:
			err = errors.Wrap(err, "unable to parse API response")
		}
	}

	return resp, err
}
