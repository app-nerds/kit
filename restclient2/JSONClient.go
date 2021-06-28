/*
 * Copyright (c) 2021. App Nerds LLC. All rights reserved
 */

package restclient2

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strings"

	"github.com/app-nerds/kit/v4/restclient2/responsegetter"
)

/*
JSONClient provides a set of methods for working with RESTful endpoints that accept
and return JSON data
*/
type JSONClient struct {
	BaseURL    string
	HTTPClient HTTPClientInterface

	authorization string
}

/*
NewJSONClient creates a new JSON-based REST client
*/
func NewJSONClient(baseURL string, httpClient HTTPClientInterface) JSONClient {
	return JSONClient{
		BaseURL:    baseURL,
		HTTPClient: httpClient,

		authorization: "",
	}
}

/*
DELETE performs an HTTP DELETE operation. You provide a path, which should exclude the
the TLD, as this is defined in BaseURL. If your request requires authorization
call WithAuthorization first, then DELETE.

Upon a 2xx response the successReceiver will be populated with the returned JSON
data, and "true,nil" will be returned. If there is an error response from the server
errorReceiver will be populated, and "false,nil" will be returned.

In the event that an error occured communicating with the server, or some other
unforseen error occurs, "false,error" is returned.
*/
func (c JSONClient) DELETE(path string, successReceiver, errorReceiver interface{}) (bool, error) {
	var (
		err      error
		request  *http.Request
		response *http.Response
	)

	if request, err = c.createRequest(path, "DELETE", c.authorization, nil); err != nil {
		return false, err
	}

	if response, err = c.HTTPClient.Do(request); err != nil {
		return false, fmt.Errorf("error executing request: %w", err)
	}

	defer response.Body.Close()
	return responsegetter.Get(response, successReceiver, errorReceiver)
}

/*
GET performs an HTTP GET operation. You provide a path, which should exclude the
the TLD, as this is defined in BaseURL. If your request requires authorization
call WithAuthorization first, then GET.

Upon a 2xx response the successReceiver will be populated with the returned JSON
data, and "true,nil" will be returned. If there is an error response from the server
errorReceiver will be populated, and "false,nil" will be returned.

In the event that an error occured communicating with the server, or some other
unforseen error occurs, "false,error" is returned.
*/
func (c JSONClient) GET(path string, successReceiver, errorReceiver interface{}) (bool, error) {
	var (
		err      error
		request  *http.Request
		response *http.Response
	)

	if request, err = c.createRequest(path, "GET", c.authorization, nil); err != nil {
		return false, err
	}

	if response, err = c.HTTPClient.Do(request); err != nil {
		return false, fmt.Errorf("error executing request: %w", err)
	}

	defer response.Body.Close()
	return responsegetter.Get(response, successReceiver, errorReceiver)
}

/*
POST performs an HTTP POST operation. You provide a path, which should exclude the
the TLD, as this is defined in BaseURL. If your request requires authorization
call WithAuthorization first, then POST.

Upon a 2xx response the successReceiver will be populated with the returned JSON
data, and "true,nil" will be returned. If there is an error response from the server
errorReceiver will be populated, and "false,nil" will be returned.

In the event that an error occured communicating with the server, or some other
unforseen error occurs, "false,error" is returned.
*/
func (c JSONClient) POST(path string, body, successReceiver, errorReceiver interface{}) (bool, error) {
	var (
		err      error
		request  *http.Request
		response *http.Response
	)

	if request, err = c.createRequest(path, "POST", c.authorization, body); err != nil {
		return false, err
	}

	if response, err = c.HTTPClient.Do(request); err != nil {
		return false, fmt.Errorf("error executing request: %w", err)
	}

	defer response.Body.Close()
	return responsegetter.Get(response, successReceiver, errorReceiver)
}

/*
PUT performs an HTTP PUT operation. You provide a path, which should exclude the
the TLD, as this is defined in BaseURL. If your request requires authorization
call WithAuthorization first, then PUT.

Upon a 2xx response the successReceiver will be populated with the returned JSON
data, and "true,nil" will be returned. If there is an error response from the server
errorReceiver will be populated, and "false,nil" will be returned.

In the event that an error occured communicating with the server, or some other
unforseen error occurs, "false,error" is returned.
*/
func (c JSONClient) PUT(path string, body, successReceiver, errorReceiver interface{}) (bool, error) {
	var (
		err      error
		request  *http.Request
		response *http.Response
	)

	if request, err = c.createRequest(path, "PUT", c.authorization, body); err != nil {
		return false, err
	}

	if response, err = c.HTTPClient.Do(request); err != nil {
		return false, fmt.Errorf("error executing request: %w", err)
	}

	defer response.Body.Close()
	return responsegetter.Get(response, successReceiver, errorReceiver)
}

/*
WithAuthorization returns an instance of RESTClient with an authorization
header set
*/
func (c JSONClient) WithAuthorization(auth string) RESTClient {
	return JSONClient{
		BaseURL:    c.BaseURL,
		HTTPClient: c.HTTPClient,

		authorization: auth,
	}
}

func (c JSONClient) buildURL(path string) string {
	return fmt.Sprintf("%s%s", c.BaseURL, path)
}

func (c JSONClient) createRequest(path, method, authorization string, body interface{}) (*http.Request, error) {
	var (
		err     error
		b       []byte
		reader  *bytes.Reader
		request *http.Request
	)

	upperMethod := strings.ToUpper(method)

	if upperMethod != "GET" && body != nil {
		if b, err = json.Marshal(body); err != nil {
			return request, fmt.Errorf("error marshaling body to JSON when creating HTTP request: %w", err)
		}

		reader = bytes.NewReader(b)
	}

	if request, err = http.NewRequest(upperMethod, c.buildURL(path), reader); err != nil {
		return request, fmt.Errorf("error creating HTTP request: %w", err)
	}

	request.Header.Add("Content-Type", "application/json")

	if authorization != "" {
		request.Header.Add("Authorization", authorization)
	}

	return request, nil
}

func (c JSONClient) getBody(response *http.Response, receiver interface{}) error {
	b, _ := io.ReadAll(response.Body)
	return json.Unmarshal(b, receiver)
}
