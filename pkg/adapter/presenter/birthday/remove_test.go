package birthday

import (
	"context"
	"testing"

	"github.com/bwmarrin/discordgo"
	"github.com/stretchr/testify/assert"
)

func TestNewBirthdayRemovePresenter(t *testing.T) {
	presenter := NewBirthdayRemovePresenter()

	assert.NotNil(t, presenter)
}

func TestNoArgumentsBirthdayRemovePresenter(t *testing.T) {
	presenter := NewBirthdayRemovePresenter()

	err := presenter.Present(context.Background(), "some string")

	assert.Error(t, err)
}

func TestNoSessionArgumentBirthdayRemovePresenter(t *testing.T) {
	presenter := NewBirthdayRemovePresenter()

	err := presenter.Present(context.Background(), "some string", 123, 123)

	assert.Error(t, err)
}

func TestNoInteractionCreateArgumentBirthdayRemovePresenter(t *testing.T) {
	presenter := NewBirthdayRemovePresenter()

	err := presenter.Present(
		context.Background(),
		"some string",
		&discordgo.Session{},
		123,
	)

	assert.Error(t, err)
}

func TestErrorBirthdayRemovePresenter(t *testing.T) {
	defer func() {
		r := recover()
		assert.NotNil(t, r)
	}()

	presenter := NewBirthdayRemovePresenter()

	err := presenter.Present(
		context.Background(),
		"some string",
		&discordgo.Session{},
		&discordgo.InteractionCreate{},
	)

	assert.Error(t, err)
}
