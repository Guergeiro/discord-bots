package birthday

import (
	"context"
	"testing"

	"github.com/bwmarrin/discordgo"
	"github.com/guergeiro/discord-bots/pkg/domain/entity"
	"github.com/stretchr/testify/assert"
)

func TestNewBirthdayTodayPresenter(t *testing.T) {
	presenter := NewBirthdayTodayPresenter()

	assert.NotNil(t, presenter)
}

func TestNoArgumentsBirthdayTodayPresenter(t *testing.T) {
	presenter := NewBirthdayTodayPresenter()

	err := presenter.Present(context.Background(), birthdayInput)

	assert.Error(t, err)
}

func TestNoSessionArgumentBirthdayTodayPresenter(t *testing.T) {
	presenter := NewBirthdayTodayPresenter()

	err := presenter.Present(context.Background(), birthdayInput, 123, 123)

	assert.Error(t, err)
}

func TestNoInteractionCreateArgumentBirthdayTodayPresenter(t *testing.T) {
	presenter := NewBirthdayTodayPresenter()

	err := presenter.Present(
		context.Background(),
		birthdayInput,
		&discordgo.Session{},
		123,
	)

	assert.Error(t, err)
}

func TestErrorLenZeroBirthdayTodayPresenter(t *testing.T) {
	defer func() {
		r := recover()
		assert.NotNil(t, r)
	}()

	presenter := NewBirthdayTodayPresenter()

	err := presenter.Present(
		context.Background(),
		[]entity.Birthday{},
		&discordgo.Session{},
		&discordgo.InteractionCreate{},
	)

	assert.Error(t, err)
}

func TestErrorBirthdayTodayPresenter(t *testing.T) {
	defer func() {
		r := recover()
		assert.NotNil(t, r)
	}()

	presenter := NewBirthdayTodayPresenter()

	err := presenter.Present(
		context.Background(),
		birthdayInput,
		&discordgo.Session{},
		&discordgo.InteractionCreate{},
	)

	assert.Error(t, err)
}
