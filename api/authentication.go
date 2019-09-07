package api

import (
	"crypto/rand"
	"fmt"
	"net"
	"net/http"

	"github.com/skratchdot/open-golang/open"
	"golang.org/x/oauth2"
)

// FetchAccessToken fetches access token
func FetchAccessToken(clientID string, clientSecret string) (string, error) {
	l, err := net.Listen("tcp", "localhost:8989")
	if err != nil {
		return "", err
	}

	defer closer(l)

	oauthConfig := &oauth2.Config{
		Scopes: []string{
			"email",
			"personal",
			"daily",
		},
		Endpoint: oauth2.Endpoint{
			AuthURL:  "https://cloud.ouraring.com/oauth/authorize",
			TokenURL: "https://api.ouraring.com/oauth/token",
		},
		ClientID:     clientID,
		ClientSecret: clientSecret,
		RedirectURL:  "http://localhost:8989",
	}

	stateBytes := make([]byte, 16)
	_, err = rand.Read(stateBytes)
	if err != nil {
		return "", err
	}

	state := fmt.Sprintf("%x", stateBytes)
	err = open.Start(oauthConfig.AuthCodeURL(state, oauth2.SetAuthURLParam("response_type", "token")))
	if err != nil {
		return "", nil
	}

	quit := make(chan string)
	go serve(l, err, quit)

	return <-quit, nil
}

func closer(l net.Listener) {
	if err := l.Close(); err != nil {
		fmt.Printf("ERROR: %#v\n", err)
	}
}

func serve(l net.Listener, err error, q chan string) {
	http.Serve(l, http.HandlerFunc(func(w http.ResponseWriter, req *http.Request) {
		if req.URL.Path == "/" {
			_, err = w.Write([]byte(`<script>location.href = "/close?" + location.hash.substring(1);</script>`))
			if err != nil {
				fmt.Printf("ERROR: %#v", err)
				return
			}
		} else {
			_, err = w.Write([]byte(`<script>window.open("about:blank","_self").close()</script>`))
			if err != nil {
				fmt.Printf("ERROR: %#v", err)
				return
			}
			w.(http.Flusher).Flush()
			q <- req.URL.Query().Get("access_token")
		}
	}))
}
