# REST Client

This package provides a set of services and mocks for working with and testing REST services. An interface is provided to abstract the type of data you work with, and a concrete JSON-data implementation is provided.

## HTTPClientInterface

**HTTPClientInterface** provides an interface to the Go HTTP client. This allows your services to mock HTTP interactions for unit testing.

```go
type HTTPClientInterface interface {
	Do(req *http.Request) (*http.Response, error)
}
```

As you can see this interface matches the Go *http.Client* struct, allowing you to use it to fulfill services that depend on the **HTTPClientInterface**.

## JSONClient

**JSONClient** is used to communicate with REST APIs that deal with JSON inputs and outputs. This service implements the *RESTClient* interface.

```go
httpClient := &http.Client{} // Here you could use MockHTTPClient for unit tests
client := restclient.NewJSONClient("http://localhost:8080", httpClient)
```

If you want your requests to include an authorization header that would look like this.

```go
httpClient := &http.Client{} // Here you could use MockHTTPClient for unit tests
client := restclient.NewJSONClient("http://localhost:8080", httpClient).WithAuthorization("Bearer " + token)
```

### DELETE

DELETE performs an HTTP DELETE operation. You provide a path, which should exclude the TLD, as this is defined in BaseURL. 

If some type of error occurs when creating the request, executing the request,
or unmarshalling the response, err will be populated. The http.Response is returned so callers can inspect the status.

```go
successResponse := SuccessStruct{}
errorResponse := ErrorStruct{}

response, err := client.DELETE("/thing", &successResponse, &errorResponse)

if err != nil {
	// Handle the error here
}

if response.Status > 299 {
	// Uh oh, non-success response. errorResponse contains whatever
	// JSON came back
}

// successResponse contains whatever JSON came back from success
```

### GET

GET performs an HTTP GET operation. You provide a path, which should exclude the TLD, as this is defined in BaseURL.

If some type of error occurs when creating the request, executing the request,
or unmarshalling the response, err will be populated. The http.Response is returned so callers can inspect the status.

```go
successResponse := SuccessStruct{}
errorResponse := ErrorStruct{}

response, err := client.GET("/thing", &successResponse, &errorResponse)

if err != nil {
	// Handle the error here
}

if response.Status > 299 {
	// Uh oh, non-success response. errorResponse contains whatever
	// JSON came back
}

// successResponse contains whatever JSON came back from success
```

## POST

POST performs an HTTP POST operation. You provide a path, which should exclude the TLD, as this is defined in BaseURL. The body of the post is an interface, allowing you to pass whatever you wish, and it will be serialized to JSON.

If some type of error occurs when creating the request, executing the request,
or unmarshalling the response, err will be populated. The http.Response is returned so callers can inspect the status.

```go
successResponse := SuccessStruct{}
errorResponse := ErrorStruct{}

body := SomeStruct{
	SomeKey1: "somevalue1",
	SomeKey2: 2,
}

response, err := client.POST("/thing", body, &successResponse, &errorResponse)

if err != nil {
	// Handle the error here
}

if response.Status > 299 {
	// Uh oh, non-success response. errorResponse contains whatever
	// JSON came back
}

// successResponse contains whatever JSON came back from success
```

## PUT

PUT performs an HTTP PUT operation. You provide a path, which should exclude the TLD, as this is defined in BaseURL. The body of the put is an interface, allowing you to pass whatever you wish, and it will be serialized to JSON.

If some type of error occurs when creating the request, executing the request,
or unmarshalling the response, err will be populated. The http.Response is returned so callers can inspect the status.

```go
successResponse := SuccessStruct{}
errorResponse := ErrorStruct{}

body := SomeStruct{
	SomeKey1: "somevalue1",
	SomeKey2: 2,
}

response, err := client.PUT("/thing", body, &successResponse, &errorResponse)

if err != nil {
	// Handle the error here
}

if response.Status > 299 {
	// Uh oh, non-success response. errorResponse contains whatever
	// JSON came back
}

// successResponse contains whatever JSON came back from success
```

## WithAuthorization

**WithAuthorization** returns a copy of this JSONClient with the *Authorization* header set. Subsequent HTTP requests will send `Authorization: <auth>`, with *auth* being whatever you set in the string argument.

```go
httpClient := &http.Client{} // Here you could use MockHTTPClient for unit tests
client := restclient.NewJSONClient("http://localhost:8080", httpClient).WithAuthorization("Bearer " + token)
```

