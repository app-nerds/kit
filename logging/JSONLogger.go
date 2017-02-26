/*
Copyright 2017 AppNinjas LLC

Licensed under the Apache License, Version 2.0 (the "License"); you may not use this file except in compliance with the License. You may obtain a copy of the License at

http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software distributed under the License is distributed on an "AS IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied. See the License for the specific language governing permissions and limitations under the License.
*/
package logging

import (
	"fmt"
	"log"
	"strings"

	"github.com/fatih/color"
)

/*
JSONLogger is a console logger that has a format of `{"applicationName": "{ApplicationName}", "type": "{Type}"", "message": "{Message}"}`
*/
type JSONLogger struct {
	Logger
}

/*
NewJSONLogger returns an instance of an ILogger interface
set to the JSON logger format
*/
func NewJSONLogger(applicationName string, minimumLogLevel LogType) *JSONLogger {
	return &JSONLogger{
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
func (logger *JSONLogger) Debugf(message string, args ...interface{}) {
	logger.writeLogf(DEBUG, message, args...)
}

/*
DisableColors turns of console coloring
*/
func (logger *JSONLogger) DisableColors() {
	logger.colorEnabled = false
}

/*
EnableColors turns on console coloring
*/
func (logger *JSONLogger) EnableColors() {
	logger.colorEnabled = true
}

/*
Errorf writes a formatted error entry to the log
*/
func (logger *JSONLogger) Errorf(message string, args ...interface{}) {
	logger.writeLogf(ERROR, message, args...)
}

/*
Infof writes a formatted info entry to the log
*/
func (logger *JSONLogger) Infof(message string, args ...interface{}) {
	logger.writeLogf(INFO, message, args...)
}

func (logger *JSONLogger) writeLogf(logType LogType, message string, args ...interface{}) {
	logLevelInt := int(logType)
	logColor := logType.Color()

	if logger.colorEnabled {
		color.Set(logColor)
	}

	if logLevelInt >= logger.logLevelInt {
		log.SetPrefix("")

		formattedMessage := fmt.Sprintf(message, args...)
		formattedMessage = logger.escape(formattedMessage)
		log.Printf(fmt.Sprintf(`{"applicationName": "%s", "type": "%s", "message": "%s"}`, logger.ApplicationName, logType.String(), formattedMessage))
	}

	if logger.colorEnabled {
		color.Unset()
	}
}

func (logger *JSONLogger) escape(s string) string {
	return strings.Replace(s, "\"", "\\\"", -1)
}
