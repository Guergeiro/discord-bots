package entity

import (
	"fmt"
	"time"
)

type Birthday struct {
	Id   string
	Date time.Time
}

func NewBirthday(
	id string,
	date time.Time,
) Birthday {
	return Birthday{
		Id:   id,
		Date: date,
	}
}

func (b Birthday) PrettyBirthday() string {
	d := b.Date.Day()
	m := b.Date.Month().String()
	y := b.Date.Year()
	return fmt.Sprintf("%d %s %d", d, m, y)
}
