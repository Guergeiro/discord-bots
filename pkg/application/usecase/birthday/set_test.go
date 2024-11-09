package birthday

import (
	"context"
	"errors"
	"testing"
	"time"

	"github.com/guergeiro/discord-bots/pkg/domain/entity"
	"github.com/stretchr/testify/assert"
)

type fakeSetRepository struct {
	shouldError bool
	output      entity.Birthday
}

func (m fakeSetRepository) InsertOne(ctx context.Context, input entity.Birthday) error {
	if m.shouldError {
		return errors.New("Some error")
	}
	return nil
}

func TestNewSetBirthdayUseCase(t *testing.T) {
	fake := &fakeSetRepository{}

	useCase := NewSetBirthdayUseCase(fake)

	assert.Equal(t, fake, useCase.repository)
}

func TestSetBirthdayUseCaseNoIdError(t *testing.T) {
	fake := &fakeSetRepository{}

	useCase := NewSetBirthdayUseCase(fake)

	output, err := useCase.Execute(context.Background())
	assert.Error(t, err)
	assert.Empty(t, output)
}

func TestSetBirthdayUseCaseNoStringError(t *testing.T) {
	fake := &fakeSetRepository{}

	useCase := NewSetBirthdayUseCase(fake)

	output, err := useCase.Execute(context.Background(), 123, time.Now())
	assert.Error(t, err)
	assert.Empty(t, output)
}

func TestSetBirthdayUseCaseNoTimeError(t *testing.T) {
	fake := &fakeSetRepository{}

	useCase := NewSetBirthdayUseCase(fake)

	output, err := useCase.Execute(context.Background(), "123", 321)
	assert.Error(t, err)
	assert.Empty(t, output)
}

func TestSetBirthdayUseCaseRepositoryError(t *testing.T) {
	expectedOutput := entity.NewBirthday("some id", time.Now())

	fake := &fakeSetRepository{
		shouldError: true,
	}

	useCase := NewSetBirthdayUseCase(fake)

	output, err := useCase.Execute(context.Background(), expectedOutput.Id, expectedOutput.Date)
	assert.Error(t, err)
	assert.Empty(t, output)
}

func TestSetBirthdayUseCaseSuccess(t *testing.T) {
	expectedOutput := entity.NewBirthday("some id", time.Now())
	fake := &fakeSetRepository{
		shouldError: false,
		output:      expectedOutput,
	}

	useCase := NewSetBirthdayUseCase(fake)

	output, err := useCase.Execute(context.Background(), expectedOutput.Id, expectedOutput.Date)
	assert.Nil(t, err)
	assert.Equal(t, expectedOutput, output)
}
