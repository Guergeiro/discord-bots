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

type fakeSetUseCase struct {
	shouldError bool
}

func (u fakeSetUseCase) Execute(
	ctx context.Context,
	args ...any,
) (entity.Birthday, error) {
	if u.shouldError {
		return entity.Birthday{}, errors.New("Some error")
	}
	return entity.Birthday{}, nil
}

type fakeSetPresenter struct {
	shouldError bool
}

func (p fakeSetPresenter) Present(
	ctx context.Context,
	input entity.Birthday,
	args ...any,
) error {
	if p.shouldError {
		return errors.New("Some error")
	}
	return nil
}

func TestNewBirthdaySetController(t *testing.T) {
	controller := NewBirthdaySetController(
		fakeSetUseCase{},
		fakeSetPresenter{},
	)

	assert.NotNil(t, controller)
}

func TestNoArgumentsBirthdaySetController(t *testing.T) {
	controller := NewBirthdaySetController(
		fakeSetUseCase{},
		fakeSetPresenter{},
	)

	err := controller.Handle(context.Background())

	assert.Error(t, err)
}

func TestNoSessionArgumentBirthdaySetController(t *testing.T) {
	controller := NewBirthdaySetController(
		fakeSetUseCase{},
		fakeSetPresenter{},
	)

	err := controller.Handle(context.Background(), 123, 123)

	assert.Error(t, err)
}

func TestNoInteractionCreateArgumentBirthdaySetController(t *testing.T) {
	controller := NewBirthdaySetController(
		fakeSetUseCase{},
		fakeSetPresenter{},
	)

	err := controller.Handle(context.Background(), &discordgo.Session{}, 123)

	assert.Error(t, err)
}

func TestNotCorrectNameBirthdaySetController(t *testing.T) {
	controller := NewBirthdaySetController(
		fakeSetUseCase{},
		fakeSetPresenter{},
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

func TestInvalidDateBirthdaySetController(t *testing.T) {
	controller := NewBirthdaySetController(
		fakeSetUseCase{},
		fakeSetPresenter{},
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
							Name: "set",
							Options: []*discordgo.ApplicationCommandInteractionDataOption{
								{
									Value: "invalid_date",
								},
							},
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

func TestErrorUseCaseBirthdaySetController(t *testing.T) {
	controller := NewBirthdaySetController(
		fakeSetUseCase{
			shouldError: true,
		},
		fakeSetPresenter{},
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
							Name: "set",
							Options: []*discordgo.ApplicationCommandInteractionDataOption{
								{
									Value: "2000-01-02",
								},
							},
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

func TestErrorPresenterBirthdaySetController(t *testing.T) {
	controller := NewBirthdaySetController(
		fakeSetUseCase{
			shouldError: false,
		},
		fakeSetPresenter{
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
							Name: "set",
							Options: []*discordgo.ApplicationCommandInteractionDataOption{
								{
									Value: "2000-01-02",
								},
							},
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

func TestSuccessBirthdaySetController(t *testing.T) {
	controller := NewBirthdaySetController(
		fakeSetUseCase{
			shouldError: false,
		},
		fakeSetPresenter{
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
							Name: "set",
							Options: []*discordgo.ApplicationCommandInteractionDataOption{
								{
									Value: "2000-01-02",
								},
							},
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

func TestSetNextBirthdaySetController(t *testing.T) {
	cBase := controller.NewBaseController()

	controller := NewBirthdaySetController(
		fakeSetUseCase{},
		fakeSetPresenter{},
	)

	controller.SetNext(cBase)

	assert.NotNil(t, controller)
}
