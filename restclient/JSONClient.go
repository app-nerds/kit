/*
 * Copyright (c) 2021. App Nerds LLC. All rights reserved
 */

package restclient

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strings"

	"github.com/app-nerds/kit/v6/restclient/responsegetter"
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
DELETE performs an HTTP DELETE operation. You provide a path, which should exclude 
the TLD, as this is defined in BaseURL. If your request requires authorization
call WithAuthorization first, then DELETE.

If some type of error occurs when creating the request, executing the request,
or unmarshalling the response, err will be populated. The http.Response is returned
so callers can inspect the status.
*/
func (c JSONClient) DELETE(path string, successReceiver, errorReceiver interface{}) (*http.Response, error) {
	var (
		err      error
		request  *http.Request
		response *http.Response
	)

	if request, err = c.createRequest(path, "DELETE", c.authorization, nil); err != nil {
		return response, fmt.Errorf("error creating HTTP request: %w", err)
	}

	if response, err = c.HTTPClient.Do(request); err != nil {
		return response, fmt.Errorf("error executing request: %w", err)
	}

	defer response.Body.Close()
	return response, responsegetter.Get(response, successReceiver, errorReceiver)
}

/*
GET performs an HTTP GET operation. You provide a path, which should exclude the
the TLD, as this is defined in BaseURL. If your request requires authorization
call WithAuthorization first, then GET.

If some type of error occurs when creating the request, executing the request,
or unmarshalling the response, err will be populated. The http.Response is returned
so callers can inspect the status.
*/
func (c JSONClient) GET(path string, successReceiver, errorReceiver interface{}) (*http.Response, error) {
	var (
		err      error
		request  *http.Request
		response *http.Response
	)

	if request, err = c.createRequest(path, "GET", c.authorization, nil); err != nil {
		return response, fmt.Errorf("error creating HTTP request: %w", err)
	}

	if response, err = c.HTTPClient.Do(request); err != nil {
		return response, fmt.Errorf("error executing request: %w", err)
	}

	defer response.Body.Close()
	return response, responsegetter.Get(response, successReceiver, errorReceiver)
}

/*
POST performs an HTTP POST operation. You provide a path, which should exclude the
the TLD, as this is defined in BaseURL. If your request requires authorization
call WithAuthorization first, then POST.

If some type of error occurs when creating the request, executing the request,
or unmarshalling the response, err will be populated. The http.Response is returned
so callers can inspect the status.
*/
func (c JSONClient) POST(path string, body, successReceiver, errorReceiver interface{}) (*http.Response, error) {
	var (
		err      error
		request  *http.Request
		response *http.Response
	)

	if request, err = c.createRequest(path, "POST", c.authorization, body); err != nil {
		return response, fmt.Errorf("error creating HTTP request: %w", err)
	}

	if response, err = c.HTTPClient.Do(request); err != nil {
		return response, fmt.Errorf("error executing request: %w", err)
	}

	defer response.Body.Close()
	return response, responsegetter.Get(response, successReceiver, errorReceiver)
}

/*
PUT performs an HTTP PUT operation. You provide a path, which should exclude the
the TLD, as this is defined in BaseURL. If your request requires authorization
call WithAuthorization first, then PUT.

If some type of error occurs when creating the request, executing the request,
or unmarshalling the response, err will be populated. The http.Response is returned
so callers can inspect the status.
*/
func (c JSONClient) PUT(path string, body, successReceiver, errorReceiver interface{}) (*http.Response, error) {
	var (
		err      error
		request  *http.Request
		response *http.Response
	)

	if request, err = c.createRequest(path, "PUT", c.authorization, body); err != nil {
		return response, fmt.Errorf("error creating HTTP request: %w", err)
	}

	if response, err = c.HTTPClient.Do(request); err != nil {
		return response, fmt.Errorf("error executing request: %w", err)
	}

	defer response.Body.Close()
	return response, responsegetter.Get(response, successReceiver, errorReceiver)
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
	u := c.buildURL(path)

	if upperMethod != "GET" && body != nil {
		if b, err = json.Marshal(body); err != nil {
			return request, fmt.Errorf("error marshaling body to JSON when creating HTTP request: %w", err)
		}

		reader = bytes.NewReader(b)
		if request, err = http.NewRequest(upperMethod, u, reader); err != nil {
			return request, fmt.Errorf("error creating HTTP request: %w", err)
		}
	} else if upperMethod == "GET" {
		if request, err = http.NewRequest(upperMethod, u, nil); err != nil {
			return request, fmt.Errorf("error creating HTTP request: %w", err)
		}
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
