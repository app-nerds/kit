package datetime_test

import (
	"testing"
	"time"

	"github.com/app-nerds/kit/v4/datetime"
)

func TestDaysAgo(t *testing.T) {
	var err error
	var actual time.Time
	service := getService()

	expected := time.Now().UTC().Add(-24 * time.Hour)

	if actual, err = service.DaysAgo(1); err != nil {
		t.Errorf(err.Error())
	}

	if actual.Format("2006-01-02T15:04:05-0700") != expected.Format("2006-01-02T15:04:05-0700") {
		t.Errorf("Expected %v but got %v", expected, actual)
	}
}

func TestToISO8601(t *testing.T) {
	service := getService()
	now := time.Now()
	expected := now.Format("2006-01-02T15:04:05-0700")
	actual := service.ToISO8601(now)

	if actual != expected {
		t.Errorf("Expected '%s' but got '%s'", expected, actual)
	}
}

func TestNowUTC(t *testing.T) {
	service := getService()
	expected := time.Now().UTC().Format("2006-01-02T15:04:05-0700")
	actual := service.NowUTC().Format("2006-01-02T15:04:05-0700")

	if actual != expected {
		t.Errorf("Expected '%s' but got '%s'", expected, actual)
	}
}

func TestParseISO8601(t *testing.T) {
	service := getService()
	expected := "2018-03-26T13:00:00-07:00"

	actual := service.ParseISO8601(expected)

	if expected != actual.Format("2006-01-02T15:04:05-07:00") {
		t.Errorf("Expected '%s' but got '%s'", expected, actual.Format("2006-01-02T15:04:05-07:00"))
	}
}

func TestValidISO8601(t *testing.T) {
	service := getService()
	input := "2018-03-26T13:00:00-07:00"

	expected := true
	actual := service.ValidISO8601(input)

	if actual != expected {
		t.Errorf("Expected date to be valid")
	}
}

func TestValidISO8601FailsWithBadDate(t *testing.T) {
	service := getService()
	input := "2018-03-26 13:00:00-07:00"

	expected := false
	actual := service.ValidISO8601(input)

	if actual != expected {
		t.Errorf("Expected date to be invalid")
	}
}

/******************************************************************************
 * Private methods
 *****************************************************************************/
func getService() *datetime.DateTimeParser {
	return &datetime.DateTimeParser{}
}
