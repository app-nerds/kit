package logging

/*
ILogger is an interface that describes how a logger should behave
*/
type ILogger interface {
	Debugf(message string, args ...interface{})
	DisableColors()
	EnableColors()
	Errorf(message string, args ...interface{})
	Infof(message string, args ...interface{})
}
