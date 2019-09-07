package api

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"net/url"
	"path"

	"github.com/pkg/errors"
)

// Client represents a HTTP client
type Client struct {
	EndpointURL *url.URL
	HTTPClient  *http.Client
	UserAgent   string
}

// NewClient creates a new http client
func NewClient(endpointURL string, httpClient *http.Client, userAgent string) (*Client, error) {
	parsedURL, err := url.ParseRequestURI(endpointURL)
	if err != nil {
		return nil, errors.Wrapf(err, "failed to parse url: %s", endpointURL)
	}

	client := &Client{
		EndpointURL: parsedURL,
		HTTPClient:  httpClient,
		UserAgent:   userAgent,
	}

	return client, nil
}

func (client *Client) newRequest(ctx context.Context, method string, subURL string, body io.Reader) (*http.Request, error) {
	endpointURL := *client.EndpointURL
	endpointURL.Path = path.Join(client.EndpointURL.Path, subURL)

	req, err := http.NewRequest(method, endpointURL.String(), body)
	if err != nil {
		return nil, err
	}

	req = req.WithContext(ctx)

	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	req.Header.Set("User-Agent", client.UserAgent)

	return req, nil
}

func decodeBody(resp *http.Response, iFace interface{}) error {
	defer resp.Body.Close()
	decoder := json.NewDecoder(resp.Body)

	return decoder.Decode(iFace)
}

// TODO: extract package or files.

// UserInfo requests GET /v1/userinfo
func (client *Client) UserInfo(ctx context.Context, token string) (*UserInfo, error) {
	subURL := fmt.Sprint("/v1/userinfo")
	httpRequest, err := client.newRequest(ctx, "GET", subURL, nil)
	if err != nil {
		return nil, err
	}
	httpRequest.Header.Set("Content-Type", "application/json")
	httpRequest.Header.Set("Authorization", fmt.Sprintf("Bearer %s", token))
	httpResponse, err := client.HTTPClient.Do(httpRequest)
	if err != nil {
		return nil, err
	}

	var apiResponse UserInfo
	log.Printf("HTTP Request: %v", httpResponse.Status)
	if err := decodeBody(httpResponse, &apiResponse); err != nil {
		return nil, err
	}

	return &apiResponse, nil
}

// Sleep requests GET /v1/sleep
func (client *Client) Sleep(ctx context.Context, token string, datePeriod DatePeriod) (*SleepPeriods, error) {
	subURL := fmt.Sprintf("/v1/sleep")
	httpRequest, err := client.newRequest(ctx, "GET", subURL, nil)
	if err != nil {
		return nil, err
	}

	httpRequest.Header.Set("Content-Type", "application/json")
	httpRequest.Header.Set("Authorization", fmt.Sprintf("Bearer %s", token))
	q := httpRequest.URL.Query()
	q.Add("start", datePeriod.StartDate)
	q.Add("end", datePeriod.EndDate)
	httpRequest.URL.RawQuery = q.Encode()
	httpResponse, err := client.HTTPClient.Do(httpRequest)
	if err != nil {
		return nil, err
	}

	var apiResponse SleepPeriods
	log.Printf("HTTP Request: %v", httpResponse.Status)
	if err := decodeBody(httpResponse, &apiResponse); err != nil {
		return nil, err
	}

	return &apiResponse, nil
}

// Activity requests GET /v1/activity
func (client *Client) Activity(ctx context.Context, token string, datePeriod DatePeriod) (*Activities, error) {
	subURL := fmt.Sprintf("/v1/activity")
	httpRequest, err := client.newRequest(ctx, "GET", subURL, nil)
	if err != nil {
		return nil, err
	}

	httpRequest.Header.Set("Content-Type", "application/json")
	httpRequest.Header.Set("Authorization", fmt.Sprintf("Bearer %s", token))
	q := httpRequest.URL.Query()
	q.Add("start", datePeriod.StartDate)
	q.Add("end", datePeriod.EndDate)
	httpRequest.URL.RawQuery = q.Encode()
	httpResponse, err := client.HTTPClient.Do(httpRequest)
	if err != nil {
		return nil, err
	}

	var apiResponse Activities
	log.Printf("HTTP Request: %v", httpResponse.Status)
	if err := decodeBody(httpResponse, &apiResponse); err != nil {
		return nil, err
	}

	return &apiResponse, nil
}
