package restclient

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"mime/multipart"
	"net/http"

	"github.com/app-nerds/kit/v6/restclient/responsegetter"
	"github.com/sirupsen/logrus"
)

type MultipartWriterConfig struct {
	Authorization string
	BaseURL       string
	DebugMode     bool
	HTTPClient    HTTPClientInterface
	Logger        *logrus.Entry
}

type MultipartWriter struct {
	baseURL    string
	debugMode  bool
	httpClient HTTPClientInterface
	logger     *logrus.Entry

	authorization string
	writer        *multipart.Writer
	body          *bytes.Buffer
}

/*
NewMultipartWriter creates a new MultipartWriter
*/
func NewMultipartWriter(config MultipartWriterConfig) *MultipartWriter {
	result := &MultipartWriter{
		baseURL:       config.BaseURL,
		debugMode:     config.DebugMode,
		httpClient:    config.HTTPClient,
		logger:        config.Logger.WithField("component", "MultipartWriter"),
		authorization: config.Authorization,
	}

	result.body = &bytes.Buffer{}
	result.writer = multipart.NewWriter(result.body)

	return result
}

/*
AddField adds a field to this multipart form. It accepts two types: string, json.Marshaler.
For the latter, this would be any struct or type that implements the json.Marshaler
interface.
*/
func (mw *MultipartWriter) AddField(name string, body interface{}) error {
	var (
		err error
		ok  bool
		b   []byte
	)

	if _, ok = body.(string); ok {
		s, _ := body.(string)
		return mw.writer.WriteField(name, s)
	}

	if b, err = json.Marshal(body); err != nil {
		return fmt.Errorf("error converting body to JSON: %w", err)
	}

	return mw.writer.WriteField(name, string(b))
}

/*
AddFile adds a file to this multipart form.
*/
func (mw *MultipartWriter) AddFile(name, fileName string, fileContents io.Reader) error {
	var (
		err        error
		fileWriter io.Writer
	)

	if fileWriter, err = mw.writer.CreateFormFile(name, fileName); err != nil {
		return fmt.Errorf("error initializing file form field: %w", err)
	}

	if _, err = io.Copy(fileWriter, fileContents); err != nil {
		return fmt.Errorf("error copying file contents to buffer: %w", err)
	}

	return nil
}

/*
POST sends a multipart form via POST. The fields being sent should be added prior to calling
POST by using the AddField and AddFile methods.
*/
func (mw *MultipartWriter) POST(path string, successReceiver, errorReceiver interface{}) (*http.Response, error) {
	var (
		err      error
		request  *http.Request
		response *http.Response
	)

	_ = mw.writer.Close()

	if request, err = http.NewRequest(http.MethodPost, mw.buildURL(path), mw.body); err != nil {
		return response, fmt.Errorf("error creating multipart POST request: %w", err)
	}

	contentType := mw.writer.FormDataContentType()
	request.Header.Add("Content-Type", contentType)

	if mw.authorization != "" {
		request.Header.Add("Authorization", mw.authorization)
	}

	if response, err = mw.httpClient.Do(request); err != nil {
		return response, fmt.Errorf("error executing multipart POST request: %w", err)
	}

	defer response.Body.Close()
	return response, responsegetter.Get(response, successReceiver, errorReceiver, mw.logger, mw.debugMode)
}

/*
PUT sends a multipart form via PUT. The fields being sent should be added prior to calling
PUT by using the AddField and AddFile methods.
*/
func (mw *MultipartWriter) PUT(path string, successReceiver, errorReceiver interface{}) (*http.Response, error) {
	var (
		err      error
		request  *http.Request
		response *http.Response
	)

	_ = mw.writer.Close()

	if request, err = http.NewRequest(http.MethodPut, mw.buildURL(path), mw.body); err != nil {
		return response, fmt.Errorf("error creating multipart POST request: %w", err)
	}

	contentType := mw.writer.FormDataContentType()
	request.Header.Add("Content-Type", contentType)

	if mw.authorization != "" {
		request.Header.Add("Authorization", mw.authorization)
	}

	if response, err = mw.httpClient.Do(request); err != nil {
		return response, fmt.Errorf("error executing multipart PUT request: %w", err)
	}

	defer response.Body.Close()
	return response, responsegetter.Get(response, successReceiver, errorReceiver, mw.logger, mw.debugMode)
}

func (mw *MultipartWriter) buildURL(path string) string {
	return fmt.Sprintf("%s%s", mw.baseURL, path)
}
