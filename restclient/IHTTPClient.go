package restclient

import "net/http"

/*
IHTTPClient is an interface over the Go HttpClient
*/
type IHTTPClient interface {
	Do(req *http.Request) (*http.Response, error)
}
