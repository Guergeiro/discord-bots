package birthday

import (
	"context"
	"fmt"
	"log"
	"regexp"
	"time"

	"github.com/bwmarrin/discordgo"
	"github.com/guergeiro/discord-bots/pkg/adapter/controller"
	"github.com/guergeiro/discord-bots/pkg/application/usecase"
	"github.com/guergeiro/discord-bots/pkg/domain/entity"
)

type BirthdayAdminOthersBirthdayController struct {
	base    *controller.BaseController[[]string]
	usecase usecase.UseCase[entity.Birthday]
}

func NewBirthdayAdminOthersBirthdayController(
	usecase usecase.UseCase[entity.Birthday],
) *BirthdayAdminOthersBirthdayController {
	return &BirthdayAdminOthersBirthdayController{
		base:    controller.NewBaseController[[]string](),
		usecase: usecase,
	}
}

var idExtractor = regexp.MustCompile(`\d+`)

func (c *BirthdayAdminOthersBirthdayController) Handle(
	ctx context.Context,
	args ...any,
) []string {
	log.Println("others-birthday")
	if len(args) == 0 {
		return c.base.Handle(ctx, args...)
	}

	i, ok := args[0].([]*discordgo.ApplicationCommandInteractionDataOption)
	if !ok {
		return c.base.Handle(ctx, args...)
	}

	dateValue := i[0].Options[0].Value

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

	idValue := i[0].Options[1].Value

	idStr, ok := idValue.(string)
	if !ok {
		log.Println("There was an error parsing the user id")
		return []string{"There was an error parsing the user id"}
	}
	id := idExtractor.FindString(idStr)

	birthday, err := c.usecase.Execute(ctx, id, date)
	if err != nil {
		log.Println(err.Error())
		return []string{err.Error()}
	}
	return []string{fmt.Sprintf(
		"<@%s> - %s", birthday.Id, birthday.PrettyBirthday(),
	)}
}

func (c *BirthdayAdminOthersBirthdayController) SetNext(
	next controller.Controller[[]string],
) {
	c.base.SetNext(next)
}
