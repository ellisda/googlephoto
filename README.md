# Google Photo API client (Go WebServer)


Here is some sample code that connects via OAuth2 to the Google Photos API, written in go.

## Steps to Get Working

Steps to compile and run:
1.  Go to https://console.developers.google.com/apis/library and Create a New Project
1. For your new project, find the "Photos Library API" in the API Library and enable it from google's side
1. Configure new Credentials: OAuth Client ID
    - Web Application, with a defined redirect URL (here I use http://127.0.0.1:8080/auth/google/callback)
1. Paste your Client ID, Secret, and Redirect URL into the source code
1. Compile and Run (note: First time, user will have to authenticate as user to google and allow access to this application)
    - NOTE: We're only requesting the read scope 


Build

    go build

Run

    go run main.go



## References

- [Google Photos API Reference](https://developers.google.com/photos/library/reference/)
- [Older picasaweb client API Sample from Bowbaq](https://github.com/Bowbaq/googlephoto)
- [Golang oauth2 client example](https://github.com/golang/oauth2/blob/master/google/example_test.go)


