/*
 * Copyright (c) 2020. App Nerds LLC. All rights reserved
 */

package restclient

import "net/http"

/*
IHTTPClient is an interface over the Go HttpClient
*/
type IHTTPClient interface {
	Do(req *http.Request) (*http.Response, error)
	SetTransport(transport *http.Transport)
}

type HTTPClient struct {
	*http.Client
}

func (h *HTTPClient) Do(req *http.Request) (*http.Response, error) {
	return h.Client.Do(req)
}

func (h *HTTPClient) SetTransport(transport *http.Transport) {
	h.Client.Transport = transport
}
