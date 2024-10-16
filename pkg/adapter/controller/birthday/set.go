package birthday

import (
	"context"
	"fmt"
	"log"
	"time"

	"github.com/bwmarrin/discordgo"
	"github.com/guergeiro/discord-bots/pkg/adapter/controller"
	"github.com/guergeiro/discord-bots/pkg/application/usecase"
	"github.com/guergeiro/discord-bots/pkg/domain/entity"
)

type BirthdaySetController struct {
	base    *controller.BaseController[[]string]
	usecase usecase.UseCase[entity.Birthday]
}

func NewBirthdaySetController(
	usecase usecase.UseCase[entity.Birthday],
) *BirthdaySetController {
	return &BirthdaySetController{
		base:    controller.NewBaseController[[]string](),
		usecase: usecase,
	}
}

const timelayout = "2006-01-02"

func (c *BirthdaySetController) Handle(
	ctx context.Context,
	args ...any,
) []string {
	log.Println("set")
	if len(args) == 0 {
		return c.base.Handle(ctx, args...)
	}

	i, ok := args[0].(*discordgo.InteractionCreate)
	if !ok {
		return c.base.Handle(ctx, args...)
	}
	name := parseName(i.ApplicationCommandData().Options)
	if name != "set" {
		return c.base.Handle(ctx, args...)
	}

	dateValue := i.ApplicationCommandData().Options[0].Options[0].Value

	dateStr, ok := dateValue.(string)
	if !ok {
		log.Println("There was an error parsing the date")
		return []string{"There was an error parsing the date"}
	}
	date, err := time.Parse(timelayout, dateStr)
	if err != nil {
		log.Println(err.Error())
		return []string{err.Error()}
	}
	id := i.Member.User.ID

	birthday, err := c.usecase.Execute(ctx, id, date)
	if err != nil {
		log.Println(err.Error())
		return []string{err.Error()}
	}
	return []string{fmt.Sprintf(
		"<@%s> - %s", birthday.Id, birthday.PrettyBirthday(),
	)}
}

func (c *BirthdaySetController) SetNext(next controller.Controller[[]string]) {
	c.base.SetNext(next)
}
