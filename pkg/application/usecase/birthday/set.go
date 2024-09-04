package birthday

import (
	"context"
	"errors"
	"time"

	"github.com/guergeiro/discord-bots/pkg/domain/entity"
	"github.com/guergeiro/discord-bots/pkg/domain/repository"
)

type SetBirthdayUseCase struct {
	repository repository.InsertOne[entity.Birthday]
}

func NewSetBirthdayUseCase(
	repository repository.InsertOne[entity.Birthday],
) SetBirthdayUseCase {
	return SetBirthdayUseCase{
		repository: repository,
	}
}

func (u SetBirthdayUseCase) Execute(
	ctx context.Context,
	args ...interface{},
) (entity.Birthday, error) {
	if len(args) != 2 {
		return entity.Birthday{}, errors.New("Requires Id and Date parameter")
	}
	id, ok := args[0].(string)
	if !ok {
		return entity.Birthday{}, errors.New("Requires Id must be of type string")
	}
	date, ok := args[1].(time.Time)
	if !ok {
		return entity.Birthday{}, errors.New("Requires Date must be of type time.Time")
	}
	birthday := entity.NewBirthday(id, date)
	err := u.repository.InsertOne(ctx, birthday)
	if err != nil {
		return entity.Birthday{}, err
	}

	return birthday, nil
}
