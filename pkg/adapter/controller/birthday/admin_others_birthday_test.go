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

type fakeAdminOthersBirthdayUseCase struct {
	shouldError bool
}

func (u fakeAdminOthersBirthdayUseCase) Execute(
	ctx context.Context,
	args ...any,
) (entity.Birthday, error) {
	if u.shouldError {
		return entity.Birthday{}, errors.New("Some error")
	}
	return entity.Birthday{}, nil
}

type fakeAdminOthersBirthdayPresenter struct {
	shouldError bool
}

func (p fakeAdminOthersBirthdayPresenter) Present(
	ctx context.Context,
	input entity.Birthday,
	args ...any,
) error {
	if p.shouldError {
		return errors.New("Some error")
	}
	return nil
}

func TestNewBirthdayAdminOthersBirthdayController(t *testing.T) {
	controller := NewBirthdayAdminOthersBirthdayController(
		fakeAdminOthersBirthdayUseCase{},
		fakeAdminOthersBirthdayPresenter{},
	)

	assert.NotNil(t, controller)
}

func TestNoArgumentsBirthdayAdminOthersBirthdayController(t *testing.T) {
	controller := NewBirthdayAdminOthersBirthdayController(
		fakeAdminOthersBirthdayUseCase{},
		fakeAdminOthersBirthdayPresenter{},
	)

	err := controller.Handle(context.Background())

	assert.Error(t, err)
}

func TestNoSessionArgumentBirthdayAdminOthersBirthdayController(t *testing.T) {
	controller := NewBirthdayAdminOthersBirthdayController(
		fakeAdminOthersBirthdayUseCase{},
		fakeAdminOthersBirthdayPresenter{},
	)

	err := controller.Handle(context.Background(), 123, 123)

	assert.Error(t, err)
}

func TestNoInteractionCreateArgumentBirthdayAdminOthersBirthdayController(t *testing.T) {
	controller := NewBirthdayAdminOthersBirthdayController(
		fakeAdminOthersBirthdayUseCase{},
		fakeAdminOthersBirthdayPresenter{},
	)

	err := controller.Handle(context.Background(), &discordgo.Session{}, 123)

	assert.Error(t, err)
}

func TestNotCorrectNameBirthdayAdminOthersBirthdayController(t *testing.T) {
	controller := NewBirthdayAdminOthersBirthdayController(
		fakeAdminOthersBirthdayUseCase{},
		fakeAdminOthersBirthdayPresenter{},
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
											Value: "2000-01-02",
										},
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

func TestInvalidDateBirthdayAdminOthersBirthdayController(t *testing.T) {
	controller := NewBirthdayAdminOthersBirthdayController(
		fakeAdminOthersBirthdayUseCase{},
		fakeAdminOthersBirthdayPresenter{},
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
									Name: "others-birthday",
									Options: []*discordgo.ApplicationCommandInteractionDataOption{
										{
											Value: "qwer",
										},
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

func TestInvalidIdBirthdayAdminOthersBirthdayController(t *testing.T) {
	controller := NewBirthdayAdminOthersBirthdayController(
		fakeAdminOthersBirthdayUseCase{},
		fakeAdminOthersBirthdayPresenter{},
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
									Name: "others-birthday",
									Options: []*discordgo.ApplicationCommandInteractionDataOption{
										{
											Value: "2000-01-02",
										},
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

func TestErrorUseCaseBirthdayAdminOthersBirthdayController(t *testing.T) {
	controller := NewBirthdayAdminOthersBirthdayController(
		fakeAdminOthersBirthdayUseCase{
			shouldError: true,
		},
		fakeAdminOthersBirthdayPresenter{},
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
									Name: "others-birthday",
									Options: []*discordgo.ApplicationCommandInteractionDataOption{
										{
											Value: "2000-01-02",
										},
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

func TestErrorPresenterBirthdayAdminOthersBirthdayController(t *testing.T) {
	controller := NewBirthdayAdminOthersBirthdayController(
		fakeAdminOthersBirthdayUseCase{
			shouldError: false,
		},
		fakeAdminOthersBirthdayPresenter{
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
									Name: "others-birthday",
									Options: []*discordgo.ApplicationCommandInteractionDataOption{
										{
											Value: "2000-01-02",
										},
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

func TestSuccessBirthdayAdminOthersBirthdayController(t *testing.T) {
	controller := NewBirthdayAdminOthersBirthdayController(
		fakeAdminOthersBirthdayUseCase{
			shouldError: false,
		},
		fakeAdminOthersBirthdayPresenter{
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
									Name: "others-birthday",
									Options: []*discordgo.ApplicationCommandInteractionDataOption{
										{
											Value: "2000-01-02",
										},
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

func TestAdminOthersBirthdayNextBirthdayAdminOthersBirthdayController(t *testing.T) {
	cBase := controller.NewBaseController()

	controller := NewBirthdayAdminOthersBirthdayController(
		fakeAdminOthersBirthdayUseCase{},
		fakeAdminOthersBirthdayPresenter{},
	)

	controller.SetNext(cBase)

	assert.NotNil(t, controller)
}
