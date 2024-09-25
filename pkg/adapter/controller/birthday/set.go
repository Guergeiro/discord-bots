package birthday

import (
	"context"
	"fmt"
	"log"
	"time"

	"github.com/guergeiro/discord-bots/pkg/application/usecase"
	"github.com/guergeiro/discord-bots/pkg/domain/entity"
)

type BirthdaySetController struct {
	usecase usecase.UseCase[entity.Birthday]
}

func NewBirthdaySetController(
	usecase usecase.UseCase[entity.Birthday],
) BirthdaySetController {
	return BirthdaySetController{
		usecase: usecase,
	}
}

const timelayout = "2006-01-02"

func (c BirthdaySetController) Handle(
	ctx context.Context,
	id string,
	dateValue interface{},
) []string {
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
	birthday, err := c.usecase.Execute(ctx, id, date)
	if err != nil {
		log.Println(err.Error())
		return []string{err.Error()}
	}
	return []string{fmt.Sprintf(
		"<@%s> - %s", birthday.Id, birthday.PrettyBirthday(),
	)}
}
