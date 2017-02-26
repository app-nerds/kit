/*
Copyright 2017 AppNinjas LLC

Licensed under the Apache License, Version 2.0 (the "License"); you may not use this file except in compliance with the License. You may obtain a copy of the License at

http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software distributed under the License is distributed on an "AS IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied. See the License for the specific language governing permissions and limitations under the License.
*/
package logging

import (
	"strings"

	"github.com/fatih/color"
)

/*
LogType represents a type and level of logging
*/
type LogType int

/*
Constants for the type and levels of logging
*/
const (
	NONE LogType = iota
	DEBUG
	INFO
	WARN
	ERROR
	FATAL
)

var logTypeNames = map[LogType]string{
	NONE:  "None",
	DEBUG: "DEBUG",
	INFO:  "INFO",
	WARN:  "WARNING",
	ERROR: "ERROR",
	FATAL: "FATAL",
}

var logTypeColors = map[LogType]color.Attribute{
	DEBUG: color.FgGreen,
	INFO:  color.FgWhite,
	WARN:  color.FgYellow,
	ERROR: color.FgRed,
	FATAL: color.FgRed,
}

/*
Color returns the color attribute for this log type
*/
func (logType LogType) Color() color.Attribute {
	return logTypeColors[logType]
}

/*
String returns the friendly name of a specified log type/level
*/
func (logType LogType) String() string {
	return logTypeNames[logType]
}

/*
StringToLogType converts a specified string to a LogType. If the string does not
match a specific log type the NONE is returned.
*/
func StringToLogType(logTypeName string) LogType {
	for logType, stringValue := range logTypeNames {
		if strings.ToLower(stringValue) == strings.ToLower(logTypeName) {
			return logType
		}
	}

	return NONE
}
