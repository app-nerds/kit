/*
 * Copyright (c) 2021. App Nerds LLC. All rights reserved
 */

package restclient

import "net/http"

/*
RESTClient defines an interface for working with RESTful endpoints
*/
type RESTClient interface {
	DELETE(path string, successReceiver, errorReceiver interface{}) (*http.Response, error)
	GET(path string, successReceiver, errorReceiver interface{}) (*http.Response, error)
	POST(path string, body, successReceiver, errorReceiver interface{}) (*http.Response, error)
	PUT(path string, body, successReceiver, errorReceiver interface{}) (*http.Response, error)
	WithAuthorization(auth string) RESTClient
}
