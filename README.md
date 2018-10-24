# Google Photo API client (Go WebServer)


Here is some sample code that consumes the Google Photos Rest API, using OAuth2 authentication.

Build

    go build

Run with a http listener to catch the auth code

    ./googlephoto

Run solely on command line (user must copy / paste auth code into terminal)

    ./googlephoto --port 0

## Steps to Get Working

Steps to configure OAuth2 client:
1. Create a new Project at https://console.developers.google.com/apis/library
1. Select your new project, then find and enable the "Photos Library API" at https://console.developers.google.com/apis/library
1. Configure new Credentials at https://console.developers.google.com/apis/credentials/
    - OAuth Client ID
    - Web Application
    - Provide the following "Authorized redirect url" - http://127.0.0.1:8080/auth/google/callback
1. Paste the generated client ID and secret into the `client.go` file
1. Run the program, authenticate to google, and grant access to the app
1. Review your app access grants at https://myaccount.google.com/permissions


## Things Left to do

There is an issue in the http listener code where the shutdown doesn't work right and causes a ~30 sec delay when performing auth. Refreshing the redirected browser window kicks it.

We're not detecting expired refresh tokens on disk or refreshing the token against the server. Delete the file on disk as a workaround.

## References

- [Google Photos API Reference](https://developers.google.com/photos/library/reference/)
- [Older picasaweb client API Sample from Bowbaq](https://github.com/Bowbaq/googlephoto)
- [Golang oauth2 client example](https://github.com/golang/oauth2/blob/master/google/example_test.go)


