package birthday

import (
	"context"
	"fmt"

	"github.com/bwmarrin/discordgo"
	"github.com/guergeiro/discord-bots/internal/env"
	"github.com/guergeiro/discord-bots/pkg/adapter/controller"
)

type BirthdayAdminController struct {
	base     *controller.BaseController
	original controller.Controller
}

func NewBirthdayAdminController(
	original controller.Controller,
) *BirthdayAdminController {
	return &BirthdayAdminController{
		base:     controller.NewBaseController(),
		original: original,
	}
}

func (c *BirthdayAdminController) Handle(
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
	if name != "admin" {
		return c.base.Handle(ctx, args...)
	}

	ADMIN_ID, err := env.Get("ADMIN_ID")
	if err != nil {
		return err
	}

	caller := i.Interaction.Member.User.ID
	if caller != ADMIN_ID {
		return fmt.Errorf("Only <@%s> is able to use this command", ADMIN_ID)
	}

	return c.original.Handle(
		ctx,
		s,
		i,
	)
}

func (c *BirthdayAdminController) SetNext(next controller.Controller) {
	c.base.SetNext(next)
}
