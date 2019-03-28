package serverstats

type ResponseTimePrecision int

const (
	PRECISION_HOUR  ResponseTimePrecision = 1
	PRECISION_DAY   ResponseTimePrecision = 2
	PRECISION_MONTH ResponseTimePrecision = 3
)
