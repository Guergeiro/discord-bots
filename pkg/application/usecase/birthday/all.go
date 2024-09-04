package birthday

import (
	"context"

	"github.com/guergeiro/discord-bots/pkg/domain/entity"
	"github.com/guergeiro/discord-bots/pkg/domain/repository"
)

type AllBirthdayUseCase struct {
	repository repository.FindAll[entity.Birthday]
}

func NewAllBirthdayUseCase(
	repository repository.FindAll[entity.Birthday],
) AllBirthdayUseCase {
	return AllBirthdayUseCase{
		repository: repository,
	}
}

func (u AllBirthdayUseCase) Execute(
	ctx context.Context,
	args ...interface{},
) ([]entity.Birthday, error) {
	birthdays, err := u.repository.FindAll(ctx)
	if err != nil {
		return []entity.Birthday{}, err
	}

	return birthdays, nil
}
