/*
 * Copyright (c) 2020. App Nerds LLC. All rights reserved
 */
package logging

import "log"

/*
Logger represents the basic instance of a logging object. Other,
more specific loggers will use this
*/
type Logger struct {
	ApplicationName string
	LogLevel        LogType
	LogFormat       LogFormat

	colorEnabled bool
	logLevelInt  int
}

/*
LogFactory returns a logger in the required format
*/
func LogFactory(logFormat LogFormat, applicationName string, minimumLogLevel LogType) ILogger {
	switch logFormat {
	case LOG_FORMAT_SIMPLE:
		return NewSimpleLogger(applicationName, minimumLogLevel)

	case LOG_FORMAT_JSON:
		return NewJSONLogger(applicationName, minimumLogLevel)

	default:
		return NewSimpleLogger(applicationName, minimumLogLevel)
	}
}

func (logger *Logger) writeLogf(logType LogType, message string, args ...interface{}) {
	log.Printf("Not implemented")
}
