# logging Package

The **logging** package provides objects and methods for logging output to the standard output console.
It allows consumers to output in various formats, including text and JSON. It also supports
log levels and color.

## Package Methods

Below is a reference of exported methods in this package.

### LogFactory(logFormat LogFormat, applicationName string, minimumLogLevel LogType) ILogger
`LogFactory` should be your primary point of entry. This is a factory method that,
when told what format, log level, and application name is, will return the correct
logger object. For example, to get a simple logger with a minimum log level output
of `INFO`:

```golang
var logger logging.ILogger
logger = logging.LogFactory(logging.LOG\_FORMAT\_SIMPLE, "Test Application", logging.INFO)
```

### NewJSONLogger(applicationName string, minimumLogLevel LogType) *JSONLogger
`NewJSONLogger` returns an instance of a new logger that outputs in JSON format. Logs
written in JSON format output an object to the console. Here is an example of what that
looks like. Note that an actual log entry would be a single line.

```javascript
{
	"applicationName": "Test Application",
	"type": "INFO",
	"message": "Database connection established"
}
```

### NewSimpleLogger(applicationName string, minimumLogLevel LogType) *SimpleLogger
`NewSimpleLogger` returns an instance of a new logger that outputs plain text.
Here is an example of what that looks like.

```text
Test Application: INFO - Database connection established
```

### StringToLogFormat(logFormatName string) LogFormat
`StringToLogFormat` takes a string and attempts to convert it to a LogFormat type.
A log format represents the format the logger will output. Currently valid values are

* SIMPLE
* JSON

### StringToLogType(logTypeName string) LogType
`StringToLogType` takes a string and attempts to convert it to a LogType.
A log type specifies the types of log entries, such as an error, or information entry.
The following are the supported log types.

* DEBUG
* INFO
* WARN
* ERROR
* FATAL

