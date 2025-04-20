package utils

import (
	"fmt"
	"strings"
	"time"
)

type JSONTime time.Time

const (
	DateOnly     = "2006-01-02"
	DateTimeISO  = "2006-01-02T15:04:05Z"
	DateTimeISOZ = time.RFC3339
)

func (t *JSONTime) UnmarshalJSON(data []byte) error {

	s := strings.Trim(string(data), "\"")
	if s == "null" || s == "" {
		*t = JSONTime(time.Time{})

		return nil
	}

	formats := []string{DateOnly, DateTimeISO, DateTimeISOZ}
	var parsedTime time.Time
	var err error

	for _, format := range formats {
		parsedTime, err = time.Parse(format, s)
		if err == nil {
			*t = JSONTime(parsedTime)

			return nil
		}
	}

	return fmt.Errorf("cannot parse time: %v", s)
}

func (t JSONTime) MarshalJSON() ([]byte, error) {
	stamp := fmt.Sprintf("\"%s\"", time.Time(t).Format(DateTimeISO))

	return []byte(stamp), nil
}

func (t JSONTime) Time() time.Time {
	return time.Time(t)
}

func NewJSONTime(t time.Time) JSONTime {
	return JSONTime(t)
}
