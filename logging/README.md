# Logging

This package provides logging-related services.

## Fireplace

Fireplace is a custom log aggregation server written by App Nerds. It provides a RESTful endpoint for writing log entries which in turn stores them in a MongoDB database. There is also a viewer application that can be used to browse and search log entries. This package provides a method to create a new instance of a [Logrus](https://github.com/sirupsen/logrus) logger that is connected to a Fireplace server.

```bash
go get github.com/app-nerds/kit/v6/logging
```

To begin you'll need two pieces of information:

* The URL to a Fireplace Server (full URL. e.g. https://someurl.com:8999)
* The password to the Fireplace Server

### Example

```go
logger := logging.NewFireplaceLogger("My application", "info", "https://someurl.com:8999", "password", nil)
```
