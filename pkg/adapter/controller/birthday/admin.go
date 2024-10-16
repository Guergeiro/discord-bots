package birthday

import (
	"context"
	"fmt"
	"log"

	"github.com/bwmarrin/discordgo"
	"github.com/guergeiro/discord-bots/internal/env"
	"github.com/guergeiro/discord-bots/pkg/adapter/controller"
)

type BirthdayAdminController struct {
	base     *controller.BaseController[[]string]
	original controller.Controller[[]string]
}

func NewBirthdayAdminController(
	original controller.Controller[[]string],
) *BirthdayAdminController {
	return &BirthdayAdminController{
		base:     controller.NewBaseController[[]string](),
		original: original,
	}
}

func (c *BirthdayAdminController) Handle(
	ctx context.Context,
	args ...any,
) []string {
	log.Println("admin")
	if len(args) == 0 {
		return c.base.Handle(ctx, args...)
	}

	i, ok := args[0].(*discordgo.InteractionCreate)
	if !ok {
		return c.base.Handle(ctx, args...)
	}

	name := parseName(i.ApplicationCommandData().Options)
	if name != "admin" {
		return c.base.Handle(ctx, args...)
	}

	ADMIN_ID, err := env.Get("ADMIN_ID")
	if err != nil {
		log.Println(err.Error())
		return []string{
			err.Error(),
		}
	}

	caller := i.Interaction.Member.User.ID
	if caller != ADMIN_ID {
		return []string{
			fmt.Sprintf("Only <@%s> is able to use this command", ADMIN_ID),
		}
	}

	return c.original.Handle(
		ctx,
		i.ApplicationCommandData().Options[0].Options,
	)
}

func (c *BirthdayAdminController) SetNext(next controller.Controller[[]string]) {
	c.base.SetNext(next)
}
