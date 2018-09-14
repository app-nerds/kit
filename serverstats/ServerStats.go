package serverstats

import (
	"container/ring"
	"net/http"
	"strconv"
	"sync"
	"time"

	"github.com/labstack/echo"
)

/*
ServerStats tracks general server statistics. This includes information
about uptime, response times and counts, and requests counts broken
down by HTTP status code. ServerStats is thread-safe due to a
write lock on requests, and a read lock on reads
*/
type ServerStats struct {
	Uptime        time.Time `json:"uptime"`
	RequestCount  uint64    `json:"requestCount"`
	ResponseTimes *ring.Ring
	Statuses      map[string]int `json:"statuses"`

	mutex sync.RWMutex
}

/*
NewServerStats creates a new ServerStats object
*/
func NewServerStats() *ServerStats {
	return &ServerStats{
		Uptime:        time.Now().UTC(),
		ResponseTimes: ring.New(1000),
		Statuses:      make(map[string]int),
	}
}

/*
Middleware is used to capture request and response stats. This is designed
to be used with the Echo framework
*/
func (s *ServerStats) Middleware(next echo.HandlerFunc) echo.HandlerFunc {
	return func(ctx echo.Context) error {
		var err error

		startTime := time.Now()

		if err = next(ctx); err != nil {
			ctx.Error(err)
		}

		endTime := time.Since(startTime)

		s.mutex.Lock()
		defer s.mutex.Unlock()

		s.RequestCount++

		s.ResponseTimes = s.ResponseTimes.Next()
		s.ResponseTimes.Value = endTime

		status := strconv.Itoa(ctx.Response().Status)
		s.Statuses[status]++

		return nil
	}
}

/*
Handler is an endpoint handler you can plug into your application
to return stat data
*/
func (s *ServerStats) Handler(ctx echo.Context) error {
	s.mutex.RLock()
	defer s.mutex.RUnlock()

	var averageResponseTime int64
	var numResponses int64
	averageResponseTime = 0
	numResponses = 0

	s.ResponseTimes.Do(func(responseTime interface{}) {
		if responseTimeDuration, ok := responseTime.(time.Duration); ok {
			averageResponseTime += int64(responseTimeDuration)
			numResponses++
		}
	})

	if numResponses > 0 {
		averageResponseTime = averageResponseTime / numResponses
	}

	result := struct {
		AverageResponseTimeInNanoseconds  int64          `json:"averageResponseTimeInNanoseconds"`
		AverageResponseTimeInMicroseconds int64          `json:"averageResponseTimeInMicroseconds"`
		AverageResponseTimeInMilliseconds int64          `json:"averageResponseTimeInMilliseconds"`
		ServerStartTime                   time.Time      `json:"serverStartTime"`
		RequestCount                      uint64         `json:"requestCount"`
		Statuses                          map[string]int `json:"statuses"`
	}{
		AverageResponseTimeInNanoseconds:  averageResponseTime,
		AverageResponseTimeInMicroseconds: averageResponseTime / 1000,
		AverageResponseTimeInMilliseconds: averageResponseTime / 1000 / 1000,
		ServerStartTime:                   s.Uptime,
		RequestCount:                      s.RequestCount,
		Statuses:                          s.Statuses,
	}

	return ctx.JSON(http.StatusOK, result)
}
