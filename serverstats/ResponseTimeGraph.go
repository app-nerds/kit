/*
 * Copyright (c) 2020. App Nerds LLC. All rights reserved
 */

package serverstats

/*
ResponseTimeGraph reports average response times for a given date/time
*/
type ResponseTimeGraph struct {
	AverageResponseTimeInNanoseconds  int64  `json:"averageResponseTimeInNanoseconds"`
	AverageResponseTimeInMicroseconds int64  `json:"averageResponseTimeInMicroseconds"`
	AverageExecutionTimeMilliseconds  int64  `json:"averageExecutionTimeMilliseconds"`
	Time                              string `json:"time"`
}

// ResponseTimeGraphCollection is a collection of ResponseTimeGraph structs
type ResponseTimeGraphCollection []*ResponseTimeGraph
