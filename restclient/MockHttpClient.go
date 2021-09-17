/*
 * Copyright (c) 2021. App Nerds LLC. All rights reserved
 */

package restclient

import (
	"net/http"
)

/*
MockHTTPClient is a mock HTTP client
*/
type MockHTTPClient struct {
	DoFunc           func(req *http.Request) (*http.Response, error)
	SetTransportFunc func(transport *http.Transport)
}

/*
Do mocks the HTTP Do method
*/
func (m *MockHTTPClient) Do(req *http.Request) (*http.Response, error) {
	return m.DoFunc(req)
}

/*
SetTransport mocks the SetTransport method
*/
func (m *MockHTTPClient) SetTransport(transport *http.Transport) {
	m.SetTransportFunc(transport)
}
