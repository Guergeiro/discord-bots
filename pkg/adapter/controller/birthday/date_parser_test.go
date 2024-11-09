package birthday

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestNotStringParseDate(t *testing.T) {
	_, err := parseDate(123)

	assert.Error(t, err)
}

func TestNotDateParseDate(t *testing.T) {
	_, err := parseDate("a;lskdjf")

	assert.Error(t, err)
}

func TestSuccessParseDate(t *testing.T) {
	strTime := "2000-01-02"
	expected, err := time.Parse(timelayout, strTime)
	assert.Nil(t, err)

	out, err := parseDate(strTime)

	assert.Nil(t, err)
	assert.Equal(t, expected, out)
}
