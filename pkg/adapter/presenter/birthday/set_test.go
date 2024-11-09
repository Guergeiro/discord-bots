package birthday

import (
	"context"
	"testing"

	"github.com/bwmarrin/discordgo"
	"github.com/guergeiro/discord-bots/pkg/domain/entity"
	"github.com/stretchr/testify/assert"
)

func TestNewBirthdaySetPresenter(t *testing.T) {
	presenter := NewBirthdaySetPresenter()

	assert.NotNil(t, presenter)
}

func TestNoArgumentsBirthdaySetPresenter(t *testing.T) {
	presenter := NewBirthdaySetPresenter()

	err := presenter.Present(context.Background(), entity.Birthday{})

	assert.Error(t, err)
}

func TestNoSessionArgumentBirthdaySetPresenter(t *testing.T) {
	presenter := NewBirthdaySetPresenter()

	err := presenter.Present(context.Background(), entity.Birthday{}, 123, 123)

	assert.Error(t, err)
}

func TestNoInteractionCreateArgumentBirthdaySetPresenter(t *testing.T) {
	presenter := NewBirthdaySetPresenter()

	err := presenter.Present(
		context.Background(),
		entity.Birthday{},
		&discordgo.Session{},
		123,
	)

	assert.Error(t, err)
}

func TestErrorBirthdaySetPresenter(t *testing.T) {
	defer func() {
		r := recover()
		assert.NotNil(t, r)
	}()

	presenter := NewBirthdaySetPresenter()

	err := presenter.Present(
		context.Background(),
		entity.Birthday{},
		&discordgo.Session{},
		&discordgo.InteractionCreate{},
	)

	assert.Error(t, err)
}
