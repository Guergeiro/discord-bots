package birthday

import (
	"context"
	"testing"

	"github.com/bwmarrin/discordgo"
	"github.com/guergeiro/discord-bots/pkg/domain/entity"
	"github.com/stretchr/testify/assert"
)

func TestNewBirthdayAllPresenter(t *testing.T) {
	presenter := NewBirthdayAllPresenter()

	assert.NotNil(t, presenter)
}

func TestNoArgumentsBirthdayAllPresenter(t *testing.T) {
	presenter := NewBirthdayAllPresenter()

	err := presenter.Present(context.Background(), birthdayInput)

	assert.Error(t, err)
}

func TestNoSessionArgumentBirthdayAllPresenter(t *testing.T) {
	presenter := NewBirthdayAllPresenter()

	err := presenter.Present(context.Background(), birthdayInput, 123, 123)

	assert.Error(t, err)
}

func TestNoInteractionCreateArgumentBirthdayAllPresenter(t *testing.T) {
	presenter := NewBirthdayAllPresenter()

	err := presenter.Present(
		context.Background(),
		birthdayInput,
		&discordgo.Session{},
		123,
	)

	assert.Error(t, err)
}

func TestErrorLenZeroBirthdayAllPresenter(t *testing.T) {
	defer func() {
		r := recover()
		assert.NotNil(t, r)
	}()

	presenter := NewBirthdayAllPresenter()

	err := presenter.Present(
		context.Background(),
		[]entity.Birthday{},
		&discordgo.Session{},
		&discordgo.InteractionCreate{},
	)

	assert.Error(t, err)
}

func TestErrorBirthdayAllPresenter(t *testing.T) {
	defer func() {
		r := recover()
		assert.NotNil(t, r)
	}()

	presenter := NewBirthdayAllPresenter()

	err := presenter.Present(
		context.Background(),
		birthdayInput,
		&discordgo.Session{},
		&discordgo.InteractionCreate{},
	)

	assert.Error(t, err)
}
