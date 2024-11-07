package birthday

import (
	"errors"
	"time"
)

const timelayout = "2006-01-02"

func parseDate(dateValue interface{}) (time.Time, error) {
	dateStr, ok := dateValue.(string)
	if !ok {
		return time.Time{}, errors.New("There was an error parsing the date")
	}
	date, err := time.Parse(timelayout, dateStr)
	if err != nil {
		return time.Time{}, err
	}
	return date, nil
}
