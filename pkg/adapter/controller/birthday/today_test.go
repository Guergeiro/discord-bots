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

type fakeTodayUseCase struct {
	shouldError bool
}

func (u fakeTodayUseCase) Execute(
	ctx context.Context,
	args ...any,
) ([]entity.Birthday, error) {
	if u.shouldError {
		return []entity.Birthday{}, errors.New("Some error")
	}
	return []entity.Birthday{}, nil
}

type fakeTodayPresenter struct {
	shouldError bool
}

func (p fakeTodayPresenter) Present(
	ctx context.Context,
	input []entity.Birthday,
	args ...any,
) error {
	if p.shouldError {
		return errors.New("Some error")
	}
	return nil
}

func TestNewBirthdayTodayController(t *testing.T) {
	controller := NewBirthdayTodayController(
		fakeTodayUseCase{},
		fakeTodayPresenter{},
	)

	assert.NotNil(t, controller)
}

func TestNoArgumentsBirthdayTodayController(t *testing.T) {
	controller := NewBirthdayTodayController(
		fakeTodayUseCase{},
		fakeTodayPresenter{},
	)

	err := controller.Handle(context.Background())

	assert.Error(t, err)
}

func TestNoSessionArgumentBirthdayTodayController(t *testing.T) {
	controller := NewBirthdayTodayController(
		fakeTodayUseCase{},
		fakeTodayPresenter{},
	)

	err := controller.Handle(context.Background(), 123, 123)

	assert.Error(t, err)
}

func TestNoInteractionCreateArgumentBirthdayTodayController(t *testing.T) {
	controller := NewBirthdayTodayController(
		fakeTodayUseCase{},
		fakeTodayPresenter{},
	)

	err := controller.Handle(context.Background(), &discordgo.Session{}, 123)

	assert.Error(t, err)
}

func TestNotCorrectNameBirthdayTodayController(t *testing.T) {
	controller := NewBirthdayTodayController(
		fakeTodayUseCase{},
		fakeTodayPresenter{},
	)

	err := controller.Handle(
		context.Background(),
		&discordgo.Session{},
		&discordgo.InteractionCreate{
			Interaction: &discordgo.Interaction{
				Type: discordgo.InteractionApplicationCommand,
				Data: discordgo.ApplicationCommandInteractionData{
					Options: []*discordgo.ApplicationCommandInteractionDataOption{
						{
							Name: "invalid",
						},
					},
				},
			},
		},
	)

	assert.Error(t, err)
}

func TestErrorUseCaseBirthdayTodayController(t *testing.T) {
	controller := NewBirthdayTodayController(
		fakeTodayUseCase{
			shouldError: true,
		},
		fakeTodayPresenter{},
	)

	err := controller.Handle(
		context.Background(),
		&discordgo.Session{},
		&discordgo.InteractionCreate{
			Interaction: &discordgo.Interaction{
				Type: discordgo.InteractionApplicationCommand,
				Data: discordgo.ApplicationCommandInteractionData{
					Options: []*discordgo.ApplicationCommandInteractionDataOption{
						{
							Name: "today",
						},
					},
				},
			},
		},
	)

	assert.Error(t, err)
}

func TestErrorPresenterBirthdayTodayController(t *testing.T) {
	controller := NewBirthdayTodayController(
		fakeTodayUseCase{
			shouldError: false,
		},
		fakeTodayPresenter{
			shouldError: true,
		},
	)

	err := controller.Handle(
		context.Background(),
		&discordgo.Session{},
		&discordgo.InteractionCreate{
			Interaction: &discordgo.Interaction{
				Type: discordgo.InteractionApplicationCommand,
				Data: discordgo.ApplicationCommandInteractionData{
					Options: []*discordgo.ApplicationCommandInteractionDataOption{
						{
							Name: "today",
						},
					},
				},
			},
		},
	)

	assert.Error(t, err)
}

func TestSuccessBirthdayTodayController(t *testing.T) {
	controller := NewBirthdayTodayController(
		fakeTodayUseCase{
			shouldError: false,
		},
		fakeTodayPresenter{
			shouldError: false,
		},
	)

	err := controller.Handle(
		context.Background(),
		&discordgo.Session{},
		&discordgo.InteractionCreate{
			Interaction: &discordgo.Interaction{
				Type: discordgo.InteractionApplicationCommand,
				Data: discordgo.ApplicationCommandInteractionData{
					Options: []*discordgo.ApplicationCommandInteractionDataOption{
						{
							Name: "today",
						},
					},
				},
			},
		},
	)

	assert.Nil(t, err)
}

func TestSetNextBirthdayTodayController(t *testing.T) {
	cBase := controller.NewBaseController()

	controller := NewBirthdayTodayController(
		fakeTodayUseCase{},
		fakeTodayPresenter{},
	)

	controller.SetNext(cBase)

	assert.NotNil(t, controller)
}
