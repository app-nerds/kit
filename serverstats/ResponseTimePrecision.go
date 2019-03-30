package serverstats

// ResponseTimePrecision describes how granular to report response times
type ResponseTimePrecision int

const (
	// PrecisionHour reports responses times averaged by hour
	PrecisionHour ResponseTimePrecision = 1

	// PrecisionDay reports respose times averaged by day
	PrecisionDay ResponseTimePrecision = 2

	// PrecisionMonth reports response times averaged by month
	PrecisionMonth ResponseTimePrecision = 3
)
