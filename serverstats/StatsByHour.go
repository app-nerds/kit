package serverstats

import (
	"sync"

	"github.com/dustin/go-humanize"
)

type StatsByHour struct {
	Hour                              int                    `json:"hour"`
	AverageFreeMemory                 uint64                 `json:"averageFreeMemory"`
	AverageFreeMemoryPretty           string                 `json:"averageFreeMemoryPretty"`
	AverageMemoryUsage                uint64                 `json:"averageMemoryUsage"`
	AverageMemoryUsagePretty          string                 `json:"averageMemoryUsagePretty"`
	AverageResponseTimeInNanoseconds  int64                  `json:"averageResponseTimeInNanoseconds"`
	AverageResponseTimeInMicroseconds int64                  `json:"averageResponseTimeInMicroseconds"`
	AverageResponseTimeInMilliseconds int64                  `json:"averageResponseTimeInMilliseconds"`
	CustomStats                       map[string]interface{} `json:"customStats"`
	RequestCount                      uint64                 `json:"requestCount"`
	Statuses                          map[string]int         `json:"statuses"`

	sync.RWMutex `json:"-"`
}

type StatsByHourCollection []*StatsByHour

func NewStatsByHour(hour int) *StatsByHour {
	return &StatsByHour{
		Hour:        hour,
		CustomStats: make(map[string]interface{}),
		Statuses:    make(map[string]int),

		RWMutex: sync.RWMutex{},
	}
}

func (sbh *StatsByHour) Calculate(s *ServerStats) {
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

	sbh.AverageFreeMemory = averageFreeMemory
	sbh.AverageFreeMemoryPretty = humanize.Bytes(averageFreeMemory)
	sbh.AverageMemoryUsage = averageMemoryUsage
	sbh.AverageMemoryUsagePretty = humanize.Bytes(averageMemoryUsage)
	sbh.AverageResponseTimeInNanoseconds = averageResponseTime
	sbh.AverageResponseTimeInMicroseconds = averageResponseTime / 1000
	sbh.AverageResponseTimeInMilliseconds = averageResponseTime / 1000 / 1000
	sbh.CustomStats = s.CustomStats
	sbh.RequestCount = s.RequestCount
	sbh.Statuses = s.Statuses
}
