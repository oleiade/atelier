package atelier

import (
	"errors"
	"fmt"
	"strconv"
	"strings"
	"time"
)

// BeginningOfTime is the time the Unix clock was started.
var BeginningOfTime = time.Date(1970, 1, 1, 0, 0, 0, 0, time.UTC)

const (
	// DaysInAWeek is the number of days in a week.
	DaysInAWeek = 7

	// DaysInAMonth is the number of days in a month.
	DaysInAMonth = 30

	// DaysInAYear is the number of days in a year.
	DaysInAYear = 365

	// DaysIn5Years is the number of days in 5 years.
	DaysIn5Years = 5 * DaysInAYear
)

// ISO8601Time is a time.Time that can be unmarshalled from an ISO8601-formatted string.
type ISO8601Time struct {
	time.Time
}

// UnmarshalJSON unmarshal an ISO8601-formatted string into an ISO8601Time.
func (ct *ISO8601Time) UnmarshalJSON(b []byte) error {
	str := strings.Trim(string(b), `"`)

	parsedTime, err := time.Parse("2006-01-02T15:04:05-0700", str)
	if err != nil {
		return fmt.Errorf("error parsing time: %w", err)
	}

	ct.Time = parsedTime

	return nil
}

// OneDayBefore returns the time one day before the given time.
func OneDayBefore(t time.Time) time.Time {
	return t.AddDate(0, 0, -1)
}

// OneWeekBefore returns the time one week before the given time.
func OneWeekBefore(t time.Time) time.Time {
	return t.AddDate(0, 0, -7)
}

// OneMonthBefore returns the time one month before the given time.
func OneMonthBefore(t time.Time) time.Time {
	return t.AddDate(0, -1, 0)
}

// OneYearBefore returns the time one year before the given time.
func OneYearBefore(t time.Time) time.Time {
	return t.AddDate(-1, 0, 0)
}

// FiveYearsBefore returns the time 5 years before the given time.
func FiveYearsBefore(t time.Time) time.Time {
	return t.AddDate(-5, 0, 0)
}

// ParseTimePeriod parses a string representation of a period and returns the corresponding time.Duration.
//
// The period string should be in the format of a number followed by a unit, such as "1d" for 1 day or "2w" for 2 weeks.
// Valid units are "d" (day), "w" (week), "m" (month), and "y" (year).
//
// The function returns an error if the period string is empty, contains an invalid number or unit, or has an
// invalid combination of number and unit.
func ParseTimePeriod(s string) (time.Duration, error) {
	if s == "" {
		return 0, errors.New("period string is empty")
	}

	// Normalize the string to lowercase.
	s = strings.ToLower(s)

	// Handle the "max" case, which is special because it doesn't have a number.
	if s == "max" {
		return time.Since(BeginningOfTime), nil
	}

	// Split the string into number and unit
	var numberStr, unit string
	for i, char := range s {
		if char < '0' || char > '9' {
			numberStr = s[:i]
			unit = strings.Trim(s[i:], " \t\n\r")
			break
		}
	}

	number, err := strconv.Atoi(numberStr)
	if err != nil {
		return 0, fmt.Errorf("invalid number in period string: %w", err)
	}

	var (
		isShortForm = func(unit string) bool { return len(unit) == 1 }
		isSingular  = func(unit string) bool { return isShortForm(unit) || !strings.HasSuffix(unit, "s") }
		isPlural    = func(unit string) bool { return isShortForm(unit) || strings.HasSuffix(unit, "s") }
	)

	// Validate the combination of number and unit
	if (number == 0 || number == 1) && !isSingular(unit) {
		return 0, fmt.Errorf("invalid unit for singular form")
	}
	if (number > 1) && !isPlural(unit) {
		return 0, fmt.Errorf("invalid unit for plural form")
	}

	var duration time.Duration
	switch unit {
	case "s", "sec", "secs", "second", "seconds":
		duration = time.Duration(number) * time.Second
	case "min", "mins", "minute", "minutes":
		duration = time.Duration(number) * time.Minute
	case "h", "hr", "hour", "hours":
		duration = time.Duration(number) * time.Hour
	case "d", "day", "days":
		duration = time.Duration(number) * 24 * time.Hour
	case "w", "week", "weeks":
		duration = time.Duration(number) * 7 * 24 * time.Hour
	case "m", "month", "months":
		duration = time.Duration(number) * 30 * 24 * time.Hour
	case "y", "year", "years":
		duration = time.Duration(number) * 365 * 24 * time.Hour
	default:
		return 0, fmt.Errorf("invalid unit in period string: %s", unit)
	}

	return duration, nil
}
