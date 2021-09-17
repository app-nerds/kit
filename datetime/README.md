# DateTimerParser

DateTimerParser is a service that makes certain date functions a little easier.

## DaysAgo

DaysAgo returns the current date minus a nmber of specified days.

**Method:**

`DaysAgo(numDays int) (time.Time, error)`

**Example:**

```go
serivce := datetime.DateTimerParser{}
threeDaysAgo, err := service.DaysAgo(3)
```

## GetUTCLocation

GetUTCLocation returns locationn information for the UTC timezone.

**Method:**

`GetUTCLocation() *time.Location`

**Example:**

```go
serivce := datetime.DateTimerParser{}
loc := service.GetUTCLocation()
```

## NowUTC

NowUTC returns the current date/time in UTC.

**Method:**

`NowUTC() time.Time`

**Example:**

```go
serivce := datetime.DateTimerParser{}
now := service.NowUTC()
```

## IsDateOlderThanNumDaysAgo

IsDateOlderThanNumDaysAgo returns true if the provided date is older than today inus the specified number of days. All times are converted to UTC.

**Method:**

`IsDateOlderThanNumDaysAgo(t time.Time, numDays int) bool`

**Example:**

```go
serivce := datetime.DateTimerParser{}
someDate, _ := time.DaysAgo(3)
isOlder := service.IsDateOlderThanNumDaysAgo(someDate, 4)

// isOlder == false
```

## Parse

Parse is a general purpose method for parsing a date string. This will attempt to parse the string the following formats:

* YYYY-MM-DDTHH:mm:ss-07:00
* YYYY-MM-DDTHH:mm:ss.000Z
* YYYY-MM-DDTHH:mm:ss
* YYYY-MM-DD
* MM/DD/YYYY H:mm A
* MM/DD/YYYY

**Method:**

`Parse(dateString string) (time.Time, error)`

**Example:**

```go
serivce := datetime.DateTimerParser{}

dateString1 := "2021-01-01T13:00:01-05:00"
dateString2 := "2021-02-01T14:01:02.999Z"
dateString3 := "2021-03-01T15:02:03"
dateString4 := "2021-04-05"
dateString5 := "05/01/2021 5:45 AM"
dateString6 := "06/02/2021"

date1, err := service.Parse(dateString1)
date2, err := service.Parse(dateString2)
date3, err := service.Parse(dateString3)
date4, err := service.Parse(dateString4)
date5, err := service.Parse(dateString5)
date6, err := service.Parse(dateString6)
```

## ParseDateTime

ParseDateTime parses a string as a basic SQL-style date (YYYY-MM-DDTHH:mm:ss).

**Method:**

`ParseDateTime(dateString string) time.Time`

**Example:**

```go
serivce := datetime.DateTimerParser{}

dateString := "2021-03-01T15:02:03"
date := service.ParseDateTime(dateString)
```

## ParseISO8601

ParseISO8601 parses a string as an ISO 8601 format YYYY-MM-DDTHH:MM:SS-07:00.
The string must have the timezone indicated as an offset.

**Method:**

`ParseISO8601(dateString string) time.Time`

**Example:**

```go
serivce := datetime.DateTimerParser{}

dateString := "2021-03-01T15:02:03-06:00"
date := service.ParseISO8601(dateString)
```

## ParseISO8601SqlUtc

ParseISO8601SqlUtc parses a string in SQL format with milliseconds and the
UTC indicator YYYY-MM-DDTHH:MM:SS.000Z.

**Method:**

`ParseISO8601SqlUtc(dateString string) time.Time`

```go
serivce := datetime.DateTimerParser{}

dateString := "2021-03-01T15:02:03.4650600"
date := service.ParseISO8601SqlUtc(dateString)
```

## ParseShortDate

ParseShortDate parses a short date YYYY-MM-DD.

**Method:**

`ParseShortDate(dateString string) time.Time`

**Example:**

```go
serivce := datetime.DateTimerParser{}

dateString := "2021-03-01"
date := service.ParseShortDate(dateString)
```

## ParseUSDateTime 

ParseUSDateTime parses a standard US date/time MM/DD/YYYY H:MM A.

**Method:**

`ParseUSDateTime(dateString string) time.Time`

**Example:**

```go
serivce := datetime.DateTimerParser{}

dateString := "01/02/2021 5:30 PM"
date := service.ParseUSDateTime(dateString)
```

## ParseUSDate

ParseUSDate parses a standard US date MM/DD/YYYY.

**Method:**

`ParseUSDate(dateString string) time.Time`

**Example:**

```go
serivce := datetime.DateTimerParser{}

dateString := "01/02/2021"
date := service.ParseUSDate(dateString)
```

## Pretty

Pretty returns a date/time formatted Jan 1 2010 at H:MMAM.

**Method:**

`Pretty(t time.Time) string`

**Example:**

```go
serivce := datetime.DateTimerParser{}

dateString := "01/02/2021 5:30 PM"
date := service.ParseUSDateTime(dateString)

pretty := service.Pretty(date)
// pretty == Jan 2 2021 at 5:30PM
```

## ToISO8601

ToISO8601 formats a time as YYYY-MM-DDTHH:MM:SS-07:00, using an offset
to indicate timezone.

**Method:**

`ToISO8601(t time.Time) string`

**Example:**

```go
serivce := datetime.DateTimerParser{}

date := time.Now()
formatted := service.ToISO8601(date)

// formatted == 2006-01-02T13:14:15-05:00
```

## ToSQLString

ToSQLString formats a time as YYYY-MM-DD HH:MM:SS. This is useful
for inserting into a database.

**Method:**

`ToSQLString(t time.Time) string`

**Example:**

```go
serivce := datetime.DateTimerParser{}

date := time.Now()
formatted := service.ToSQLString(date)

// formatted == 2006-01-02 13:14:15
```

## ToUSDate

ToUSDate formats a time as MM/DD/YYYY.

**Method:**

`ToUSDate(t time.Time) string`

**Example:**

```go
serivce := datetime.DateTimerParser{}

date := time.Now()
formatted := service.ToUSDate(date)

// formatted == 01/02/2006
```

## ToUSDateTime

ToUSDateTime formats a time as MM/DD/YYYY H:MM AM.

**Method:**

`ToUSDateTime(t time.Time) string`

**Example:**

```go
serivce := datetime.DateTimerParser{}

date := time.Now()
formatted := service.ToUSDateTime(date)

// formatted == 01/02/2006 3:45 PM
```

## ToUSTime

ToUSTime formats a time as H:MM AM.

**Method:**

`ToUSTime(t time.Time) string`

**Example:**

```go
serivce := datetime.DateTimerParser{}

date := time.Now()
formatted := service.ToUSTime(date)

// formatted == 3:45 PM
```

## ValidDateTime

ValidDateTime returns true if the string is YYYY-MM-DDTHH:MM:SS.

**Method:**

`ValidDateTime(dateString string) bool`

**Example:**

```go
serivce := datetime.DateTimerParser{}
isValid1 := service.ValidDateTime("2021-01-01T15:33:02")
isValid2 := service.ValidDateTime("01/02/2021")

// isValid1 == true
// isValid2 == false
```

## ValidISO8601 

ValidISO8601 returns true if the string is YYYY-MM-DDTHH:MM:SS-07:00.

**Method:**

`ValidISO8601(dateString string) bool`

**Example:**

```go
serivce := datetime.DateTimerParser{}
isValid1 := service.ValidISO8601("2021-01-01T15:33:02-03:00")
isValid2 := service.ValidISO8601("01/02/2021")

// isValid1 == true
// isValid2 == false
```

## ValidShortDate

ValidShortDate returns true if the string is YYYY-MM-DD.

**Method:**

`ValidShortDate(dateString string) bool`

**Example:**

```go
serivce := datetime.DateTimerParser{}
isValid1 := service.ValidShortDate("2021-01-01")
isValid2 := service.ValidShortDate("01/02/2021")

// isValid1 == true
// isValid2 == false
```

## ValidISO8601SqlUtc

ValidISO8601SqlUtc returns true if the string is YYYY-MM-DDTHH:MM:SS.000Z.

**Method:**

`ValidISO8601SqlUtc(dateString string) bool`

**Example:**

```go
serivce := datetime.DateTimerParser{}
isValid1 := service.ValidISO8601SqlUtc("2021-01-01T10:00:01.234Z")
isValid2 := service.ValidISO8601SqlUtc("01/02/2021")

// isValid1 == true
// isValid2 == false
```

## ValidUSDateTime

ValidUSDateTime returns true if the string is MM/DD/YYYY H:MM AM.

**Method:**

`ValidUSDateTime(dateString string) bool`

**Example:**

```go
serivce := datetime.DateTimerParser{}
isValid1 := service.ValidUSDateTime("02/01/2021 6:30 AM")
isValid2 := service.ValidUSDateTime("01/02/2021")

// isValid1 == true
// isValid2 == false
```

## ValidUSDate

ValidUSDate returns true if the string is MM/DD/YYYY.

**Method:**

`ValidUSDate(dateString string) bool`

**Example:**

```go
serivce := datetime.DateTimerParser{}
isValid1 := service.ValidUSDate("02/01/2021")
isValid2 := service.ValidUSDate("2021-01-02")

// isValid1 == true
// isValid2 == false
```

