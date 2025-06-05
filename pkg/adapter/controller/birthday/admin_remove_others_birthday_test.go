package birthday

import (
	"context"
	"errors"
	"testing"

	"github.com/bwmarrin/discordgo"
	"github.com/guergeiro/discord-bots/pkg/adapter/controller"
	"github.com/stretchr/testify/assert"
)

type fakeAdminRemoveOthersBirthdayUseCase struct {
	shouldError bool
}

func (u fakeAdminRemoveOthersBirthdayUseCase) Execute(
	ctx context.Context,
	args ...any,
) (string, error) {
	if u.shouldError {
		return "", errors.New("Some error")
	}
	return "", nil
}

type fakeAdminRemoveOthersBirthdayPresenter struct {
	shouldError bool
}

func (p fakeAdminRemoveOthersBirthdayPresenter) Present(
	ctx context.Context,
	input string,
	args ...any,
) error {
	if p.shouldError {
		return errors.New("Some error")
	}
	return nil
}

func TestNewBirthdayAdminRemoveOthersBirthdayController(t *testing.T) {
	controller := NewBirthdayAdminRemoveOthersBirthdayController(
		fakeAdminRemoveOthersBirthdayUseCase{},
		fakeAdminRemoveOthersBirthdayPresenter{},
	)

	assert.NotNil(t, controller)
}

func TestNoArgumentsBirthdayAdminRemoveOthersBirthdayController(t *testing.T) {
	controller := NewBirthdayAdminRemoveOthersBirthdayController(
		fakeAdminRemoveOthersBirthdayUseCase{},
		fakeAdminRemoveOthersBirthdayPresenter{},
	)

	err := controller.Handle(context.Background())

	assert.Error(t, err)
}

func TestNoSessionArgumentBirthdayAdminRemoveOthersBirthdayController(t *testing.T) {
	controller := NewBirthdayAdminRemoveOthersBirthdayController(
		fakeAdminRemoveOthersBirthdayUseCase{},
		fakeAdminRemoveOthersBirthdayPresenter{},
	)

	err := controller.Handle(context.Background(), 123, 123)

	assert.Error(t, err)
}

func TestNoInteractionCreateArgumentBirthdayAdminRemoveOthersBirthdayController(t *testing.T) {
	controller := NewBirthdayAdminRemoveOthersBirthdayController(
		fakeAdminRemoveOthersBirthdayUseCase{},
		fakeAdminRemoveOthersBirthdayPresenter{},
	)

	err := controller.Handle(context.Background(), &discordgo.Session{}, 123)

	assert.Error(t, err)
}

func TestNotCorrectNameBirthdayAdminRemoveOthersBirthdayController(t *testing.T) {
	controller := NewBirthdayAdminRemoveOthersBirthdayController(
		fakeAdminRemoveOthersBirthdayUseCase{},
		fakeAdminRemoveOthersBirthdayPresenter{},
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
									Options: []*discordgo.ApplicationCommandInteractionDataOption{
										{
											Value: "<@123456789>",
										},
									},
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

func TestInvalidIdBirthdayAdminRemoveOthersBirthdayController(t *testing.T) {
	controller := NewBirthdayAdminRemoveOthersBirthdayController(
		fakeAdminRemoveOthersBirthdayUseCase{},
		fakeAdminRemoveOthersBirthdayPresenter{},
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
									Name: "remove-others-birthday",
									Options: []*discordgo.ApplicationCommandInteractionDataOption{
										{
											Value: 123,
										},
									},
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

func TestErrorUseCaseBirthdayAdminRemoveOthersBirthdayController(t *testing.T) {
	controller := NewBirthdayAdminRemoveOthersBirthdayController(
		fakeAdminRemoveOthersBirthdayUseCase{
			shouldError: true,
		},
		fakeAdminRemoveOthersBirthdayPresenter{},
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
									Name: "remove-others-birthday",
									Options: []*discordgo.ApplicationCommandInteractionDataOption{
										{
											Value: "<@123456789>",
										},
									},
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

func TestErrorPresenterBirthdayAdminRemoveOthersBirthdayController(t *testing.T) {
	controller := NewBirthdayAdminRemoveOthersBirthdayController(
		fakeAdminRemoveOthersBirthdayUseCase{
			shouldError: false,
		},
		fakeAdminRemoveOthersBirthdayPresenter{
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
									Name: "remove-others-birthday",
									Options: []*discordgo.ApplicationCommandInteractionDataOption{
										{
											Value: "<@123456789>",
										},
									},
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

func TestSuccessBirthdayAdminRemoveOthersBirthdayController(t *testing.T) {
	controller := NewBirthdayAdminRemoveOthersBirthdayController(
		fakeAdminRemoveOthersBirthdayUseCase{
			shouldError: false,
		},
		fakeAdminRemoveOthersBirthdayPresenter{
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
									Name: "remove-others-birthday",
									Options: []*discordgo.ApplicationCommandInteractionDataOption{
										{
											Value: "<@123456789>",
										},
									},
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

func TestAdminRemoveOthersBirthdayNextBirthdayAdminRemoveOthersBirthdayController(t *testing.T) {
	cBase := controller.NewBaseController()

	controller := NewBirthdayAdminRemoveOthersBirthdayController(
		fakeAdminRemoveOthersBirthdayUseCase{},
		fakeAdminRemoveOthersBirthdayPresenter{},
	)

	controller.SetNext(cBase)

	assert.NotNil(t, controller)
}
