package logging

import (
	"log"

	"github.com/fatih/color"
)

/*
SimpleLogger is a basic console logger that has a format of `{ApplicationName}: {Type} - {Message}`
*/
type SimpleLogger struct {
	Logger
}

/*
NewSimpleLogger returns an instance of an ILogger interface
set to the simple logger format
*/
func NewSimpleLogger(applicationName string, minimumLogLevel LogType) ILogger {
	return &SimpleLogger{
		Logger: Logger{
			ApplicationName: applicationName,
			LogLevel:        minimumLogLevel,

			colorEnabled: false,
			logLevelInt:  int(minimumLogLevel),
		},
	}
}

/*
Debugf writes a formatted debug entry to the log
*/
func (logger *SimpleLogger) Debugf(message string, args ...interface{}) {
	logger.writeLogf(DEBUG, message, args...)
}

/*
DisableColors turns of console coloring
*/
func (logger *SimpleLogger) DisableColors() {
	logger.colorEnabled = false
}

/*
EnableColors turns on console coloring
*/
func (logger *SimpleLogger) EnableColors() {
	logger.colorEnabled = true
}

/*
Errorf writes a formatted error entry to the log
*/
func (logger *SimpleLogger) Errorf(message string, args ...interface{}) {
	logger.writeLogf(ERROR, message, args...)
}

/*
Infof writes a formatted info entry to the log
*/
func (logger *SimpleLogger) Infof(message string, args ...interface{}) {
	logger.writeLogf(INFO, message, args...)
}

func (logger *SimpleLogger) writeLogf(logType LogType, message string, args ...interface{}) {
	logLevelInt := int(logType)
	logColor := logType.Color()

	if logger.colorEnabled {
		color.Set(logColor)
	}

	if logLevelInt >= logger.logLevelInt {
		log.SetPrefix(logger.ApplicationName + ": " + logType.String() + " - ")
		log.Printf(message+"\n", args...)
	}

	if logger.colorEnabled {
		color.Unset()
	}
}
