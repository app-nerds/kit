/*
Copyright 2017 AppNinjas LLC

Licensed under the Apache License, Version 2.0 (the "License"); you may not use this file except in compliance with the License. You may obtain a copy of the License at

http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software distributed under the License is distributed on an "AS IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied. See the License for the specific language governing permissions and limitations under the License.
*/
package logging

import "strings"

/*
LogFormat describes how to format log messages
*/
type LogFormat int

/*
Constants for the available log formats
*/
const (
	LOG_FORMAT_SIMPLE LogFormat = iota
	LOG_FORMAT_JSON
)

var logFormatNames = map[LogFormat]string{
	LOG_FORMAT_SIMPLE: "Simple",
	LOG_FORMAT_JSON:   "JSON",
}

/*
String returns the friendly name of a specified log format
*/
func (format LogFormat) String() string {
	return logFormatNames[format]
}

/*
StringToLogFormat converts a specified string to a LogFormat. If the string does not
match a specific log type the SIMPLE is returned.
*/
func StringToLogFormat(logFormatName string) LogFormat {
	for logFormat, stringValue := range logFormatNames {
		if strings.ToLower(stringValue) == strings.ToLower(logFormatName) {
			return logFormat
		}
	}

	return LOG_FORMAT_SIMPLE
}
