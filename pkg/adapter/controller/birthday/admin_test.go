package birthday

import (
	"context"
	"errors"
	"testing"

	"github.com/bwmarrin/discordgo"
	"github.com/guergeiro/discord-bots/pkg/adapter/controller"
	"github.com/stretchr/testify/assert"
)

type fakeOriginalController struct {
	shouldError bool
}

func (c fakeOriginalController) Handle(ctx context.Context, args ...any) error {
	if c.shouldError {
		return errors.New("Some error")
	}
	return nil
}

func (c fakeOriginalController) SetNext(next controller.Controller) {
}

func TestNewBirthdayAdminController(t *testing.T) {
	controller := NewBirthdayAdminController(
		fakeOriginalController{},
	)

	assert.NotNil(t, controller)
}

func TestNoArgumentsBirthdayAdminController(t *testing.T) {
	controller := NewBirthdayAdminController(
		fakeOriginalController{},
	)

	err := controller.Handle(context.Background())

	assert.Error(t, err)
}

func TestNoSessionArgumentBirthdayAdminController(t *testing.T) {
	controller := NewBirthdayAdminController(
		fakeOriginalController{},
	)

	err := controller.Handle(context.Background(), 123, 123)

	assert.Error(t, err)
}

func TestNoInteractionCreateArgumentBirthdayAdminController(t *testing.T) {
	controller := NewBirthdayAdminController(
		fakeOriginalController{},
	)

	err := controller.Handle(context.Background(), &discordgo.Session{}, 123)

	assert.Error(t, err)
}

func TestNotCorrectNameBirthdayAdminController(t *testing.T) {
	controller := NewBirthdayAdminController(
		fakeOriginalController{},
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

func TestNoEnvVariableBirthdayAdminController(t *testing.T) {
	controller := NewBirthdayAdminController(
		fakeOriginalController{},
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
							Name: "admin",
						},
					},
				},
			},
		},
	)

	assert.Error(t, err)
}

func TestDifferentEnvVariableBirthdayAdminController(t *testing.T) {
	admin := "123456789"
	t.Setenv("ADMIN_ID", admin)

	controller := NewBirthdayAdminController(
		fakeOriginalController{},
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
							Name: "admin",
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

func TestErrorOriginalBirthdayAdminController(t *testing.T) {
	admin := "123456789"
	t.Setenv("ADMIN_ID", admin)

	controller := NewBirthdayAdminController(
		fakeOriginalController{
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
							Name: "admin",
						},
					},
				},
				Member: &discordgo.Member{
					User: &discordgo.User{
						ID: admin,
					},
				},
			},
		},
	)

	assert.Error(t, err)
}

func TestSuccessBirthdayAdminController(t *testing.T) {
	admin := "123456789"
	t.Setenv("ADMIN_ID", admin)

	controller := NewBirthdayAdminController(
		fakeOriginalController{
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
							Name: "admin",
						},
					},
				},
				Member: &discordgo.Member{
					User: &discordgo.User{
						ID: admin,
					},
				},
			},
		},
	)

	assert.Nil(t, err)
}

func TestSetNextBirthdayAdminController(t *testing.T) {
	cBase := controller.NewBaseController()

	controller := NewBirthdayAdminController(
		fakeOriginalController{},
	)

	controller.SetNext(cBase)

	assert.NotNil(t, controller)
}
