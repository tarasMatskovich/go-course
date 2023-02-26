package model

import (
	"library/config"
	"strings"
	"time"
)

type BookTime struct {
	Date time.Time
}

func (t BookTime) MarshalJSON() ([]byte, error) {
	format := config.Configuration.TimeFormat
	date := t.Date.Format(format)
	b := make([]byte, 0, len(date)+2)
	b = append(b, '"')
	b = t.Date.AppendFormat(b, format)
	b = append(b, '"')

	return b, nil
}

func (t *BookTime) UnmarshalJSON(data []byte) error {
	format := config.Configuration.TimeFormat
	date := strings.Trim(string(data), "\"")

	parsedDate, err := time.Parse(format, date)

	if err != nil {
		return err
	}

	t.Date = parsedDate

	return nil
}
