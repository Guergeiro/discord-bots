package usecase

import "context"

type UseCase[O any] interface {
	Execute(ctx context.Context, args ...any) (O, error)
}
