/*
 * Copyright (c) 2021. App Nerds LLC. All rights reserved
 */

package responsegetter

import (
	"encoding/json"
	"io"
	"net/http"
	"strings"

	"github.com/sirupsen/logrus"
)

/*
Get retrieves the HTTP response body. If the response is successful the body
is written into "successReceiver". If not it is written into "errorReceiver".
The type of value that is written depends on the response Content-Type.
*/
func Get(response *http.Response, successReceiver, errorReceiver interface{}, logger *logrus.Entry, debugMode bool) error {
	contentType := response.Header.Get("Content-Type")
	p := successReceiver

	if !isSuccess(response) {
		p = errorReceiver
	}

	if strings.Contains(contentType, "application/json") {
		return getJSON(response, p, logger.WithField("contentType", contentType), debugMode)
	}

	return getString(response, p, logger.WithField("contentType", contentType), debugMode)
}

func isSuccess(response *http.Response) bool {
	return response.StatusCode >= 200 && response.StatusCode < 300
}

func getJSON(response *http.Response, receiver interface{}, logger *logrus.Entry, debugMode bool) error {
	b, _ := io.ReadAll(response.Body)

	if debugMode {
		logger.WithFields(logrus.Fields{
			"body":       string(b),
			"statusCode": response.StatusCode,
		}).Info("retrieved JSON")
	}

	return json.Unmarshal(b, receiver)
}

func getString(response *http.Response, receiver interface{}, logger *logrus.Entry, debugMode bool) error {
	b, _ := io.ReadAll(response.Body)

	p := receiver.(*string)
	*p = string(b)

	if debugMode {
		logger.WithFields(logrus.Fields{
			"body":       string(b),
			"statusCode": response.StatusCode,
		}).Info("retrieved string")
	}

	return nil
}
