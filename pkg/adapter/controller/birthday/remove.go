package birthday

import (
	"context"
	"fmt"
	"log"

	"github.com/bwmarrin/discordgo"
	"github.com/guergeiro/discord-bots/pkg/adapter/controller"
	"github.com/guergeiro/discord-bots/pkg/application/usecase"
)

type BirthdayRemoveController struct {
	base    *controller.BaseController[[]string]
	usecase usecase.UseCase[string]
}

func NewBirthdayRemoveController(
	usecase usecase.UseCase[string],
) *BirthdayRemoveController {
	return &BirthdayRemoveController{
		base:    controller.NewBaseController[[]string](),
		usecase: usecase,
	}
}

func (c *BirthdayRemoveController) Handle(
	ctx context.Context,
	args ...any,
) []string {
	log.Println("remove")
	if len(args) == 0 {
		return c.base.Handle(ctx, args...)
	}

	i, ok := args[0].(*discordgo.InteractionCreate)
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
		log.Println(err.Error())
		return []string{err.Error()}
	}
	return []string{fmt.Sprintf("Birthday removed")}
}

func (c *BirthdayRemoveController) SetNext(
	next controller.Controller[[]string],
) {
	c.base.SetNext(next)
}
