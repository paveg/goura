package api

import (
	"net/http"
	"net/http/httptest"
	"net/url"
)

// NewMockServer initializes mockServer
func NewMockServer() (*http.ServeMux, *url.URL) {
	mux := http.NewServeMux()
	server := httptest.NewServer(mux)
	mockServerURL, _ := url.Parse(server.URL)

	return mux, mockServerURL
}

// NewTestClient initializes mockClient
func NewTestClient(mockServerURL *url.URL) *Client {
	endpointURL := mockServerURL.String() + "/api"
	httpClient := &http.Client{}
	userAgent := "test client"
	client, _ := NewClient(endpointURL, httpClient, userAgent)

	return client
}
