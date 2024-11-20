package entity

import (
	"github.com/stretchr/testify/assert"
	"testing"
	"time"
)

func TestNewBirthday(t *testing.T) {
	id := "123"
	date := time.Date(2000, time.January, 15, 0, 0, 0, 0, time.UTC)

	birthday := NewBirthday(id, date)

	assert.Equal(t, id, birthday.Id)
	assert.Equal(t, date, birthday.Date)
}

func TestPrettyBirthday(t *testing.T) {
	id := "123"
	date := time.Date(2000, time.January, 15, 0, 0, 0, 0, time.UTC)

	birthday := NewBirthday(id, date)
	expectedOutput := "15 January 2000"

	assert.Equal(t, expectedOutput, birthday.PrettyBirthday())
}
