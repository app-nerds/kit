package serverstats

import (
	"time"
)

/*
ResponseTime is used to track how much time a request took to
execute, and what time (of day) it happened
*/
type ResponseTime struct {
	ExecutionTime time.Duration
	Time          time.Time
}
