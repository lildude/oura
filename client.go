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
	"time"

	"github.com/pkg/errors"
)

var (
	BaseURL   = "https://api.ouraring.com/"
	userAgent = "go-oura"
)

// Client holds configuration items for the Oura client and provides methods that interact with the Oura API.
type Client struct {
	baseURL *url.URL

	UserAgent string
	client    *http.Client
}

// TimeSeriesData is time series data used by various other methods.
type TimeSeriesData struct {
	Interval  float32   `json:"interval"`
	Items     []float32 `json:"items"`
	Timestamp time.Time `json:"timestamp"`
}

func (e *ErrorDetail) Error() string {
	return e.Detail
}

// ErrorDetail holds the details of an error message
type ErrorDetail struct {
	Status int    `json:"status,omitempty"`
	Title  string `json:"title,omitempty"`
	Detail string `json:"detail"`
}

// NewClient returns a new Oura API client. If a nil httpClient is
// provided, http.DefaultClient will be used. To use API methods which require
// authentication, provide an http.Client that will perform the authentication
// for you (such as that provided by the golang.org/x/oauth2 library).
func NewClient(cc *http.Client) *Client {
	if cc == nil {
		cc = http.DefaultClient
	}
	baseURL, _ := url.Parse(BaseURL)

	c := &Client{baseURL: baseURL, UserAgent: userAgent, client: cc}
	return c
}

// NewRequest creates an HTTP Request. The client baseURL is checked to confirm that it has a trailing
// slash. A relative URL should be provided without the leading slash. If a non-nil body is provided
// it will be JSON encoded and included in the request.
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
	if c.UserAgent != "" {
		req.Header.Set("User-Agent", c.UserAgent)
	}
	return req, nil
}

// Do sends a request and returns the response. An error is returned if the request cannot
// be sent or if the API returns an error. If a response is received, the body response body
// is decoded and stored in the value pointed to by v.
func (c *Client) Do(ctx context.Context, req *http.Request, v interface{}) (*http.Response, error) {
	req = req.WithContext(ctx)
	resp, err := c.client.Do(req)
	if err != nil {
		return nil, err
	}

	data, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	resp.Body.Close()

	// Anything other than a HTTP 2xx response code is treated as an error.
	if resp.StatusCode >= 300 {
		e := ErrorDetail{}
		err := json.Unmarshal(data, &e)
		if err != nil {
			return resp, err
		}

		err = errors.New(http.StatusText(resp.StatusCode) + ": " + e.Detail)
		return resp, err
	}

	if v != nil && len(data) != 0 {
		err = json.Unmarshal(data, v)

		switch err {
		case nil:
		case io.EOF:
			err = nil
		default:
		}
	}

	return resp, err
}

// parametiseDate takes the arguments and URL encodes them into a string
// where the dates are ISO 8601 date strings without times.
func parametiseDate(path, start, end, next string) string {
	params := url.Values{}

	if start != "" {
		params.Add("start_date", start)
	}
	if end != "" {
		params.Add("end_date", end)
	}
	if next != "" {
		params.Add("next_token", next)
	}
	if len(params) > 0 {
		path += fmt.Sprintf("?%s", params.Encode())
	}
	return path
}

// parametiseDate takes the arguments and URL encodes them into a string
// where the dates are ISO 8601 date strings with times.
func parametiseDatetime(path, start, end, next string) string {
	params := url.Values{}

	if start != "" {
		params.Add("start_datetime", start)
	}
	if end != "" {
		params.Add("end_datetime", end)
	}
	if next != "" {
		params.Add("next_token", next)
	}
	if len(params) > 0 {
		path += fmt.Sprintf("?%s", params.Encode())
	}
	return path
}
