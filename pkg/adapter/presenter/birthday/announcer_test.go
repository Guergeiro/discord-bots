package birthday

import (
	"context"
	"testing"

	"github.com/bwmarrin/discordgo"
	"github.com/guergeiro/discord-bots/pkg/domain/entity"
	"github.com/stretchr/testify/assert"
)

func TestNewBirthdayAnnouncerPresenter(t *testing.T) {
	presenter := NewBirthdayAnnouncerPresenter("123")

	assert.NotNil(t, presenter)
}

func TestNoArgumentsBirthdayAnnouncerPresenter(t *testing.T) {
	presenter := NewBirthdayAnnouncerPresenter("123")

	err := presenter.Present(context.Background(), birthdayInput)

	assert.Error(t, err)
}

func TestNoSessionArgumentBirthdayAnnouncerPresenter(t *testing.T) {
	presenter := NewBirthdayAnnouncerPresenter("123")

	err := presenter.Present(context.Background(), birthdayInput, 123)

	assert.Error(t, err)
}

func TestLenZeroBirthdayAnnouncerPresenter(t *testing.T) {
	presenter := NewBirthdayAnnouncerPresenter("123")

	err := presenter.Present(
		context.Background(),
		[]entity.Birthday{},
		&discordgo.Session{},
	)

	assert.Nil(t, err)
}

func TestErrorBirthdayAnnouncerPresenter(t *testing.T) {
	defer func() {
		r := recover()
		assert.NotNil(t, r)
	}()

	presenter := NewBirthdayAnnouncerPresenter("123")

	err := presenter.Present(
		context.Background(),
		birthdayInput,
		&discordgo.Session{},
	)

	assert.Error(t, err)
}
