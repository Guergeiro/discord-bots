package birthday

import (
	"context"
	"errors"
	"testing"

	"github.com/bwmarrin/discordgo"
	"github.com/guergeiro/discord-bots/pkg/adapter/controller"
	"github.com/guergeiro/discord-bots/pkg/domain/entity"
	"github.com/stretchr/testify/assert"
)

type fakeAnnouncerUseCase struct {
	shouldError bool
}

func (u fakeAnnouncerUseCase) Execute(
	ctx context.Context,
	args ...any,
) ([]entity.Birthday, error) {
	if u.shouldError {
		return []entity.Birthday{}, errors.New("Some error")
	}
	return []entity.Birthday{}, nil
}

type fakeAnnouncerPresenter struct {
	shouldError bool
}

func (p fakeAnnouncerPresenter) Present(
	ctx context.Context,
	input []entity.Birthday,
	args ...any,
) error {
	if p.shouldError {
		return errors.New("Some error")
	}
	return nil
}

func TestNewBirthdayAnnouncerController(t *testing.T) {
	controller := NewBirthdayAnnouncerController(
		fakeAnnouncerUseCase{},
		fakeAnnouncerPresenter{},
	)

	assert.NotNil(t, controller)
}

func TestNoArgumentsBirthdayAnnouncerController(t *testing.T) {
	controller := NewBirthdayAnnouncerController(
		fakeAnnouncerUseCase{},
		fakeAnnouncerPresenter{},
	)

	err := controller.Handle(context.Background())

	assert.Error(t, err)
}

func TestNoSessionArgumentBirthdayAnnouncerController(t *testing.T) {
	controller := NewBirthdayAnnouncerController(
		fakeAnnouncerUseCase{},
		fakeAnnouncerPresenter{},
	)

	err := controller.Handle(context.Background(), 123)

	assert.Error(t, err)
}

func TestErrorUseCaseBirthdayAnnouncerController(t *testing.T) {
	controller := NewBirthdayAnnouncerController(
		fakeAnnouncerUseCase{
			shouldError: true,
		},
		fakeAnnouncerPresenter{},
	)

	err := controller.Handle(
		context.Background(),
		&discordgo.Session{},
	)

	assert.Error(t, err)
}

func TestErrorPresenterBirthdayAnnouncerController(t *testing.T) {
	controller := NewBirthdayAnnouncerController(
		fakeAnnouncerUseCase{
			shouldError: false,
		},
		fakeAnnouncerPresenter{
			shouldError: true,
		},
	)

	err := controller.Handle(
		context.Background(),
		&discordgo.Session{},
	)

	assert.Error(t, err)
}

func TestSuccessBirthdayAnnouncerController(t *testing.T) {
	controller := NewBirthdayAnnouncerController(
		fakeAnnouncerUseCase{
			shouldError: false,
		},
		fakeAnnouncerPresenter{
			shouldError: false,
		},
	)

	err := controller.Handle(
		context.Background(),
		&discordgo.Session{},
	)

	assert.Nil(t, err)
}

func TestSetNextBirthdayAnnouncerController(t *testing.T) {
	cBase := controller.NewBaseController()

	controller := NewBirthdayAnnouncerController(
		fakeAnnouncerUseCase{},
		fakeAnnouncerPresenter{},
	)

	controller.SetNext(cBase)

	assert.NotNil(t, controller)
}
