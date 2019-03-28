package serverstats

import "time"

type ResponseTimeGraph struct {
	AverageExecutionTimeMilliseconds int64     `json:"averageExecutionTimeMilliseconds"`
	Time                             time.Time `json:"time"`
}

type ResponseTimeGraphCollection []*ResponseTimeGraph
