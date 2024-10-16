package controller

import (
	"context"
	"log"
)

type Controller[O any] interface {
	Handle(ctx context.Context, args ...any) O
	SetNext(next Controller[O])
}

type BaseController[O any] struct {
	next Controller[O]
}

func NewBaseController[O any]() *BaseController[O] {
	return &BaseController[O]{}
}

func (c *BaseController[O]) Handle(
	ctx context.Context,
	args ...any,
) O {
	if c.next != nil {
		return c.next.Handle(ctx, args...)
	}
	log.Println("No handlers available")
	var result O
	return result
}

func (c *BaseController[O]) SetNext(next Controller[O]) {
	c.next = next
}
