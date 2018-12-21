package restclient

import (
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"net/url"
	"strings"
)

/*
IRestClient defines an interface for working with RESTful endpoints
*/
type IRestClient interface {
	BuildURL(query string) string
	ConvertMapToQueryString(m HTTPParameters) string
	ConvertParametersToReader(parameters HTTPParameters) io.Reader
	CreateRequest(query, method string, body io.Reader, authorization string) (*http.Request, error)
	ExecuteRequest(request *http.Request) (*http.Response, error)
	GetResponseBytes(response *http.Response) []byte
	GetResponseJSON(response *http.Response, receiver interface{}) error
	GetResponseString(response *http.Response) string
}

/*
RestClient provides methods for working with RESTful endpoints
*/
type RestClient struct {
	BaseURL    string
	HTTPClient IHTTPClient
}

/*
HTTPParameters is a map of strings representing HTTP parameters. These
are useful for POST and PUT operations
*/
type HTTPParameters map[string]string

/*
BuildURL takes a query and builds a full URL
*/
func (c *RestClient) BuildURL(query string) string {
	return fmt.Sprintf("%s/%s", c.BaseURL, query)
}

/*
ConvertMapToQueryString takes a map and converts it to a key/value
pairs
*/
func (c *RestClient) ConvertMapToQueryString(m HTTPParameters) string {
	result := &strings.Builder{}

	index := 0
	len := len(m)

	result.WriteString("")

	for key, value := range m {
		result.WriteString(key)
		result.WriteString("=")
		result.WriteString(url.QueryEscape(value))

		if index < len-1 {
			result.WriteString("&")
		}

		index++
	}

	return result.String()
}

/*
ConvertParametersToReader takes a set of HTTPParameters and converts them
to key/value pairs suitable for use in an HTTP request
*/
func (c *RestClient) ConvertParametersToReader(parameters HTTPParameters) io.Reader {
	parametersString := c.ConvertMapToQueryString(parameters)
	return strings.NewReader(parametersString)
}

/*
CreateRequest builds a request object for a provided url and
HTTP method
*/
func (c *RestClient) CreateRequest(query, method string, body io.Reader, authorization string) (*http.Request, error) {
	request, err := http.NewRequest(method, c.BuildURL(query), body)

	if err != nil {
		return nil, err
	}

	if authorization != "" {
		request.Header.Add("Authorization", authorization)
	}

	return request, nil
}

/*
ExecuteRequest execute an arbitrary request to GoBucket
*/
func (c *RestClient) ExecuteRequest(request *http.Request) (*http.Response, error) {
	var err error
	var response *http.Response

	response, err = c.HTTPClient.Do(request)
	return response, err
}

/*
GetResponseBytes returns the response content from the body as a byte slice
*/
func (c *RestClient) GetResponseBytes(response *http.Response) []byte {
	result, _ := ioutil.ReadAll(response.Body)
	return result
}

/*
GetResponseJSON converts response bytes and unmarshals to a receiver object
*/
func (c *RestClient) GetResponseJSON(response *http.Response, receiver interface{}) error {
	bytes := c.GetResponseBytes(response)
	return json.Unmarshal(bytes, receiver)
}

/*
GetResponseString returns the response content from the body as a string
*/
func (c *RestClient) GetResponseString(response *http.Response) string {
	result := c.GetResponseBytes(response)
	return string(result)
}
