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

type fakeAdminRetriggerUseCase struct {
	shouldError bool
}

func (u fakeAdminRetriggerUseCase) Execute(
	ctx context.Context,
	args ...any,
) ([]entity.Birthday, error) {
	if u.shouldError {
		return []entity.Birthday{}, errors.New("Some error")
	}
	return []entity.Birthday{}, nil
}

type fakeAdminRetriggerPresenter struct {
	shouldError bool
}

func (p fakeAdminRetriggerPresenter) Present(
	ctx context.Context,
	input []entity.Birthday,
	args ...any,
) error {
	if p.shouldError {
		return errors.New("Some error")
	}
	return nil
}

func TestNewBirthdayAdminRetriggerController(t *testing.T) {
	controller := NewBirthdayAdminRetriggerController(
		fakeAdminRetriggerUseCase{},
		fakeAdminRetriggerPresenter{},
	)

	assert.NotNil(t, controller)
}

func TestNoArgumentsBirthdayAdminRetriggerController(t *testing.T) {
	controller := NewBirthdayAdminRetriggerController(
		fakeAdminRetriggerUseCase{},
		fakeAdminRetriggerPresenter{},
	)

	err := controller.Handle(context.Background())

	assert.Error(t, err)
}

func TestNoSessionArgumentBirthdayAdminRetriggerController(t *testing.T) {
	controller := NewBirthdayAdminRetriggerController(
		fakeAdminRetriggerUseCase{},
		fakeAdminRetriggerPresenter{},
	)

	err := controller.Handle(context.Background(), 123, 123)

	assert.Error(t, err)
}

func TestNoInteractionCreateArgumentBirthdayAdminRetriggerController(t *testing.T) {
	controller := NewBirthdayAdminRetriggerController(
		fakeAdminRetriggerUseCase{},
		fakeAdminRetriggerPresenter{},
	)

	err := controller.Handle(context.Background(), &discordgo.Session{}, 123)

	assert.Error(t, err)
}

func TestNotCorrectNameBirthdayAdminRetriggerController(t *testing.T) {
	controller := NewBirthdayAdminRetriggerController(
		fakeAdminRetriggerUseCase{},
		fakeAdminRetriggerPresenter{},
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
							Options: []*discordgo.ApplicationCommandInteractionDataOption{
								{
									Name: "invalid",
								},
							},
						},
					},
				},
			},
		},
	)

	assert.Error(t, err)
}

func TestErrorUseCaseBirthdayAdminRetriggerController(t *testing.T) {
	controller := NewBirthdayAdminRetriggerController(
		fakeAdminRetriggerUseCase{
			shouldError: true,
		},
		fakeAdminRetriggerPresenter{},
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
							Options: []*discordgo.ApplicationCommandInteractionDataOption{
								{
									Name: "retrigger",
								},
							},
						},
					},
				},
			},
		},
	)

	assert.Error(t, err)
}

func TestErrorPresenterBirthdayAdminRetriggerController(t *testing.T) {
	controller := NewBirthdayAdminRetriggerController(
		fakeAdminRetriggerUseCase{
			shouldError: false,
		},
		fakeAdminRetriggerPresenter{
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
							Options: []*discordgo.ApplicationCommandInteractionDataOption{
								{
									Name: "retrigger",
								},
							},
						},
					},
				},
			},
		},
	)

	assert.Error(t, err)
}

func TestSuccessBirthdayAdminRetriggerController(t *testing.T) {
	controller := NewBirthdayAdminRetriggerController(
		fakeAdminRetriggerUseCase{
			shouldError: false,
		},
		fakeAdminRetriggerPresenter{
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
							Options: []*discordgo.ApplicationCommandInteractionDataOption{
								{
									Name: "retrigger",
								},
							},
						},
					},
				},
			},
		},
	)

	assert.Nil(t, err)
}

func TestSetNextBirthdayAdminRetriggerController(t *testing.T) {
	cBase := controller.NewBaseController()

	controller := NewBirthdayAdminRetriggerController(
		fakeAdminRetriggerUseCase{},
		fakeAdminRetriggerPresenter{},
	)

	controller.SetNext(cBase)

	assert.NotNil(t, controller)
}
