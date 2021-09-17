package restclient

/*
MockRESTClient is a mock for RESTClient
*/
type MockRESTClient struct {
	DELETEFunc            func(path string, successReceiver, errorReceiver interface{}) (bool, error)
	GETFunc               func(path string, successReceiver, errorReceiver interface{}) (bool, error)
	POSTFunc              func(path string, body, successReceiver, errorReceiver interface{}) (bool, error)
	PUTFunc               func(path string, body, successReceiver, errorReceiver interface{}) (bool, error)
	WithAuthorizationFunc func(auth string) RESTClient
}

/*
DELETE is a mock method
*/
func (m MockRESTClient) DELETE(path string, successReceiver, errorReceiver interface{}) (bool, error) {
	return m.DELETEFunc(path, successReceiver, errorReceiver)
}

/*
GET is a mock method
*/
func (m MockRESTClient) GET(path string, successReceiver, errorReceiver interface{}) (bool, error) {
	return m.GETFunc(path, successReceiver, errorReceiver)
}

/*
POST is a mock method
*/
func (m MockRESTClient) POST(path string, body, successReceiver, errorReceiver interface{}) (bool, error) {
	return m.POSTFunc(path, body, successReceiver, errorReceiver)
}

/*
PUT is a mock method
*/
func (m MockRESTClient) PUT(path string, body, successReceiver, errorReceiver interface{}) (bool, error) {
	return m.PUTFunc(path, body, successReceiver, errorReceiver)
}

/*
WithAuthorization is a mock method
*/
func (m MockRESTClient) WithAuthorization(auth string) RESTClient {
	return m.WithAuthorizationFunc(auth)
}
