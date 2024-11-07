package birthday

import (
	"context"
	"github.com/bwmarrin/discordgo"
	"github.com/guergeiro/discord-bots/pkg/adapter/controller"
	"github.com/guergeiro/discord-bots/pkg/adapter/presenter"
	"github.com/guergeiro/discord-bots/pkg/application/usecase"
	"github.com/guergeiro/discord-bots/pkg/domain/entity"
)

type BirthdayAnnouncerController struct {
	base      *controller.BaseController
	usecase   usecase.UseCase[[]entity.Birthday]
	presenter presenter.Presenter[[]entity.Birthday]
}

func NewBirthdayAnnouncerController(
	usecase usecase.UseCase[[]entity.Birthday],
	presenter presenter.Presenter[[]entity.Birthday],
) *BirthdayAnnouncerController {
	return &BirthdayAnnouncerController{
		base:      controller.NewBaseController(),
		usecase:   usecase,
		presenter: presenter,
	}
}

func (c *BirthdayAnnouncerController) Handle(
	ctx context.Context,
	args ...any,
) error {
	if len(args) != 1 {
		return c.base.Handle(ctx, args...)
	}

	s, ok := args[0].(*discordgo.Session)
	if !ok {
		return c.base.Handle(ctx, args...)
	}

	birthdays, err := c.usecase.Execute(ctx)
	if err != nil {
		return err
	}
	return c.presenter.Present(ctx, birthdays, s)
}

func (c *BirthdayAnnouncerController) SetNext(
	next controller.Controller,
) {
	c.base.SetNext(next)
}
