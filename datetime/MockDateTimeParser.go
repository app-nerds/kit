package datetime

import "time"

type MockDateTimeParser struct {
	DaysAgoFunc                   func(numDays int) (time.Time, error)
	GetUTCLocationFunc            func() *time.Location
	NowUTCFunc                    func() time.Time
	IsDateOlderThanNumDaysAgoFunc func(t time.Time, numDays int) bool
	ParseFunc                     func(dateString string) (time.Time, error)
	ParseDateTimeFunc             func(dateString string) time.Time
	ParseISO8601Func              func(dateString string) time.Time
	ParseISO8601SqlUtcFunc        func(dateString string) time.Time
	ParseShortDateFunc            func(dateString string) time.Time
	ParseUSDateTimeFunc           func(dateString string) time.Time
	PrettyFunc                    func(t time.Time) string
	ToISO8601Func                 func(t time.Time) string
	ToSQLStringFunc               func(t time.Time) string
	ToUSDateFunc                  func(t time.Time) string
	ToUSDateTimeFunc              func(t time.Time) string
	ToUSTimeFunc                  func(t time.Time) string
	ValidDateTimeFunc             func(dateString string) bool
	ValidISO8601Func              func(dateString string) bool
	ValidShortDateFunc            func(dateString string) bool
	ValidISO8601SqlUtcFunc        func(dateString string) bool
	ValidUSDateTimeFunc           func(dateString string) bool
}

func (m *MockDateTimeParser) DaysAgo(numDays int) (time.Time, error) {
	return m.DaysAgoFunc(numDays)
}

func (m *MockDateTimeParser) GetUTCLocation() *time.Location {
	return m.GetUTCLocationFunc()
}

func (m *MockDateTimeParser) NowUTC() time.Time {
	return m.NowUTCFunc()
}

func (m *MockDateTimeParser) IsDateOlderThanNumDaysAgo(t time.Time, numDays int) bool {
	return m.IsDateOlderThanNumDaysAgoFunc(t, numDays)
}

func (m *MockDateTimeParser) Parse(dateString string) (time.Time, error) {
	return m.ParseFunc(dateString)
}

func (m *MockDateTimeParser) ParseDateTime(dateString string) time.Time {
	return m.ParseDateTimeFunc(dateString)
}

func (m *MockDateTimeParser) ParseISO8601(dateString string) time.Time {
	return m.ParseISO8601Func(dateString)
}

func (m *MockDateTimeParser) ParseISO8601SqlUtc(dateString string) time.Time {
	return m.ParseISO8601SqlUtcFunc(dateString)
}

func (m *MockDateTimeParser) ParseShortDate(dateString string) time.Time {
	return m.ParseShortDateFunc(dateString)
}

func (m *MockDateTimeParser) ParseUSDateTime(dateString string) time.Time {
	return m.ParseUSDateTimeFunc(dateString)
}

func (m *MockDateTimeParser) Pretty(t time.Time) string {
	return m.PrettyFunc(t)
}

func (m *MockDateTimeParser) ToISO8601(t time.Time) string {
	return m.ToISO8601Func(t)
}

func (m *MockDateTimeParser) ToSQLString(t time.Time) string {
	return m.ToSQLStringFunc(t)
}

func (m *MockDateTimeParser) ToUSDate(t time.Time) string {
	return m.ToUSDateFunc(t)
}

func (m *MockDateTimeParser) ToUSDateTime(t time.Time) string {
	return m.ToUSDateTimeFunc(t)
}

func (m *MockDateTimeParser) ToUSTime(t time.Time) string {
	return m.ToUSTimeFunc(t)
}

func (m *MockDateTimeParser) ValidDateTime(dateString string) bool {
	return m.ValidDateTimeFunc(dateString)
}

func (m *MockDateTimeParser) ValidISO8601(dateString string) bool {
	return m.ValidISO8601Func(dateString)
}

func (m *MockDateTimeParser) ValidShortDate(dateString string) bool {
	return m.ValidShortDateFunc(dateString)
}

func (m *MockDateTimeParser) ValidISO8601SqlUtc(dateString string) bool {
	return m.ValidISO8601SqlUtcFunc(dateString)
}

func (m *MockDateTimeParser) ValidUSDateTime(dateString string) bool {
	return m.ValidUSDateTimeFunc(dateString)
}
