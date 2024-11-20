package birthday

import (
	"context"
	"errors"
	"testing"

	"github.com/stretchr/testify/assert"
)

type fakeRemoveOneRepository struct {
	shouldError bool
}

func (m fakeRemoveOneRepository) RemoveOne(ctx context.Context, id string) error {
	if m.shouldError {
		return errors.New("Some error")
	}
	return nil
}

func TestNewRemoveBirthdayUseCase(t *testing.T) {
	fake := &fakeRemoveOneRepository{}

	useCase := NewRemoveBirthdayUseCase(fake)

	assert.Equal(t, fake, useCase.repository)
}

func TestRemoveBirthdayUseCaseNoIdError(t *testing.T) {
	fake := &fakeRemoveOneRepository{}

	useCase := NewRemoveBirthdayUseCase(fake)

	output, err := useCase.Execute(context.Background())
	assert.Error(t, err)
	assert.Empty(t, output)
}

func TestRemoveBirthdayUseCaseNoStringError(t *testing.T) {
	fake := &fakeRemoveOneRepository{}

	useCase := NewRemoveBirthdayUseCase(fake)

	output, err := useCase.Execute(context.Background(), 123)
	assert.Error(t, err)
	assert.Empty(t, output)
}

func TestRemoveBirthdayUseCaseRepositoryError(t *testing.T) {
	fake := &fakeRemoveOneRepository{
		shouldError: true,
	}

	useCase := NewRemoveBirthdayUseCase(fake)

	output, err := useCase.Execute(context.Background(), "some id")
	assert.Error(t, err)
	assert.Empty(t, output)
}

func TestRemoveBirthdayUseCaseSuccess(t *testing.T) {
	fake := &fakeRemoveOneRepository{
		shouldError: false,
	}

	useCase := NewRemoveBirthdayUseCase(fake)

	_, err := useCase.Execute(context.Background(), "some id")
	assert.Nil(t, err)
}
