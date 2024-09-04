package birthday

import (
	"fmt"
	"strings"
	"time"
)

const timelayout = "2006-01-02"

type Birthday struct {
	db map[time.Time]string
}

func (b Birthday) Handle() string {

	message := []string{
		"@everyone",
		"this is a test, nothing to see here",
	}
	return strings.Join(message, "\n")
}

func (b Birthday) Find(day time.Time) (string, error) {
	if v, ok := b.db[day]; ok {
		return v, nil
	}
	return "", fmt.Errorf("No birthday")
}

func (b Birthday) Upsert(day time.Time, id string) error {
	b.db[day] = id
	return nil
}

func (b Birthday) Delete(day time.Time) error {
	delete(b.db, day)
	return nil
}
