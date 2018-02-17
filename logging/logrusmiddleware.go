package logging

import (
	"io"
	"strconv"
	"time"

	"github.com/labstack/echo"
	"github.com/labstack/gommon/log"
	"github.com/sirupsen/logrus"
)

type MiddlewareLogger struct {
	*logrus.Logger
}

func (l MiddlewareLogger) Level() log.Lvl {
	switch l.Logger.Level {
	case logrus.DebugLevel:
		return log.DEBUG
	case logrus.WarnLevel:
		return log.WARN
	case logrus.ErrorLevel:
		return log.ERROR
	case logrus.InfoLevel:
		return log.INFO
	default:
		l.Panic("Invalid level")
	}

	return log.OFF
}

func (l MiddlewareLogger) SetPrefix(s string) {
	// TODO
}

func (l MiddlewareLogger) Prefix() string {
	// TODO.  Is this even valid?  I'm not sure it can be translated since
	// logrus uses a Formatter interface.  Which seems to me to probably be
	// a better way to do it.
	return ""
}

func (l MiddlewareLogger) SetLevel(lvl log.Lvl) {
	switch lvl {
	case log.DEBUG:
		logrus.SetLevel(logrus.DebugLevel)
	case log.WARN:
		logrus.SetLevel(logrus.WarnLevel)
	case log.ERROR:
		logrus.SetLevel(logrus.ErrorLevel)
	case log.INFO:
		logrus.SetLevel(logrus.InfoLevel)
	default:
		l.Panic("Invalid level")
	}
}

func (l MiddlewareLogger) Output() io.Writer {
	return l.Out
}

func (l MiddlewareLogger) SetOutput(w io.Writer) {
	logrus.SetOutput(w)
}

func (l MiddlewareLogger) Printj(j log.JSON) {
	logrus.WithFields(logrus.Fields(j)).Print()
}

func (l MiddlewareLogger) Debugj(j log.JSON) {
	logrus.WithFields(logrus.Fields(j)).Debug()
}

func (l MiddlewareLogger) Infoj(j log.JSON) {
	logrus.WithFields(logrus.Fields(j)).Info()
}

func (l MiddlewareLogger) Warnj(j log.JSON) {
	logrus.WithFields(logrus.Fields(j)).Warn()
}

func (l MiddlewareLogger) Errorj(j log.JSON) {
	logrus.WithFields(logrus.Fields(j)).Error()
}

func (l MiddlewareLogger) Fatalj(j log.JSON) {
	logrus.WithFields(logrus.Fields(j)).Fatal()
}

func (l MiddlewareLogger) Panicj(j log.JSON) {
	logrus.WithFields(logrus.Fields(j)).Panic()
}

type LogrusMiddleware struct {
	logger *logrus.Logger
}

func (l *LogrusMiddleware) logrusMiddlewareHandler(ctx echo.Context, next echo.HandlerFunc) error {
	request := ctx.Request()
	response := ctx.Response()

	start := time.Now()

	if err := next(ctx); err != nil {
		ctx.Error(err)
	}

	stop := time.Now()

	p := request.URL.Path
	if p == "" {
		p = "/"
	}

	bytesIn := request.Header.Get(echo.HeaderContentLength)
	if bytesIn == "" {
		bytesIn = "0"
	}

	l.logger.WithFields(map[string]interface{}{
		"time":          time.Now().Format(time.RFC3339),
		"remote_ip":     ctx.RealIP(),
		"uri":           request.RequestURI,
		"method":        request.Method,
		"status":        response.Status,
		"latency":       strconv.FormatInt(stop.Sub(start).Nanoseconds()/1000, 10),
		"latency_human": stop.Sub(start).String(),
	}).Info("Handled request")

	return nil
}

func (l *LogrusMiddleware) middleware(next echo.HandlerFunc) echo.HandlerFunc {
	return func(ctx echo.Context) error {
		return l.logrusMiddlewareHandler(ctx, next)
	}
}

func Hook() echo.MiddlewareFunc {
	l := &LogrusMiddleware{
		logger: logrus.New().WithField("who", "Request Logger").Logger,
	}

	return l.middleware
}

func HookWithLogger(existingLogger *logrus.Logger) echo.MiddlewareFunc {
	l := &LogrusMiddleware{
		logger: existingLogger,
	}

	return l.middleware
}

func HookWithExisting(existingLogger *logrus.Entry) echo.MiddlewareFunc {
	l := &LogrusMiddleware{
		logger: existingLogger.Logger,
	}

	return l.middleware
}
