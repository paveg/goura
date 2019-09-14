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

// Client represents a HTTP client
type Client struct {
	EndpointURL *url.URL
	HTTPClient  *http.Client
	UserAgent   string
	AccessToken string
}

// NewClient creates a new http client
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

	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
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

func (client *Client) getRequestWithDatePeriod(ctx context.Context, subURL string, datePeriod oura.DatePeriod) (*http.Response, error) {
	httpRequest, err := client.newRequest(ctx, "GET", subURL, nil)
	if err != nil {
		return nil, err
	}

	httpRequest.Header.Set("Content-Type", "application/json")
	q := httpRequest.URL.Query()
	q.Add("start", datePeriod.StartDate)
	q.Add("end", datePeriod.EndDate)
	httpRequest.URL.RawQuery = q.Encode()
	httpResponse, err := client.HTTPClient.Do(httpRequest)
	if err != nil {
		return nil, err
	}

	return httpResponse, nil
}
