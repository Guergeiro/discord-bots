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

type fakeAllUseCase struct {
	shouldError bool
}

func (u fakeAllUseCase) Execute(
	ctx context.Context,
	args ...any,
) ([]entity.Birthday, error) {
	if u.shouldError {
		return []entity.Birthday{}, errors.New("Some error")
	}
	return []entity.Birthday{}, nil
}

type fakeAllPresenter struct {
	shouldError bool
}

func (p fakeAllPresenter) Present(
	ctx context.Context,
	input []entity.Birthday,
	args ...any,
) error {
	if p.shouldError {
		return errors.New("Some error")
	}
	return nil
}

func TestNewBirthdayAllController(t *testing.T) {
	controller := NewBirthdayAllController(
		fakeAllUseCase{},
		fakeAllPresenter{},
	)

	assert.NotNil(t, controller)
}

func TestNoArgumentsBirthdayAllController(t *testing.T) {
	controller := NewBirthdayAllController(
		fakeAllUseCase{},
		fakeAllPresenter{},
	)

	err := controller.Handle(context.Background())

	assert.Error(t, err)
}

func TestNoSessionArgumentBirthdayAllController(t *testing.T) {
	controller := NewBirthdayAllController(
		fakeAllUseCase{},
		fakeAllPresenter{},
	)

	err := controller.Handle(context.Background(), 123, 123)

	assert.Error(t, err)
}

func TestNoInteractionCreateArgumentBirthdayAllController(t *testing.T) {
	controller := NewBirthdayAllController(
		fakeAllUseCase{},
		fakeAllPresenter{},
	)

	err := controller.Handle(context.Background(), &discordgo.Session{}, 123)

	assert.Error(t, err)
}

func TestNotCorrectNameBirthdayAllController(t *testing.T) {
	controller := NewBirthdayAllController(
		fakeAllUseCase{},
		fakeAllPresenter{},
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

func TestErrorUseCaseBirthdayAllController(t *testing.T) {
	controller := NewBirthdayAllController(
		fakeAllUseCase{
			shouldError: true,
		},
		fakeAllPresenter{},
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
							Name: "all",
						},
					},
				},
			},
		},
	)

	assert.Error(t, err)
}

func TestErrorPresenterBirthdayAllController(t *testing.T) {
	controller := NewBirthdayAllController(
		fakeAllUseCase{
			shouldError: false,
		},
		fakeAllPresenter{
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
							Name: "all",
						},
					},
				},
			},
		},
	)

	assert.Error(t, err)
}

func TestSuccessBirthdayAllController(t *testing.T) {
	controller := NewBirthdayAllController(
		fakeAllUseCase{
			shouldError: false,
		},
		fakeAllPresenter{
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
							Name: "all",
						},
					},
				},
			},
		},
	)

	assert.Nil(t, err)
}

func TestSetNextBirthdayAllController(t *testing.T) {
	cBase := controller.NewBaseController()

	controller := NewBirthdayAllController(
		fakeAllUseCase{},
		fakeAllPresenter{},
	)

	controller.SetNext(cBase)

	assert.NotNil(t, controller)
}
