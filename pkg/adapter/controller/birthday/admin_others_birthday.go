package birthday

import (
	"context"
	"errors"
	"regexp"

	"github.com/bwmarrin/discordgo"
	"github.com/guergeiro/discord-bots/pkg/adapter/controller"
	"github.com/guergeiro/discord-bots/pkg/adapter/presenter"
	"github.com/guergeiro/discord-bots/pkg/application/usecase"
	"github.com/guergeiro/discord-bots/pkg/domain/entity"
)

type BirthdayAdminOthersBirthdayController struct {
	base      *controller.BaseController
	usecase   usecase.UseCase[entity.Birthday]
	presenter presenter.Presenter[entity.Birthday]
}

func NewBirthdayAdminOthersBirthdayController(
	usecase usecase.UseCase[entity.Birthday],
	presenter presenter.Presenter[entity.Birthday],
) *BirthdayAdminOthersBirthdayController {
	return &BirthdayAdminOthersBirthdayController{
		base:      controller.NewBaseController(),
		usecase:   usecase,
		presenter: presenter,
	}
}

var idExtractor = regexp.MustCompile(`\d+`)

func (c *BirthdayAdminOthersBirthdayController) Handle(
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
	if name != "others-birthday" {
		return c.base.Handle(ctx, args...)
	}

	dateValue := i.ApplicationCommandData().Options[0].Options[0].Options[0].Value

	date, err := parseDate(dateValue)
	if err != nil {
		return err
	}
	idValue := i.ApplicationCommandData().Options[0].Options[0].Options[1].Value

	idStr, ok := idValue.(string)
	if !ok {
		return errors.New("There was an error parsing the user id")
	}
	id := idExtractor.FindString(idStr)

	birthday, err := c.usecase.Execute(ctx, id, date)
	if err != nil {
		return err
	}
	return c.presenter.Present(ctx, birthday, s, i)
}

func (c *BirthdayAdminOthersBirthdayController) SetNext(
	next controller.Controller,
) {
	c.base.SetNext(next)
}
