package sqldatabase

import (
	"context"
	"database/sql"
	"fmt"
	"strconv"
	"time"
)

/*
GetDBContext returns a context with the timeout set to the value
configured in the application config. Timeouts are in seconds
*/
func GetDBContext(timeout int) (context.Context, context.CancelFunc) {
	return context.WithTimeout(context.Background(), time.Second*time.Duration(timeout))
}

/*
LimitAndOffset returns a string used in a SQL query to limit the number
of records and set the initial offet
*/
func LimitAndOffset(page, pageSize int) string {
	offset := page * pageSize
	return fmt.Sprintf(" LIMIT %d OFFSET %d ", pageSize, offset)
}

/*
NullBool returns the bool value from a SQL bool. If the SQL value
is null false is returned.
*/
func NullBool(value sql.NullBool) bool {
	if !value.Valid {
		return false
	}

	return value.Bool
}

/*
NullDate parses a string date in the format of 2006-01-02 to time.Time. If that
format doesn't work then 2006-01-02T15:04:05Z is used.
*/
func NullDate(value sql.NullString) time.Time {
	parsedDate := NullDateWithFormat(value, "2006-01-02")

	if parsedDate.IsZero() {
		parsedDate = NullDateWithFormat(value, "2006-01-02T15:04:05Z")
	}

	return parsedDate
}

/*
NullDateWithFormat parses a string date in the format provided. If the provided
SQL value is null then an empty time struct is returned
*/
func NullDateWithFormat(value sql.NullString, format string) time.Time {
	var (
		result      time.Time
		stringValue string
	)

	if !value.Valid {
		return result
	}

	stringValue = value.String

	result, _ = time.Parse(format, stringValue)
	return result
}

/*
NullFloat returns the float value from a SQL float. If the SQL value
is null 0.0 is returned
*/
func NullFloat(value sql.NullFloat64) float64 {
	if !value.Valid {
		return 0.0
	}

	return value.Float64
}

/*
NullFloatFromString returns the float value from a SQL string. If the
the SQL string is null 0.0 is returned.
*/
func NullFloatFromString(value sql.NullString) float64 {
	var (
		stringValue string
		result      float64
	)

	if !value.Valid {
		return result
	}

	stringValue = value.String

	result, _ = strconv.ParseFloat(stringValue, 64)
	return result
}

/*
NullInt returns the int value from a SQL int64. If the SQL value is
null then 0 is returned
*/
func NullInt(value sql.NullInt64) int {
	var (
		int64Value int64
		result     int
	)

	if !value.Valid {
		return result
	}

	int64Value = value.Int64
	return int(int64Value)
}

/*
NullString returns the string value from a SQL string. Empty is returned
if the SQL value is null
*/
func NullString(value sql.NullString) string {
	if value.Valid {
		return value.String
	}

	return ""
}

/*
NullTime returns the time.Time value of a SQL time. If the SQL time
is null, an empty time value is returned
*/
func NullTime(value sql.NullTime) time.Time {
	if value.Valid {
		return value.Time
	}

	return time.Time{}
}
