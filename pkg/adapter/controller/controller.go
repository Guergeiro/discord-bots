package controller

import (
	"context"
	"errors"
)

type Controller interface {
	Handle(ctx context.Context, args ...any) error
	SetNext(next Controller)
}

type BaseController struct {
	next Controller
}

func NewBaseController() *BaseController {
	return &BaseController{}
}

func (c *BaseController) Handle(
	ctx context.Context,
	args ...any,
) error {
	if c.next != nil {
		return c.next.Handle(ctx, args...)
	}
	return errors.New("No handlers available")
}

func (c *BaseController) SetNext(next Controller) {
	c.next = next
}
