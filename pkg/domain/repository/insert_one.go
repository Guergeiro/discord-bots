package repository

import "context"

type InsertOne[I any] interface {
	InsertOne(ctx context.Context, input I) error
}
