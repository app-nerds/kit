/*
 * Copyright (c) 2020. App Nerds LLC. All rights reserved
 */

package datetime

import (
	"fmt"
	"time"

	"github.com/pkg/errors"
)

/*
IDateTimeParser is an interface describing methods that parse and format dates and times
*/
type IDateTimeParser interface {
	DaysAgo(numDays int) (time.Time, error)
	GetUTCLocation() *time.Location
	NowUTC() time.Time
	IsDateOlderThanNumDaysAgo(t time.Time, numDays int) bool
	Parse(dateString string) (time.Time, error)
	ParseDateTime(dateString string) time.Time
	ParseISO8601(dateString string) time.Time
	ParseISO8601SqlUtc(dateString string) time.Time
	ParseShortDate(dateString string) time.Time
	ParseUSDateTime(dateString string) time.Time
	Pretty(t time.Time) string
	ToISO8601(t time.Time) string
	ToSQLString(t time.Time) string
	ToUSDate(t time.Time) string
	ToUSDateTime(t time.Time) string
	ToUSTime(t time.Time) string
	ValidDateTime(dateString string) bool
	ValidISO8601(dateString string) bool
	ValidShortDate(dateString string) bool
	ValidISO8601SqlUtc(dateString string) bool
	ValidUSDateTime(dateString string) bool
}

/*
DateTimeParser provides methods for parsing and formatting dates
*/
type DateTimeParser struct{}

/*
DaysAgo returns the current date minus the number of specified days
*/
func (service *DateTimeParser) DaysAgo(numDays int) (time.Time, error) {
	var hoursAgo time.Duration
	var err error

	hoursAgoString := fmt.Sprintf("%dh", -24*numDays)
	if hoursAgo, err = time.ParseDuration(hoursAgoString); err != nil {
		return time.Now().UTC(), errors.Wrapf(err, "Unable to convert days to hours string")
	}

	result := time.Now().UTC().Add(hoursAgo)
	return result, nil
}

/*
GetUTCLocation returns location information for the UTC timezone
*/
func (service *DateTimeParser) GetUTCLocation() *time.Location {
	return service.NowUTC().Location()
}

/*
IsDateOlderThanNumDaysAgo returns true if the provided date is older
than today minus the specified number of days. All times are UTC
*/
func (service *DateTimeParser) IsDateOlderThanNumDaysAgo(t time.Time, numDays int) bool {
	hours := (24 * numDays) * -1
	expiration := service.NowUTC().Add(time.Duration(hours) * time.Hour)
	return t.Before(expiration)
}

/*
NowUTC returns the current date/time in UTC
*/
func (service *DateTimeParser) NowUTC() time.Time {
	return time.Now().UTC()
}

/*
Parse is a general purposes method for parsing a date string. This will
attempt to parse the string in the following formats:

  * SQL (with full milliseconds and UTC indicator)
  * SQL (YYYY-MM-DDTHH:MM:SS)
  * Short Date (YYYY-MM-DD)
  * US Date/Time (MM/DD/YYYY H:MM A)
*/
func (service *DateTimeParser) Parse(dateString string) (time.Time, error) {
	if service.ValidISO8601(dateString) {
		return service.ParseISO8601(dateString), nil
	}

	if service.ValidISO8601SqlUtc(dateString) {
		return service.ParseISO8601SqlUtc(dateString), nil
	}

	if service.ValidDateTime(dateString) {
		return service.ParseDateTime(dateString), nil
	}

	if service.ValidShortDate(dateString) {
		return service.ParseShortDate(dateString), nil
	}

	if service.ValidUSDateTime(dateString) {
		return service.ParseUSDateTime(dateString), nil
	}

	return service.NowUTC(), fmt.Errorf("Unknown date/time format: %s", dateString)
}

/*
ParseDateTime parses a string as a basic SQL date YYYY-MM-DDTHH:MM:SS
*/
func (service *DateTimeParser) ParseDateTime(dateString string) time.Time {
	result, _ := time.Parse("2006-01-02T15:04:05", dateString)
	return result
}

/*
ParseISO8601 parses a string as an ISO 8601 format YYYY-MM-DDTHH:MM:SS-07:00.
The string must have the timezone indicated as an offset.
*/
func (service *DateTimeParser) ParseISO8601(dateString string) time.Time {
	result, _ := time.Parse("2006-01-02T15:04:05-07:00", dateString)
	return result
}

/*
ParseISO8601SqlUtc parses a string in SQL format with milliseconds and the
UTC indicator YYYY-MM-DDTHH:MM:SS.000Z
*/
func (service *DateTimeParser) ParseISO8601SqlUtc(dateString string) time.Time {
	result, _ := time.Parse("2006-01-02T15:04:05.999Z", dateString)
	return result
}

/*
ParseShortDate parses a short date YYYY-MM-DD
*/
func (service *DateTimeParser) ParseShortDate(dateString string) time.Time {
	result, _ := time.Parse("2006-01-02", dateString)
	return result
}

/*
ParseUSDateTime parses a standard US date/time MM/DD/YYYY H:MM A
*/
func (service *DateTimeParser) ParseUSDateTime(dateString string) time.Time {
	result, _ := time.Parse("01/02/2006 3:04 PM", dateString)
	return result
}

/*
Pretty returns a date/time formatted Jan 1 2010 at H:MMAM
*/
func (service *DateTimeParser) Pretty(t time.Time) string {
	var result string

	result = t.Format("Jan _2 2006 at 3:04PM")
	return result
}

/*
ToISO8601 formats a time as YYYY-MM-DDTHH:MM:SS-07:00, using an offset
to indicate timezone
*/
func (service *DateTimeParser) ToISO8601(t time.Time) string {
	return t.Format("2006-01-02T15:04:05-0700")
}

/*
ToSQLString formats a time as YYYY-MM-DD HH:MM:SS. This is useful
for inserting into a database
*/
func (service *DateTimeParser) ToSQLString(t time.Time) string {
	return t.Format("2006-01-02 15:04:05")
}

/*
ToUSDate formats a time as MM/DD/YYYY
*/
func (service *DateTimeParser) ToUSDate(t time.Time) string {
	return t.Format("01/02/2006")
}

/*
ToUSDateTime formats a time as MM/DD/YYYY H:MM AM
*/
func (service *DateTimeParser) ToUSDateTime(t time.Time) string {
	return t.Format("01/02/2006 3:04 PM")
}

/*
ToUSTime formats a time as H:MM AM
*/
func (service *DateTimeParser) ToUSTime(t time.Time) string {
	return t.Format("3:04 PM")
}

/*
ValidDateTime returns true if the string is YYYY-MM-DDTHH:MM:SS
*/
func (service *DateTimeParser) ValidDateTime(dateString string) bool {
	if _, err := time.Parse("2006-01-02T15:04:05", dateString); err != nil {
		return false
	}

	return true
}

/*
ValidISO8601 returns true if the string is YYYY-MM-DDTHH:MM:SS-07:00
*/
func (service *DateTimeParser) ValidISO8601(dateString string) bool {
	if _, err := time.Parse("2006-01-02T15:04:05-07:00", dateString); err != nil {
		return false
	}

	return true
}

/*
ValidShortDate returns true if the string is YYYY-MM-DD
*/
func (service *DateTimeParser) ValidShortDate(dateString string) bool {
	if _, err := time.Parse("2006-01-02", dateString); err != nil {
		return false
	}

	return true
}

/*
ValidISO8601SqlUtc returns true if the string is YYYY-MM-DDTHH:MM:SS.000Z
*/
func (service *DateTimeParser) ValidISO8601SqlUtc(dateString string) bool {
	if _, err := time.Parse("2006-01-02T15:04:05.999Z", dateString); err != nil {
		return false
	}

	return true
}

/*
ValidUSDateTime returns true if the string is MM/DD/YYYY H:MM AM
*/
func (service *DateTimeParser) ValidUSDateTime(dateString string) bool {
	if _, err := time.Parse("01/02/2006 3:04 PM", dateString); err != nil {
		return false
	}

	return true
}
