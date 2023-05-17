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
	"github.com/sirupsen/logrus"
)

/*
JSONClientConfig is used to configure a JSONClient struct
*/
type JSONClientConfig struct {
	BaseURL         string
	DebugMode       bool
	HTTPClient      HTTPClientInterface
	Logger          *logrus.Entry
	CustomHeaders   map[string]string
	OmitContentType bool
}

/*
JSONClient provides a set of methods for working with RESTful endpoints that accept
and return JSON data
*/
type JSONClient struct {
	baseURL         string
	debugMode       bool
	httpClient      HTTPClientInterface
	logger          *logrus.Entry
	customHeaders   map[string]string
	omitContentType bool

	authorization string
}

/*
NewJSONClient creates a new JSON-based REST client
*/
func NewJSONClient(config JSONClientConfig) JSONClient {
	return JSONClient{
		baseURL:         config.BaseURL,
		debugMode:       config.DebugMode,
		httpClient:      config.HTTPClient,
		logger:          config.Logger.WithField("component", "JSONClient"),
		customHeaders:   config.CustomHeaders,
		omitContentType: config.OmitContentType,

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

	if response, err = c.httpClient.Do(request); err != nil {
		return response, fmt.Errorf("error executing request: %w", err)
	}

	defer response.Body.Close()
	return response, responsegetter.Get(response, successReceiver, errorReceiver, c.logger, c.debugMode)
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

	if response, err = c.httpClient.Do(request); err != nil {
		return response, fmt.Errorf("error executing request: %w", err)
	}

	defer response.Body.Close()

	return response, responsegetter.Get(response, successReceiver, errorReceiver, c.logger, c.debugMode)
}

/*
NewMultipartWriter returns a MultipartWriter. This is used to POST and PUT multipart forms.
Use this when you need to POST or PUT files, for example.
*/
func (c JSONClient) NewMultipartWriter() *MultipartWriter {
	config := MultipartWriterConfig{
		Authorization: c.authorization,
		BaseURL:       c.baseURL,
		DebugMode:     c.debugMode,
		HTTPClient:    c.httpClient,
		Logger:        c.logger,
	}

	return NewMultipartWriter(config)
}

/*
POST performs an HTTP POST operation. You provide a path, which should exclude
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

	if response, err = c.httpClient.Do(request); err != nil {
		return response, fmt.Errorf("error executing request: %w", err)
	}

	defer response.Body.Close()
	return response, responsegetter.Get(response, successReceiver, errorReceiver, c.logger, c.debugMode)
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

	if response, err = c.httpClient.Do(request); err != nil {
		return response, fmt.Errorf("error executing request: %w", err)
	}

	defer response.Body.Close()
	return response, responsegetter.Get(response, successReceiver, errorReceiver, c.logger, c.debugMode)
}

/*
WithAuthorization returns an instance of RESTClient with an authorization
header set
*/
func (c JSONClient) WithAuthorization(auth string) RESTClient {
	return JSONClient{
		baseURL:    c.baseURL,
		debugMode:  c.debugMode,
		httpClient: c.httpClient,
		logger:     c.logger,

		authorization: auth,
	}
}

func (c JSONClient) buildURL(path string) string {
	return fmt.Sprintf("%s%s", c.baseURL, path)
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

	if !c.omitContentType {
		request.Header.Add("Content-Type", "application/json")
	}

	if authorization != "" {
		request.Header.Add("Authorization", authorization)
	}

	if c.customHeaders != nil {
		for k, v := range c.customHeaders {
			request.Header.Add(k, v)
		}
	}

	if c.debugMode {
		c.logger.WithFields(logrus.Fields{
			"method":        upperMethod,
			"url":           u,
			"authorization": authorization,
			"customHeaders": c.customHeaders,
		}).Info("request created")
	}

	return request, nil
}

func (c JSONClient) getBody(response *http.Response, receiver interface{}) error {
	b, _ := io.ReadAll(response.Body)
	return json.Unmarshal(b, receiver)
}
