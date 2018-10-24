package main

import (
	"context"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/pkg/browser"
	"golang.org/x/oauth2"
	"golang.org/x/oauth2/google"
)

//GetClient returns an http.Client with OAuth2 refresh token builtin
func GetClient(port int) (*http.Client, error) {
	conf := &oauth2.Config{
		// Get Oauth2 credentials from https://console.developers.google.com
		ClientID:     "<Insert Your OAuth2 Client ID>",
		ClientSecret: "<Insert Your OAuth2 Client Secret>",
		RedirectURL:  "<Insert Your OAuth2 RedirectURL>",
		Scopes: []string{
			"https://www.googleapis.com/auth/photoslibrary.readonly",
		},
		Endpoint: google.Endpoint,
	}

	var token oauth2.Token
	tokenData, err := ioutil.ReadFile("./token.json")
	if err != nil {
		fmt.Println("Refresh Token not present - initiating User Auth")
		token, err = GetToken(port, conf)
		fmt.Printf("Refresh Token Retrieved: %v\n", token.RefreshToken)
		if err != nil {
			return nil, err
		}
	} else {
		fmt.Printf("Found Refresh Token locally: %v\n", tokenData)
		err = json.Unmarshal(tokenData, &token)
		if err != nil {
			return nil, err
		}
	}

	return conf.Client(oauth2.NoContext, &token), nil
}

//GetToken will launch an HTTP listener on designated port to catch the
// redirect from the google authentication
// (Note: redirect URL must match project configuration on google server)
func GetToken(port int, conf *oauth2.Config) (oauth2.Token, error) {
	var authCode string
	var err error
	if port != 0 {
		authCode, err = getAuthCodeFromHTTPCallback(conf, port)
		check(err)
	} else {
		authCode, err = getAuthCodeFromCLI(conf, port)
		check(err)
	}

	token, err := conf.Exchange(oauth2.NoContext, authCode)
	check(err)

	data, err := json.Marshal(token)
	check(err)

	ioutil.WriteFile("./token.json", data, 0400)
	return *token, nil
}

func getAuthCodeFromHTTPCallback(conf *oauth2.Config, port int) (string, error) {
	url := conf.AuthCodeURL("state", oauth2.AccessTypeOffline)
	browser.OpenURL(url)
	fmt.Printf("You must allow this application access to your google photos, please see %v\n", url)

	ch := make(chan string, 100)
	mux := http.NewServeMux()

	mux.HandleFunc("/auth/google/callback", func(w http.ResponseWriter, r *http.Request) {
		if code, ok := r.URL.Query()["code"]; ok && len(code) >= 1 {
			fmt.Printf("Recv Code:%v\n", code)
			ch <- code[0]
			w.Write([]byte(`<!doctype html><html><head />
			<body>Success - Please close this tab now</body>
			</html>`))
		} else {
			w.WriteHeader(400)
		}
	})
	s := http.Server{
		Addr:    fmt.Sprintf(":%d", port),
		Handler: mux,
	}
	go func() {
		if err := s.ListenAndServe(); err != http.ErrServerClosed {
			log.Fatalf("Callback ListenAndServe returned err:%v\n", err)
		}
	}()

	//Block till callback handler has retrieved the refresh token
	authCode := <-ch

	if err := s.Shutdown(context.Background()); err != nil && err != http.ErrServerClosed {
		log.Fatalf("Callback Shutdown returned err:%v\n", err)
	}
	return authCode, nil
}

func getAuthCodeFromCLI(conf *oauth2.Config, port int) (string, error) {
	url := conf.AuthCodeURL("state", oauth2.AccessTypeOffline)
	fmt.Printf("Visit the URL for the auth dialog: %v", url)

	fmt.Println("Your browser should have redirected to: http://localhost/?state=state&code=<code>")
	fmt.Print("Paste the code: ")

	authCode := ""
	fmt.Scanln(&authCode)

	return authCode, nil
}
