package birthday

import (
	"context"
	"time"

	"github.com/guergeiro/discord-bots/pkg/domain/entity"
	"github.com/guergeiro/discord-bots/pkg/domain/repository"
)

type TodayBirthdayUseCase struct {
	repository repository.FindByDate[entity.Birthday]
}

func NewTodayBirthdayUseCase(
	repository repository.FindByDate[entity.Birthday],
) TodayBirthdayUseCase {
	return TodayBirthdayUseCase{
		repository: repository,
	}
}

func (u TodayBirthdayUseCase) Execute(
	ctx context.Context,
	args ...interface{},
) ([]entity.Birthday, error) {
	now := time.Now()
	today := time.Date(
		now.Year(),
		now.Month(),
		now.Day(),
		0,
		0,
		0,
		0,
		time.UTC,
	)
	birthdays, err := u.repository.FindByDate(ctx, today)
	if err != nil {
		return []entity.Birthday{}, err
	}

	return birthdays, nil
}
