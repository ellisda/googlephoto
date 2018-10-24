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

1. There is an issue in the http listener code where the shutdown doesn't work right and causes a ~30 sec delay when performing auth. Refreshing the redirected browser window kicks it.

1. We're not detecting expired refresh tokens on disk or refreshing the token against the server. Delete the file on disk as a workaround.

1. Add some Client API structs and methods for Album/MediaItem routes, maybe with code generation from a swagger spec? (see [API reference](https://developers.google.com/photos/library/reference/))

1. Do something useful (like ask google to provide the "storage quality" in the returned media items)

## References

- Google Photos [API Reference](https://developers.google.com/photos/library/reference/)
- Discussion about which IDs to persist locally - Reference: [Access media items](https://developers.google.com/photos/library/guides/access-media-items)
- Mention that [items uploaded through this API will be stored at "original quality"](https://developers.google.com/photos/library/guides/api-limits-quotas#photo-storage-quality)
- Older picasaweb client API Sample from [Bowbaq](https://github.com/Bowbaq/googlephoto)
- Golang oauth2 client [example](https://github.com/golang/oauth2/blob/master/google/example_test.go)


