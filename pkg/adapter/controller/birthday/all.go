package birthday

import (
	"context"
	"fmt"
	"log"
	"slices"

	"deedles.dev/xiter"
	"github.com/bwmarrin/discordgo"
	"github.com/guergeiro/discord-bots/pkg/adapter/controller"
	"github.com/guergeiro/discord-bots/pkg/application/usecase"
	"github.com/guergeiro/discord-bots/pkg/domain/entity"
)

type BirthdayAllController struct {
	base    *controller.BaseController[[]string]
	usecase usecase.UseCase[[]entity.Birthday]
}

func NewBirthdayAllController(
	usecase usecase.UseCase[[]entity.Birthday],
) *BirthdayAllController {
	return &BirthdayAllController{
		base:    controller.NewBaseController[[]string](),
		usecase: usecase,
	}
}

func (c *BirthdayAllController) Handle(
	ctx context.Context,
	args ...any,
) []string {
	log.Println("all")
	if len(args) == 0 {
		return c.base.Handle(ctx, args...)
	}

	i, ok := args[0].(*discordgo.InteractionCreate)
	if !ok {
		return c.base.Handle(ctx, args...)
	}

	name := parseName(i.ApplicationCommandData().Options)
	if name != "all" {
		return c.base.Handle(ctx, args...)
	}

	birthdays, err := c.usecase.Execute(ctx)
	if err != nil {
		log.Println(err.Error())
		return []string{err.Error()}
	}
	if len(birthdays) == 0 {
		return []string{
			"No birthdays exist",
		}
	}
	output := slices.Collect(
		xiter.Map(
			slices.Values(birthdays),
			func(birthday entity.Birthday) string {
				return fmt.Sprintf(
					"<@%s> - %s", birthday.Id, birthday.PrettyBirthday(),
				)
			},
		),
	)
	return slices.Insert(output, 0, "These are all the birthdays:")
}

func (c *BirthdayAllController) SetNext(next controller.Controller[[]string]) {
	c.base.SetNext(next)
}
