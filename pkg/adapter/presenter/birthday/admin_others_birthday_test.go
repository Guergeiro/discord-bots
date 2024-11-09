package birthday

import (
	"context"
	"testing"

	"github.com/bwmarrin/discordgo"
	"github.com/guergeiro/discord-bots/pkg/domain/entity"
	"github.com/stretchr/testify/assert"
)

func TestNewBirthdayAdminOthersBirthdayPresenter(t *testing.T) {
	presenter := NewBirthdayAdminOthersBirthdayPresenter()

	assert.NotNil(t, presenter)
}

func TestNoArgumentsBirthdayAdminOthersBirthdayPresenter(t *testing.T) {
	presenter := NewBirthdayAdminOthersBirthdayPresenter()

	err := presenter.Present(context.Background(), entity.Birthday{})

	assert.Error(t, err)
}

func TestNoSessionArgumentBirthdayAdminOthersBirthdayPresenter(t *testing.T) {
	presenter := NewBirthdayAdminOthersBirthdayPresenter()

	err := presenter.Present(context.Background(), entity.Birthday{}, 123, 123)

	assert.Error(t, err)
}

func TestNoInteractionCreateArgumentBirthdayAdminOthersBirthdayPresenter(t *testing.T) {
	presenter := NewBirthdayAdminOthersBirthdayPresenter()

	err := presenter.Present(
		context.Background(),
		entity.Birthday{},
		&discordgo.Session{},
		123,
	)

	assert.Error(t, err)
}

func TestErrorBirthdayAdminOthersBirthdayPresenter(t *testing.T) {
	defer func() {
		r := recover()
		assert.NotNil(t, r)
	}()

	presenter := NewBirthdayAdminOthersBirthdayPresenter()

	err := presenter.Present(
		context.Background(),
		entity.Birthday{},
		&discordgo.Session{},
		&discordgo.InteractionCreate{},
	)

	assert.Error(t, err)
}
