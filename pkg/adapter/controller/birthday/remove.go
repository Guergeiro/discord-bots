package birthday

import (
	"context"

	"github.com/bwmarrin/discordgo"
	"github.com/guergeiro/discord-bots/pkg/adapter/controller"
	"github.com/guergeiro/discord-bots/pkg/adapter/presenter"
	"github.com/guergeiro/discord-bots/pkg/application/usecase"
)

type BirthdayRemoveController struct {
	base      *controller.BaseController
	usecase   usecase.UseCase[string]
	presenter presenter.Presenter[string]
}

func NewBirthdayRemoveController(
	usecase usecase.UseCase[string],
	presenter presenter.Presenter[string],
) *BirthdayRemoveController {
	return &BirthdayRemoveController{
		base:      controller.NewBaseController(),
		usecase:   usecase,
		presenter: presenter,
	}
}

func (c *BirthdayRemoveController) Handle(
	ctx context.Context,
	args ...any,
) error {
	if len(args) != 2 {
		return c.base.Handle(ctx, args...)
	}

	s, ok := args[0].(*discordgo.Session)
	if !ok {
		return c.base.Handle(ctx, args...)
	}

	i, ok := args[1].(*discordgo.InteractionCreate)
	if !ok {
		return c.base.Handle(ctx, args...)
	}
	name := parseName(i.ApplicationCommandData().Options)
	if name != "remove" {
		return c.base.Handle(ctx, args...)
	}

	id := i.Member.User.ID

	_, err := c.usecase.Execute(ctx, id)
	if err != nil {
		return err
	}
	return c.presenter.Present(ctx, "Birthday removed", s, i)
}

func (c *BirthdayRemoveController) SetNext(
	next controller.Controller,
) {
	c.base.SetNext(next)
}
