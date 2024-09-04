package birthday

import (
	"context"
	"errors"

	"github.com/guergeiro/discord-bots/pkg/domain/repository"
)

type RemoveBirthdayUseCase struct {
	repository repository.RemoveOne
}

func NewRemoveBirthdayUseCase(
	repository repository.RemoveOne,
) RemoveBirthdayUseCase {
	return RemoveBirthdayUseCase{
		repository: repository,
	}
}

func (u RemoveBirthdayUseCase) Execute(
	ctx context.Context,
	args ...interface{},
) (string, error) {
	if len(args) != 1 {
		return "", errors.New("Requires Id parameter")
	}
	id, ok := args[0].(string)
	if !ok {
		return "", errors.New("Requires Id must be of type string")
	}
	err := u.repository.RemoveOne(ctx, id)
	if err != nil {
		return "", err
	}

	return "REMOVED ONE RECORD", nil
}
