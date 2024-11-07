package birthday

import (
	"context"

	"github.com/bwmarrin/discordgo"
	"github.com/guergeiro/discord-bots/pkg/adapter/controller"
	"github.com/guergeiro/discord-bots/pkg/adapter/presenter"
	"github.com/guergeiro/discord-bots/pkg/application/usecase"
	"github.com/guergeiro/discord-bots/pkg/domain/entity"
)

type BirthdaySetController struct {
	base      *controller.BaseController
	usecase   usecase.UseCase[entity.Birthday]
	presenter presenter.Presenter[entity.Birthday]
}

func NewBirthdaySetController(
	usecase usecase.UseCase[entity.Birthday],
	presenter presenter.Presenter[entity.Birthday],
) *BirthdaySetController {
	return &BirthdaySetController{
		base:      controller.NewBaseController(),
		usecase:   usecase,
		presenter: presenter,
	}
}

func (c *BirthdaySetController) Handle(
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
	if name != "set" {
		return c.base.Handle(ctx, args...)
	}

	dateValue := i.ApplicationCommandData().Options[0].Options[0].Value

	date, err := parseDate(dateValue)
	if err != nil {
		return err
	}
	id := i.Member.User.ID

	birthday, err := c.usecase.Execute(ctx, id, date)
	if err != nil {
		return err
	}
	return c.presenter.Present(ctx, birthday, s, i)
}

func (c *BirthdaySetController) SetNext(next controller.Controller) {
	c.base.SetNext(next)
}
