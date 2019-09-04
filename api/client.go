package api

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
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

func newClient(endpointURL string, httpClient *http.Client, userAgent string) (*Client, error) {
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

// UserInfo requests GET /v1/userinfo
func (client *Client) UserInfo(ctx context.Context) (*UserInfoResponse, error) {
	subURL := fmt.Sprintln("/v1/userinfo")
	httpRequest, err := client.newRequest(ctx, "GET", subURL, nil)
	if err != nil {
		return nil, err
	}

	httpResponse, err := client.HTTPClient.Do(httpRequest)
	if err != nil {
		return nil, err
	}

	var apiResponse UserInfoResponse
	if err := decodeBody(httpResponse, &apiResponse); err != nil {
		return nil, err
	}

	return &apiResponse, nil
}
