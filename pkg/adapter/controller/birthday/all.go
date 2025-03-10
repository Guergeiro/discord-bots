package birthday

import (
	"context"

	"github.com/bwmarrin/discordgo"
	"github.com/guergeiro/discord-bots/pkg/adapter/controller"
	"github.com/guergeiro/discord-bots/pkg/adapter/presenter"
	"github.com/guergeiro/discord-bots/pkg/application/usecase"
	"github.com/guergeiro/discord-bots/pkg/domain/entity"
)

type BirthdayAllController struct {
	base      *controller.BaseController
	usecase   usecase.UseCase[[]entity.Birthday]
	presenter presenter.Presenter[[]entity.Birthday]
}

func NewBirthdayAllController(
	usecase usecase.UseCase[[]entity.Birthday],
	presenter presenter.Presenter[[]entity.Birthday],
) *BirthdayAllController {
	return &BirthdayAllController{
		base:      controller.NewBaseController(),
		usecase:   usecase,
		presenter: presenter,
	}
}

func (c *BirthdayAllController) Handle(
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
	if name != "all" {
		return c.base.Handle(ctx, args...)
	}

	birthdays, err := c.usecase.Execute(ctx)
	if err != nil {
		return err
	}
	return c.presenter.Present(ctx, birthdays, s, i)
}

func (c *BirthdayAllController) SetNext(next controller.Controller) {
	c.base.SetNext(next)
}
