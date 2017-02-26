/*
Copyright 2017 AppNinjas LLC

Licensed under the Apache License, Version 2.0 (the "License"); you may not use this file except in compliance with the License. You may obtain a copy of the License at

http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software distributed under the License is distributed on an "AS IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied. See the License for the specific language governing permissions and limitations under the License.
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
