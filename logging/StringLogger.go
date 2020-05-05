/*
 * Copyright (c) 2020. App Nerds LLC. All rights reserved
 */
package logging

import "fmt"

/*
StringLogger is a logger that stores log entries in an array that has a format of `{ApplicationName}: {Type} - {Message}`.
This is useful for things like unit tests
*/
type StringLogger struct {
	Logger

	Messages []string
}

/*
NewStringLogger returns an instance of an ILogger interface
set to the simple logger format
*/
func NewStringLogger(applicationName string, minimumLogLevel LogType) *StringLogger {
	return &StringLogger{
		Logger: Logger{
			ApplicationName: applicationName,
			LogLevel:        minimumLogLevel,

			colorEnabled: false,
			logLevelInt:  int(minimumLogLevel),
		},
		Messages: make([]string, 0, 100),
	}
}

/*
Debugf writes a formatted debug entry to the log
*/
func (logger *StringLogger) Debugf(message string, args ...interface{}) {
	logger.writeLogf(DEBUG, message, args...)
}

/*
DisableColors turns of console coloring
*/
func (logger *StringLogger) DisableColors() {
	logger.colorEnabled = false
}

/*
EnableColors turns on console coloring
*/
func (logger *StringLogger) EnableColors() {
	logger.colorEnabled = true
}

/*
Errorf writes a formatted error entry to the log
*/
func (logger *StringLogger) Errorf(message string, args ...interface{}) {
	logger.writeLogf(ERROR, message, args...)
}

/*
Infof writes a formatted info entry to the log
*/
func (logger *StringLogger) Infof(message string, args ...interface{}) {
	logger.writeLogf(INFO, message, args...)
}

func (logger *StringLogger) writeLogf(logType LogType, message string, args ...interface{}) {
	logLevelInt := int(logType)

	if logLevelInt >= logger.logLevelInt {
		prefix := fmt.Sprintf("%s: %s - ", logger.ApplicationName, logType.String())
		message := fmt.Sprintf(message, args...)

		wholeMessage := fmt.Sprintf("%s%s", prefix, message)
		logger.Messages = append(logger.Messages, wholeMessage)
	}
}
