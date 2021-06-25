/*
 * Copyright (c) 2021. App Nerds LLC. All rights reserved
 */

package restclient2

/*
RESTClient defines an interface for working with RESTful endpoints
*/
type RESTClient interface {
	DELETE(path string, successReceiver, errorReceiver interface{}) (bool, error)
	GET(path string, successReceiver, errorReceiver interface{}) (bool, error)
	POST(path string, body, successReceiver, errorReceiver interface{}) (bool, error)
	PUT(path string, body, successReceiver, errorReceiver interface{}) (bool, error)
	WithAuthorization(auth string) RESTClient
}
