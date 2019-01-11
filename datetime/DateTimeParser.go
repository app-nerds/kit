package datetime

import (
	"fmt"
	"time"
)

type IDateTimeParser interface {
	GetUTCLocation() *time.Location
	NowUTC() time.Time
	IsDateOlderThanNumDaysAgo(t time.Time, numDays int) bool
	ToISO8601(t time.Time) string
	ToSQLString(t time.Time) string
	ToUSDate(t time.Time) string
	ToUSDateTime(t time.Time) string
	ToUSTime(t time.Time) string
	Parse(dateString string) (time.Time, error)
	ParseDateTime(dateString string) time.Time
	ParseISO8601(dateString string) time.Time
	ParseShortDate(dateString string) time.Time
	ParseISO8601SqlUtc(dateString string) time.Time
	ParseUSDateTime(dateString string) time.Time
	ValidDateTime(dateString string) bool
	ValidISO8601(dateString string) bool
	ValidShortDate(dateString string) bool
	ValidISO8601SqlUtc(dateString string) bool
	ValidUSDateTime(dateString string) bool
}

type DateTimeParser struct{}

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

func (service *DateTimeParser) ToISO8601(t time.Time) string {
	return t.Format("2006-01-02T15:04:05-0700")
}

func (service *DateTimeParser) ToSQLString(t time.Time) string {
	return t.Format("2006-01-02 15:04:05")
}

func (service *DateTimeParser) ToUSDate(t time.Time) string {
	return t.Format("01/02/2006")
}

func (service *DateTimeParser) ToUSDateTime(t time.Time) string {
	return t.Format("01/02/2006 3:04 PM")
}

func (service *DateTimeParser) ToUSTime(t time.Time) string {
	return t.Format("3:04 PM")
}

func (service *DateTimeParser) NowUTC() time.Time {
	return time.Now().UTC()
}

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

func (service *DateTimeParser) ParseDateTime(dateString string) time.Time {
	result, _ := time.Parse("2006-01-02T15:04:05", dateString)
	return result
}

func (service *DateTimeParser) ParseISO8601(dateString string) time.Time {
	result, _ := time.Parse("2006-01-02T15:04:05-07:00", dateString)
	return result
}

func (service *DateTimeParser) ParseShortDate(dateString string) time.Time {
	result, _ := time.Parse("2006-01-02", dateString)
	return result
}

func (service *DateTimeParser) ParseISO8601SqlUtc(dateString string) time.Time {
	result, _ := time.Parse("2006-01-02T15:04:05.999Z", dateString)
	return result
}

func (service *DateTimeParser) ParseUSDateTime(dateString string) time.Time {
	result, _ := time.Parse("01/02/2006 3:04 PM", dateString)
	return result
}

func (service *DateTimeParser) ValidDateTime(dateString string) bool {
	if _, err := time.Parse("2006-01-02T15:04:05", dateString); err != nil {
		return false
	}

	return true
}

func (service *DateTimeParser) ValidISO8601(dateString string) bool {
	if _, err := time.Parse("2006-01-02T15:04:05-07:00", dateString); err != nil {
		return false
	}

	return true
}

func (service *DateTimeParser) ValidShortDate(dateString string) bool {
	if _, err := time.Parse("2006-01-02", dateString); err != nil {
		return false
	}

	return true
}

func (service *DateTimeParser) ValidISO8601SqlUtc(dateString string) bool {
	if _, err := time.Parse("2006-01-02T15:04:05.999Z", dateString); err != nil {
		return false
	}

	return true
}

func (service *DateTimeParser) ValidUSDateTime(dateString string) bool {
	if _, err := time.Parse("01/02/2006 3:04 PM", dateString); err != nil {
		return false
	}

	return true
}
