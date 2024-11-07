package presenter

import "context"

type Presenter[I any] interface {
	Present(ctx context.Context, input I, args ...any) error
}
