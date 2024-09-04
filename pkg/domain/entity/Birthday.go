package entity

import "time"

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
