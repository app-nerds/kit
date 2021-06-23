/*
 * Copyright (c) 2021. App Nerds LLC. All rights reserved
 */

package restclient2

import "net/http"

/*
HTTPClientInterface is an interface over the Go HttpClient
*/
type HTTPClientInterface interface {
	Do(req *http.Request) (*http.Response, error)
	SetTransport(transport *http.Transport)
}

/*
HTTPClient implements methods as a wrapper over the Go HttpClient
*/
type HTTPClient struct {
	*http.Client
}

func (h *HTTPClient) Do(req *http.Request) (*http.Response, error) {
	return h.Client.Do(req)
}

func (h *HTTPClient) SetTransport(transport *http.Transport) {
	h.Client.Transport = transport
}
