package birthday

import (
	"context"
	"fmt"

	"github.com/guergeiro/discord-bots/pkg/application/usecase"
)

type BirthdayRemoveController struct {
	usecase usecase.UseCase[string]
}

func NewBirthdayRemoveController(
	usecase usecase.UseCase[string],
) BirthdayRemoveController {
	return BirthdayRemoveController{
		usecase: usecase,
	}
}

func (c BirthdayRemoveController) Handle(
	ctx context.Context,
	id string,
) []string {
	_, err := c.usecase.Execute(ctx, id)
	if err != nil {
		return []string{err.Error()}
	}
	return []string{fmt.Sprintf("Birthday removed")}
}
