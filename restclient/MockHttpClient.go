/*
 * Copyright (c) 2020. App Nerds LLC. All rights reserved
 */

package restclient

import (
	"net/http"
)

type MockHttpClient struct {
	DoFunc           func(req *http.Request) (*http.Response, error)
	SetTransportFunc func(transport *http.Transport)
}

func (m *MockHttpClient) Do(req *http.Request) (*http.Response, error) {
	return m.DoFunc(req)
}

func (m *MockHttpClient) SetTransport(transport *http.Transport) {
	m.SetTransportFunc(transport)
}
