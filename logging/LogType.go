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
