package api

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"path"

	"github.com/paveg/goura/oura"
	"github.com/pkg/errors"
)

const (
	// V2BasePath is the base path for v2 API endpoints.
	V2BasePath = "/v2/usercollection"
)

// Client represents a HTTP client for Oura API.
type Client struct {
	EndpointURL *url.URL
	HTTPClient  *http.Client
	UserAgent   string
	AccessToken string
}

// NewClient creates a new http client.
func NewClient(endpointURL string, httpClient *http.Client, userAgent, token string) (*Client, error) {
	parsedURL, err := url.ParseRequestURI(endpointURL)
	if err != nil {
		return nil, errors.Wrapf(err, "failed to parse url: %s", endpointURL)
	}

	client := &Client{
		EndpointURL: parsedURL,
		HTTPClient:  httpClient,
		UserAgent:   userAgent,
		AccessToken: token,
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

	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("User-Agent", client.UserAgent)
	if client.AccessToken != "" {
		req.Header.Set("Authorization", fmt.Sprintf("Bearer %s", client.AccessToken))
	}

	return req, nil
}

func decodeBody(resp *http.Response, iFace interface{}) error {
	defer resp.Body.Close()
	decoder := json.NewDecoder(resp.Body)

	return decoder.Decode(iFace)
}

// checkResponse checks the HTTP response for errors.
func checkResponse(resp *http.Response) error {
	if resp.StatusCode >= 200 && resp.StatusCode < 300 {
		return nil
	}

	body, _ := io.ReadAll(resp.Body)
	return fmt.Errorf("API error: status=%d, body=%s", resp.StatusCode, string(body))
}

// getWithDatePeriod makes a GET request with date range query parameters.
func (client *Client) getWithDatePeriod(ctx context.Context, subURL string, datePeriod oura.DatePeriod) (*http.Response, error) {
	httpRequest, err := client.newRequest(ctx, "GET", subURL, nil)
	if err != nil {
		return nil, err
	}

	q := httpRequest.URL.Query()
	if datePeriod.StartDate != "" {
		q.Add("start_date", datePeriod.StartDate)
	}
	if datePeriod.EndDate != "" {
		q.Add("end_date", datePeriod.EndDate)
	}
	httpRequest.URL.RawQuery = q.Encode()

	httpResponse, err := client.HTTPClient.Do(httpRequest)
	if err != nil {
		return nil, err
	}

	return httpResponse, nil
}

// getWithDateTimePeriod makes a GET request with datetime range query parameters.
func (client *Client) getWithDateTimePeriod(ctx context.Context, subURL string, period oura.DateTimePeriod) (*http.Response, error) {
	httpRequest, err := client.newRequest(ctx, "GET", subURL, nil)
	if err != nil {
		return nil, err
	}

	q := httpRequest.URL.Query()
	if period.StartDateTime != "" {
		q.Add("start_datetime", period.StartDateTime)
	}
	if period.EndDateTime != "" {
		q.Add("end_datetime", period.EndDateTime)
	}
	httpRequest.URL.RawQuery = q.Encode()

	httpResponse, err := client.HTTPClient.Do(httpRequest)
	if err != nil {
		return nil, err
	}

	return httpResponse, nil
}

// getWithNextToken makes a GET request with pagination token.
func (client *Client) getWithNextToken(ctx context.Context, subURL string, datePeriod oura.DatePeriod, nextToken string) (*http.Response, error) {
	httpRequest, err := client.newRequest(ctx, "GET", subURL, nil)
	if err != nil {
		return nil, err
	}

	q := httpRequest.URL.Query()
	if datePeriod.StartDate != "" {
		q.Add("start_date", datePeriod.StartDate)
	}
	if datePeriod.EndDate != "" {
		q.Add("end_date", datePeriod.EndDate)
	}
	if nextToken != "" {
		q.Add("next_token", nextToken)
	}
	httpRequest.URL.RawQuery = q.Encode()

	httpResponse, err := client.HTTPClient.Do(httpRequest)
	if err != nil {
		return nil, err
	}

	return httpResponse, nil
}

// getByDocumentID makes a GET request for a single document by ID.
func (client *Client) getByDocumentID(ctx context.Context, subURL string, documentID string) (*http.Response, error) {
	fullURL := fmt.Sprintf("%s/%s", subURL, documentID)
	httpRequest, err := client.newRequest(ctx, "GET", fullURL, nil)
	if err != nil {
		return nil, err
	}

	httpResponse, err := client.HTTPClient.Do(httpRequest)
	if err != nil {
		return nil, err
	}

	return httpResponse, nil
}
