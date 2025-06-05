package birthday

import (
	"context"
	"errors"

	"github.com/bwmarrin/discordgo"
	"github.com/guergeiro/discord-bots/pkg/adapter/controller"
	"github.com/guergeiro/discord-bots/pkg/adapter/presenter"
	"github.com/guergeiro/discord-bots/pkg/application/usecase"
)

type BirthdayAdminRemoveOthersBirthdayController struct {
	base      *controller.BaseController
	usecase   usecase.UseCase[string]
	presenter presenter.Presenter[string]
}

func NewBirthdayAdminRemoveOthersBirthdayController(
	usecase usecase.UseCase[string],
	presenter presenter.Presenter[string],
) *BirthdayAdminRemoveOthersBirthdayController {
	return &BirthdayAdminRemoveOthersBirthdayController{
		base:      controller.NewBaseController(),
		usecase:   usecase,
		presenter: presenter,
	}
}

func (c *BirthdayAdminRemoveOthersBirthdayController) Handle(
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

	name := parseName(i.ApplicationCommandData().Options[0].Options)
	if name != "remove-others-birthday" {
		return c.base.Handle(ctx, args...)
	}

	idValue := i.ApplicationCommandData().Options[0].Options[0].Options[0].Value

	idStr, ok := idValue.(string)
	if !ok {
		return errors.New("There was an error parsing the user id")
	}
	id := idExtractor.FindString(idStr)

	_, err := c.usecase.Execute(ctx, id)
	if err != nil {
		return err
	}
	return c.presenter.Present(ctx, "Birthday removed", s, i)
}

func (c *BirthdayAdminRemoveOthersBirthdayController) SetNext(
	next controller.Controller,
) {
	c.base.SetNext(next)
}
