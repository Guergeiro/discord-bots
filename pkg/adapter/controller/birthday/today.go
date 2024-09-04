package birthday

import (
	"context"
	"fmt"
	"slices"

	"deedles.dev/xiter"
	"github.com/guergeiro/discord-bots/pkg/application/usecase"
	"github.com/guergeiro/discord-bots/pkg/domain/entity"
)

type BirthdayTodayController struct {
	usecase usecase.UseCase[[]entity.Birthday]
}

func NewBirthdayTodayController(
	usecase usecase.UseCase[[]entity.Birthday],
) BirthdayTodayController {
	return BirthdayTodayController{
		usecase: usecase,
	}
}

func (c BirthdayTodayController) Handle(ctx context.Context) []string {
	birthdays, err := c.usecase.Execute(ctx)
	if err != nil {
		return []string{err.Error()}
	}
	if len(birthdays) == 0 {
		return []string{}
	}
	output := slices.Collect(
		xiter.Map(
			slices.Values(birthdays),
			func(birthday entity.Birthday) string {
				return fmt.Sprintf("%s - %s", birthday.Id, birthday.Date)
			},
		),
	)
	return output
}
