package repository

import "context"

type FindAll[O any] interface {
	FindAll(ctx context.Context) ([]O, error)
}
