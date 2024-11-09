package birthday

import (
	"context"
	"errors"
	"testing"

	"github.com/bwmarrin/discordgo"
	"github.com/guergeiro/discord-bots/pkg/adapter/controller"
	"github.com/stretchr/testify/assert"
)

type fakeRemoveUseCase struct {
	shouldError bool
}

func (u fakeRemoveUseCase) Execute(
	ctx context.Context,
	args ...any,
) (string, error) {
	if u.shouldError {
		return "", errors.New("Some error")
	}
	return "", nil
}

type fakeRemovePresenter struct {
	shouldError bool
}

func (p fakeRemovePresenter) Present(
	ctx context.Context,
	input string,
	args ...any,
) error {
	if p.shouldError {
		return errors.New("Some error")
	}
	return nil
}

func TestNewBirthdayRemoveController(t *testing.T) {
	controller := NewBirthdayRemoveController(
		fakeRemoveUseCase{},
		fakeRemovePresenter{},
	)

	assert.NotNil(t, controller)
}

func TestNoArgumentsBirthdayRemoveController(t *testing.T) {
	controller := NewBirthdayRemoveController(
		fakeRemoveUseCase{},
		fakeRemovePresenter{},
	)

	err := controller.Handle(context.Background())

	assert.Error(t, err)
}

func TestNoSessionArgumentBirthdayRemoveController(t *testing.T) {
	controller := NewBirthdayRemoveController(
		fakeRemoveUseCase{},
		fakeRemovePresenter{},
	)

	err := controller.Handle(context.Background(), 123, 123)

	assert.Error(t, err)
}

func TestNoInteractionCreateArgumentBirthdayRemoveController(t *testing.T) {
	controller := NewBirthdayRemoveController(
		fakeRemoveUseCase{},
		fakeRemovePresenter{},
	)

	err := controller.Handle(context.Background(), &discordgo.Session{}, 123)

	assert.Error(t, err)
}

func TestNotCorrectNameBirthdayRemoveController(t *testing.T) {
	controller := NewBirthdayRemoveController(
		fakeRemoveUseCase{},
		fakeRemovePresenter{},
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
				Member: &discordgo.Member{
					User: &discordgo.User{
						ID: "some_string",
					},
				},
			},
		},
	)

	assert.Error(t, err)
}

func TestErrorUseCaseBirthdayRemoveController(t *testing.T) {
	controller := NewBirthdayRemoveController(
		fakeRemoveUseCase{
			shouldError: true,
		},
		fakeRemovePresenter{},
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
							Name: "remove",
						},
					},
				},
				Member: &discordgo.Member{
					User: &discordgo.User{
						ID: "some_string",
					},
				},
			},
		},
	)

	assert.Error(t, err)
}

func TestErrorPresenterBirthdayRemoveController(t *testing.T) {
	controller := NewBirthdayRemoveController(
		fakeRemoveUseCase{
			shouldError: false,
		},
		fakeRemovePresenter{
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
							Name: "remove",
						},
					},
				},
				Member: &discordgo.Member{
					User: &discordgo.User{
						ID: "some_string",
					},
				},
			},
		},
	)

	assert.Error(t, err)
}

func TestSuccessBirthdayRemoveController(t *testing.T) {
	controller := NewBirthdayRemoveController(
		fakeRemoveUseCase{
			shouldError: false,
		},
		fakeRemovePresenter{
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
							Name: "remove",
						},
					},
				},
				Member: &discordgo.Member{
					User: &discordgo.User{
						ID: "some_string",
					},
				},
			},
		},
	)

	assert.Nil(t, err)
}

func TestSetNextBirthdayRemoveController(t *testing.T) {
	cBase := controller.NewBaseController()

	controller := NewBirthdayRemoveController(
		fakeRemoveUseCase{},
		fakeRemovePresenter{},
	)

	controller.SetNext(cBase)

	assert.NotNil(t, controller)
}
