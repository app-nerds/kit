/*
 * Copyright (c) 2020. App Nerds LLC. All rights reserved
 */

package serverstats

import (
	"container/ring"
	"fmt"
	"net/http"
	"runtime"
	"strconv"
	"sync"
	"time"

	"github.com/dustin/go-humanize"
	"github.com/labstack/echo/v4"
	"github.com/shirou/gopsutil/mem"
)

type ServerStatsOptions struct {
	NumMemStatsToKeep      int
	NumResponseTimesToKeep int
}

/*
ServerStats tracks general server statistics. This includes information
about uptime, response times and counts, and requests counts broken
down by HTTP status code. ServerStats is thread-safe due to a
write lock on requests, and a read lock on reads
*/
type ServerStats struct {
	AverageFreeSystemMemory *ring.Ring
	AverageMemoryUsage      *ring.Ring
	CustomStats             map[string]interface{} `json:"customStats"`
	Uptime                  time.Time              `json:"uptime"`
	RequestCount            uint64                 `json:"requestCount"`
	ResponseTimes           *ring.Ring
	StatsByDayCollection    StatsByDayCollection
	Statuses                map[string]int `json:"statuses"`
	customMiddleware        func(ctx echo.Context, serverStats *ServerStats)

	sync.RWMutex
}

/*
NewServerStats creates a new ServerStats object
*/
func NewServerStats(customMiddleware func(ctx echo.Context, serverStats *ServerStats)) *ServerStats {
	return &ServerStats{
		AverageFreeSystemMemory: ring.New(100),
		AverageMemoryUsage:      ring.New(100),
		customMiddleware:        customMiddleware,
		CustomStats:             make(map[string]interface{}),
		Uptime:                  time.Now().UTC(),
		ResponseTimes:           ring.New(1000),
		Statuses:                make(map[string]int),

		RWMutex: sync.RWMutex{},
	}
}

func NewServerStatsWithOptions(options ServerStatsOptions, customMiddleware func(ctx echo.Context, serverStats *ServerStats)) *ServerStats {
	return &ServerStats{
		AverageFreeSystemMemory: ring.New(options.NumMemStatsToKeep),
		AverageMemoryUsage:      ring.New(options.NumMemStatsToKeep),
		customMiddleware:        customMiddleware,
		CustomStats:             make(map[string]interface{}),
		Uptime:                  time.Now().UTC(),
		ResponseTimes:           ring.New(options.NumResponseTimesToKeep),
		Statuses:                make(map[string]int),

		RWMutex: sync.RWMutex{},
	}
}

/*
GetAverageResponseTimeGraph returns an array of response time objects. The precision
specifies at what interval you wish to get data for. For example, passing
Hour gets you response times averaged by hour. Passing Day gets you response
times averaged by day
*/
func (s *ServerStats) GetAverageResponseTimeGraph(precision ResponseTimePrecision) ResponseTimeGraphCollection {
	switch precision {
	case PrecisionHour:
		return s.breakResponsesIntoHours()

	case PrecisionMonth:
		return s.breakResponsesIntoMonths()

	default:
		return s.breakResponsesIntoDays()
	}
}

func (s *ServerStats) breakResponsesIntoHours() ResponseTimeGraphCollection {
	result := make(ResponseTimeGraphCollection, 0, 20)
	timeIndex := make(map[string]int64)
	currentTimeToIndex := ""
	count := 0

	s.RLock()
	defer s.RUnlock()

	s.ResponseTimes.Do(func(r interface{}) {
		var ok bool
		var responseTime ResponseTime

		if responseTime, ok = r.(ResponseTime); ok {
			timeToIndex := responseTime.Time.Format("2006-01-02T15:00:00")

			if _, ok = timeIndex[timeToIndex]; !ok {
				if currentTimeToIndex != "" {
					newResponseTimeGraph := &ResponseTimeGraph{
						AverageResponseTimeInNanoseconds:  timeIndex[currentTimeToIndex] / int64(count),
						AverageResponseTimeInMicroseconds: (timeIndex[currentTimeToIndex] / int64(count)) / 1000,
						AverageExecutionTimeMilliseconds:  (timeIndex[currentTimeToIndex] / int64(count)) / 1000 / 1000,
						Time:                              currentTimeToIndex,
					}

					result = append(result, newResponseTimeGraph)
				}

				timeIndex[timeToIndex] = 0
				currentTimeToIndex = timeToIndex
				count = 0
			}

			timeIndex[timeToIndex] += int64(responseTime.ExecutionTime)
			count++
		}
	})

	if currentTimeToIndex != "" {
		newResponseTimeGraph := &ResponseTimeGraph{
			AverageResponseTimeInNanoseconds:  timeIndex[currentTimeToIndex] / int64(count),
			AverageResponseTimeInMicroseconds: (timeIndex[currentTimeToIndex] / int64(count)) / 1000,
			AverageExecutionTimeMilliseconds:  (timeIndex[currentTimeToIndex] / int64(count)) / 1000 / 1000,
			Time:                              currentTimeToIndex,
		}

		result = append(result, newResponseTimeGraph)
	}

	return result
}

func (s *ServerStats) breakResponsesIntoDays() ResponseTimeGraphCollection {
	result := make(ResponseTimeGraphCollection, 0, 20)
	timeIndex := make(map[string]int64)
	currentTimeToIndex := ""
	count := 0

	s.RLock()
	defer s.RUnlock()

	s.ResponseTimes.Do(func(r interface{}) {
		var ok bool
		var responseTime ResponseTime

		if responseTime, ok = r.(ResponseTime); ok {
			timeToIndex := responseTime.Time.Format("2006-01-02T00:00:00")

			if _, ok = timeIndex[timeToIndex]; !ok {
				if currentTimeToIndex != "" {
					newResponseTimeGraph := &ResponseTimeGraph{
						AverageResponseTimeInNanoseconds:  timeIndex[currentTimeToIndex] / int64(count),
						AverageResponseTimeInMicroseconds: (timeIndex[currentTimeToIndex] / int64(count)) / 1000,
						AverageExecutionTimeMilliseconds:  (timeIndex[currentTimeToIndex] / int64(count)) / 1000 / 1000,
						Time:                              currentTimeToIndex,
					}

					result = append(result, newResponseTimeGraph)
				}

				timeIndex[timeToIndex] = 0
				currentTimeToIndex = timeToIndex
				count = 0
			}

			timeIndex[timeToIndex] += int64(responseTime.ExecutionTime)
			count++
		}
	})

	if currentTimeToIndex != "" {
		newResponseTimeGraph := &ResponseTimeGraph{
			AverageResponseTimeInNanoseconds:  timeIndex[currentTimeToIndex] / int64(count),
			AverageResponseTimeInMicroseconds: (timeIndex[currentTimeToIndex] / int64(count)) / 1000,
			AverageExecutionTimeMilliseconds:  (timeIndex[currentTimeToIndex] / int64(count)) / 1000 / 1000,
			Time:                              currentTimeToIndex,
		}

		result = append(result, newResponseTimeGraph)
	}

	return result
}

func (s *ServerStats) breakResponsesIntoMonths() ResponseTimeGraphCollection {
	result := make(ResponseTimeGraphCollection, 0, 20)
	timeIndex := make(map[string]int64)
	currentTimeToIndex := ""
	count := 0

	s.RLock()
	defer s.RUnlock()

	s.ResponseTimes.Do(func(r interface{}) {
		var ok bool
		var responseTime ResponseTime

		if responseTime, ok = r.(ResponseTime); ok {
			timeToIndex := fmt.Sprintf("%d-%02d-01T00:00:00", responseTime.Time.Year(), responseTime.Time.Month())

			if _, ok = timeIndex[timeToIndex]; !ok {
				if currentTimeToIndex != "" {
					newResponseTimeGraph := &ResponseTimeGraph{
						AverageResponseTimeInNanoseconds:  timeIndex[currentTimeToIndex] / int64(count),
						AverageResponseTimeInMicroseconds: (timeIndex[currentTimeToIndex] / int64(count)) / 1000,
						AverageExecutionTimeMilliseconds:  (timeIndex[currentTimeToIndex] / int64(count)) / 1000 / 1000,
						Time:                              currentTimeToIndex,
					}

					result = append(result, newResponseTimeGraph)
				}

				timeIndex[timeToIndex] = 0
				currentTimeToIndex = timeToIndex
				count = 0
			}

			timeIndex[timeToIndex] += int64(responseTime.ExecutionTime)
			count++
		}
	})

	if currentTimeToIndex != "" {
		newResponseTimeGraph := &ResponseTimeGraph{
			AverageResponseTimeInNanoseconds:  timeIndex[currentTimeToIndex] / int64(count),
			AverageResponseTimeInMicroseconds: (timeIndex[currentTimeToIndex] / int64(count)) / 1000,
			AverageExecutionTimeMilliseconds:  (timeIndex[currentTimeToIndex] / int64(count)) / 1000 / 1000,
			Time:                              currentTimeToIndex,
		}

		result = append(result, newResponseTimeGraph)
	}

	return result
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

		s.Lock()
		defer s.Unlock()

		s.RequestCount++

		s.ResponseTimes = s.ResponseTimes.Next()
		s.ResponseTimes.Value = ResponseTime{
			ExecutionTime: endTime,
			Time:          startTime.UTC(),
		}

		s.AverageFreeSystemMemory = s.AverageFreeSystemMemory.Next()
		s.AverageMemoryUsage = s.AverageMemoryUsage.Next()

		memStats := &runtime.MemStats{}
		runtime.ReadMemStats(memStats)

		var vMemStats *mem.VirtualMemoryStat
		vMemStats, _ = mem.VirtualMemory()

		s.AverageFreeSystemMemory.Value = vMemStats.Available
		s.AverageMemoryUsage.Value = memStats.Sys

		status := strconv.Itoa(ctx.Response().Status)
		s.Statuses[status]++

		if s.customMiddleware != nil {
			s.customMiddleware(ctx, s)
		}

		return nil
	}
}

/*
NewMiddlewareWithTimeTracking returns a middleware that tracks stats by day and hour. You
provide it a pointer to a StatsByDayCollection and this will update stats grouped by
day (starting at midnight) and hour.
*/
// TODO: Perhaps keep the statsByDayCollection in ServerStats locally. cause otherwise this ain't working
func (s *ServerStats) NewMiddlewareWithTimeTracking() echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(ctx echo.Context) error {
			var err error

			startTime := time.Now()

			if err = next(ctx); err != nil {
				ctx.Error(err)
			}

			endTime := time.Since(startTime)

			s.Lock()
			defer s.Unlock()

			/*
			 * Get the date starting at midnight, and the hour integer
			 */
			day := time.Date(startTime.Year(), startTime.Month(), startTime.Day(), 0, 0, 0, 0, startTime.Location())
			hour := startTime.Hour()

			s.RequestCount++

			s.ResponseTimes = s.ResponseTimes.Next()
			s.ResponseTimes.Value = ResponseTime{
				ExecutionTime: endTime,
				Time:          startTime.UTC(),
			}

			s.AverageFreeSystemMemory = s.AverageFreeSystemMemory.Next()
			s.AverageMemoryUsage = s.AverageMemoryUsage.Next()

			memStats := &runtime.MemStats{}
			runtime.ReadMemStats(memStats)

			var vMemStats *mem.VirtualMemoryStat
			vMemStats, _ = mem.VirtualMemory()

			s.AverageFreeSystemMemory.Value = vMemStats.Available
			s.AverageMemoryUsage.Value = memStats.Sys

			status := strconv.Itoa(ctx.Response().Status)
			s.Statuses[status]++

			if s.customMiddleware != nil {
				s.customMiddleware(ctx, s)
			}

			/*
			 * Find the day and hour, then update the hour structure
			 */
			var byDay *StatsByDay
			var byHour *StatsByHour

			for _, d := range s.StatsByDayCollection {
				if d.Date.Equal(day) {
					byDay = d
					break
				}
			}

			if byDay != nil {
				for _, h := range byDay.HourlyStats {
					if h.Hour == hour {
						byHour = h
						break
					}
				}
			} else {
				byDay = NewStatsByDay(day)
				byHour = NewStatsByHour(hour)

				byDay.HourlyStats = append(byDay.HourlyStats, byHour)
				s.StatsByDayCollection = append(s.StatsByDayCollection, byDay)
			}

			if byHour != nil {
				byHour.Calculate(s)
			}

			return nil
		}
	}
}

/*
Handler is an endpoint handler you can plug into your application
to return stat data
*/
func (s *ServerStats) Handler(ctx echo.Context) error {
	s.RLock()
	defer s.RUnlock()

	var averageResponseTime int64
	var numResponses int64
	var averageFreeMemory uint64
	var averageMemoryUsage uint64

	averageResponseTime = 0
	numResponses = 0

	s.ResponseTimes.Do(func(responseTime interface{}) {
		if responseTimeDuration, ok := responseTime.(ResponseTime); ok {
			averageResponseTime += int64(responseTimeDuration.ExecutionTime)
			numResponses++
		}
	})

	if numResponses > 0 {
		averageResponseTime = averageResponseTime / numResponses
	}

	averageFreeMemory = 0
	numResponses = 0

	s.AverageFreeSystemMemory.Do(func(iFreeMemory interface{}) {
		if freeMemory, ok := iFreeMemory.(uint64); ok {
			averageFreeMemory += freeMemory
			numResponses++
		}
	})

	if numResponses > 0 {
		averageFreeMemory = averageFreeMemory / uint64(numResponses)
	}

	averageMemoryUsage = 0
	numResponses = 0

	s.AverageMemoryUsage.Do(func(iMemUse interface{}) {
		if memUse, ok := iMemUse.(uint64); ok {
			averageMemoryUsage += memUse
			numResponses++
		}
	})

	if numResponses > 0 {
		averageMemoryUsage = averageMemoryUsage / uint64(numResponses)
	}

	result := struct {
		AverageFreeMemory                 uint64                 `json:"averageFreeMemory"`
		AverageFreeMemoryPretty           string                 `json:"averageFreeMemoryPretty"`
		AverageMemoryUsage                uint64                 `json:"averageMemoryUsage"`
		AverageMemoryUsagePretty          string                 `json:"averageMemoryUsagePretty"`
		AverageResponseTimeInNanoseconds  int64                  `json:"averageResponseTimeInNanoseconds"`
		AverageResponseTimeInMicroseconds int64                  `json:"averageResponseTimeInMicroseconds"`
		AverageResponseTimeInMilliseconds int64                  `json:"averageResponseTimeInMilliseconds"`
		CustomStats                       map[string]interface{} `json:"customStats"`
		ServerStartTime                   time.Time              `json:"serverStartTime"`
		RequestCount                      uint64                 `json:"requestCount"`
		Statuses                          map[string]int         `json:"statuses"`
	}{
		AverageFreeMemory:                 averageFreeMemory,
		AverageFreeMemoryPretty:           humanize.Bytes(averageFreeMemory),
		AverageMemoryUsage:                averageMemoryUsage,
		AverageMemoryUsagePretty:          humanize.Bytes(averageMemoryUsage),
		AverageResponseTimeInNanoseconds:  averageResponseTime,
		AverageResponseTimeInMicroseconds: averageResponseTime / 1000,
		AverageResponseTimeInMilliseconds: averageResponseTime / 1000 / 1000,
		CustomStats:                       s.CustomStats,
		ServerStartTime:                   s.Uptime,
		RequestCount:                      s.RequestCount,
		Statuses:                          s.Statuses,
	}

	return ctx.JSON(http.StatusOK, result)
}
