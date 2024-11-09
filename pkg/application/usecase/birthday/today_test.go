package birthday

import (
	"context"
	"errors"
	"testing"
	"time"

	"github.com/guergeiro/discord-bots/pkg/domain/entity"
	"github.com/stretchr/testify/assert"
)

type fakeTodayRepository struct {
	shouldError bool
	output      []entity.Birthday
}

func (m fakeTodayRepository) FindByDate(ctx context.Context, date time.Time) ([]entity.Birthday, error) {
	if m.shouldError {
		return []entity.Birthday{}, errors.New("Some error")
	}
	return m.output, nil
}

func TestNewTodayBirthdayUseCase(t *testing.T) {
	fake := &fakeTodayRepository{}

	useCase := NewTodayBirthdayUseCase(fake)

	assert.Equal(t, fake, useCase.repository)
}

func TestTodayBirthdayUseCaseRepositoryError(t *testing.T) {
	fake := &fakeTodayRepository{
		shouldError: true,
	}

	useCase := NewTodayBirthdayUseCase(fake)

	output, err := useCase.Execute(context.Background())
	assert.Error(t, err)
	assert.Empty(t, output)
}

func TestTodayBirthdayUseCaseSuccess(t *testing.T) {
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

	fake := &fakeTodayRepository{
		shouldError: false,
		output:      expectedOutput,
	}

	useCase := NewTodayBirthdayUseCase(fake)

	output, err := useCase.Execute(context.Background())
	assert.Nil(t, err)
	assert.Equal(t, expectedOutput, output)
}
