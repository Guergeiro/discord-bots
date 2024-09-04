package repository

import (
	"context"
	"time"
)

type FindByDate[O any] interface {
	FindByDate(ctx context.Context, date time.Time) ([]O, error)
}
