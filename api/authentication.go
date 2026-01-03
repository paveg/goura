package api

import (
	"context"
	"crypto/rand"
	"encoding/json"
	"fmt"
	"net"
	"net/http"
	"time"

	"github.com/skratchdot/open-golang/open"
	"golang.org/x/oauth2"
)

const (
	// OuraAuthURL is the OAuth2 authorization endpoint.
	OuraAuthURL = "https://cloud.ouraring.com/oauth/authorize" //nolint:gosec // URL, not a credential
	// OuraTokenURL is the OAuth2 token endpoint.
	OuraTokenURL = "https://api.ouraring.com/oauth/token" //nolint:gosec // URL, not a credential
	// DefaultRedirectURL is the default local callback URL.
	DefaultRedirectURL = "http://localhost:8989"
)

// TokenResponse represents the OAuth2 token response.
type TokenResponse struct {
	AccessToken  string `json:"access_token"`
	TokenType    string `json:"token_type"`
	ExpiresIn    int    `json:"expires_in"`
	RefreshToken string `json:"refresh_token,omitempty"`
	Scope        string `json:"scope,omitempty"`
}

// OAuthConfig holds the OAuth2 configuration.
type OAuthConfig struct {
	ClientID     string
	ClientSecret string
	RedirectURL  string
	Scopes       []string
}

// DefaultScopes returns the default OAuth2 scopes for Oura API v2.
func DefaultScopes() []string {
	return []string{
		"email",
		"personal",
		"daily",
		"heartrate",
		"workout",
		"session",
		"tag",
		"spo2",
	}
}

// NewOAuthConfig creates a new OAuth configuration with default settings.
func NewOAuthConfig(clientID, clientSecret string) *OAuthConfig {
	return &OAuthConfig{
		ClientID:     clientID,
		ClientSecret: clientSecret,
		RedirectURL:  DefaultRedirectURL,
		Scopes:       DefaultScopes(),
	}
}

// oauth2Config returns the golang.org/x/oauth2 configuration.
func (c *OAuthConfig) oauth2Config() *oauth2.Config {
	return &oauth2.Config{
		ClientID:     c.ClientID,
		ClientSecret: c.ClientSecret,
		RedirectURL:  c.RedirectURL,
		Scopes:       c.Scopes,
		Endpoint: oauth2.Endpoint{
			AuthURL:  OuraAuthURL,
			TokenURL: OuraTokenURL,
		},
	}
}

// FetchAccessToken fetches an access token using the authorization code flow.
// This is the recommended flow for OAuth2 and provides refresh token support.
func FetchAccessToken(clientID string, clientSecret string) (string, error) {
	config := NewOAuthConfig(clientID, clientSecret)
	return config.FetchToken()
}

// FetchToken performs the OAuth2 authorization code flow and returns the access token.
func (c *OAuthConfig) FetchToken() (string, error) {
	token, err := c.FetchFullToken()
	if err != nil {
		return "", err
	}
	return token.AccessToken, nil
}

// FetchFullToken performs the OAuth2 authorization code flow and returns the full token response.
func (c *OAuthConfig) FetchFullToken() (*TokenResponse, error) {
	l, err := net.Listen("tcp", "localhost:8989")
	if err != nil {
		return nil, fmt.Errorf("failed to start local server: %w", err)
	}
	defer closer(l)

	oauth2Conf := c.oauth2Config()

	stateBytes := make([]byte, 16)
	_, err = rand.Read(stateBytes)
	if err != nil {
		return nil, fmt.Errorf("failed to generate state: %w", err)
	}

	state := fmt.Sprintf("%x", stateBytes)

	// Use authorization code flow (response_type=code) instead of implicit flow
	authURL := oauth2Conf.AuthCodeURL(state)
	err = open.Start(authURL)
	if err != nil {
		return nil, fmt.Errorf("failed to open browser: %w", err)
	}

	codeChan := make(chan string)
	errChan := make(chan error)

	go serveAuthCallback(l, state, codeChan, errChan)

	select {
	case code := <-codeChan:
		// Exchange authorization code for tokens
		token, err := oauth2Conf.Exchange(context.Background(), code)
		if err != nil {
			return nil, fmt.Errorf("failed to exchange code for token: %w", err)
		}

		return &TokenResponse{
			AccessToken:  token.AccessToken,
			TokenType:    token.TokenType,
			RefreshToken: token.RefreshToken,
		}, nil

	case err := <-errChan:
		return nil, err

	case <-time.After(5 * time.Minute):
		return nil, fmt.Errorf("timeout waiting for authorization callback")
	}
}

// RefreshAccessToken refreshes an access token using a refresh token.
func RefreshAccessToken(clientID, clientSecret, refreshToken string) (*TokenResponse, error) {
	config := NewOAuthConfig(clientID, clientSecret)
	return config.RefreshToken(refreshToken)
}

// RefreshToken exchanges a refresh token for a new access token.
func (c *OAuthConfig) RefreshToken(refreshToken string) (*TokenResponse, error) {
	oauth2Conf := c.oauth2Config()

	token := &oauth2.Token{
		RefreshToken: refreshToken,
	}

	tokenSource := oauth2Conf.TokenSource(context.Background(), token)
	newToken, err := tokenSource.Token()
	if err != nil {
		return nil, fmt.Errorf("failed to refresh token: %w", err)
	}

	return &TokenResponse{
		AccessToken:  newToken.AccessToken,
		TokenType:    newToken.TokenType,
		RefreshToken: newToken.RefreshToken,
	}, nil
}

func closer(l net.Listener) {
	if err := l.Close(); err != nil {
		fmt.Printf("ERROR: %#v\n", err)
	}
}

func serveAuthCallback(l net.Listener, expectedState string, codeChan chan string, errChan chan error) {
	server := &http.Server{
		ReadHeaderTimeout: 10 * time.Second,
		Handler: http.HandlerFunc(func(w http.ResponseWriter, req *http.Request) {
			if req.URL.Path == "/" {
				// Check for authorization code
				code := req.URL.Query().Get("code")
				state := req.URL.Query().Get("state")
				errorParam := req.URL.Query().Get("error")

				if errorParam != "" {
					errorDesc := req.URL.Query().Get("error_description")
					w.WriteHeader(http.StatusBadRequest)
					_, _ = w.Write([]byte(fmt.Sprintf("<html><body><h1>Authorization Failed</h1><p>%s: %s</p></body></html>", errorParam, errorDesc)))
					w.(http.Flusher).Flush()
					errChan <- fmt.Errorf("authorization error: %s - %s", errorParam, errorDesc)
					return
				}

				if code != "" && state == expectedState {
					_, _ = w.Write([]byte(`<html><body><h1>Authorization Successful!</h1><p>You can close this window now.</p><script>window.close()</script></body></html>`))
					w.(http.Flusher).Flush()
					codeChan <- code
					return
				}

				if state != expectedState && state != "" {
					w.WriteHeader(http.StatusBadRequest)
					_, _ = w.Write([]byte(`<html><body><h1>State Mismatch</h1><p>Invalid state parameter. Please try again.</p></body></html>`))
					w.(http.Flusher).Flush()
					errChan <- fmt.Errorf("state mismatch: expected %s, got %s", expectedState, state)
					return
				}

				// Fallback for implicit flow (legacy support)
				_, _ = w.Write([]byte(`<script>location.href = "/close?" + location.hash.substring(1);</script>`))
			} else if req.URL.Path == "/close" {
				// Legacy implicit flow callback
				accessToken := req.URL.Query().Get("access_token")
				if accessToken != "" {
					_, _ = w.Write([]byte(`<html><body><h1>Authorization Successful!</h1><p>You can close this window now.</p><script>window.close()</script></body></html>`))
					w.(http.Flusher).Flush()
					codeChan <- accessToken
				}
			}
		}),
	}

	_ = server.Serve(l)
}

// ValidateToken checks if the token is valid by making a request to the personal info endpoint.
func ValidateToken(accessToken string) error {
	client := &http.Client{Timeout: 10 * time.Second}

	req, err := http.NewRequest("GET", "https://api.ouraring.com/v2/usercollection/personal_info", nil)
	if err != nil {
		return err
	}

	req.Header.Set("Authorization", fmt.Sprintf("Bearer %s", accessToken))

	resp, err := client.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		var errorResp struct {
			Detail string `json:"detail"`
		}
		_ = json.NewDecoder(resp.Body).Decode(&errorResp)
		return fmt.Errorf("token validation failed: status=%d, detail=%s", resp.StatusCode, errorResp.Detail)
	}

	return nil
}
