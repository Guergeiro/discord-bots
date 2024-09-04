package repository

import "context"

type RemoveOne interface {
	RemoveOne(ctx context.Context, id string) error
}
