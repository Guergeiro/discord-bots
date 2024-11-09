package birthday

import (
	"context"
	"errors"
	"testing"
	"time"

	"github.com/guergeiro/discord-bots/pkg/domain/entity"
	"github.com/stretchr/testify/assert"
)

type fakeFindAllRepository struct {
	output      []entity.Birthday
	shouldError bool
}

func (m fakeFindAllRepository) FindAll(ctx context.Context) ([]entity.Birthday, error) {
	if m.shouldError {
		return []entity.Birthday{}, errors.New("Some error")
	}
	return m.output, nil
}

func TestNewAllBirthdayUseCase(t *testing.T) {
	fake := &fakeFindAllRepository{}

	useCase := NewAllBirthdayUseCase(fake)

	assert.Equal(t, fake, useCase.repository)
}

func TestAllBirthdayUseCaseRepositoryError(t *testing.T) {
	fake := &fakeFindAllRepository{
		shouldError: true,
	}

	useCase := NewAllBirthdayUseCase(fake)

	output, err := useCase.Execute(context.Background())
	assert.Error(t, err)
	assert.Empty(t, output)
}

func TestAllBirthdayUseCaseSuccess(t *testing.T) {
	expectedOutput := []entity.Birthday{
		{
			Id:   "1",
			Date: time.Date(2000, time.January, 15, 0, 0, 0, 0, time.UTC),
		},
		{
			Id:   "2",
			Date: time.Date(1995, time.March, 10, 0, 0, 0, 0, time.UTC),
		},
	}
	fake := &fakeFindAllRepository{
		shouldError: false,
		output:      expectedOutput,
	}

	useCase := NewAllBirthdayUseCase(fake)

	output, err := useCase.Execute(context.Background())
	assert.Nil(t, err)
	assert.Equal(t, expectedOutput, output)
}
