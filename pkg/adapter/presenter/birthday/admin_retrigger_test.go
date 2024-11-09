package birthday

import (
	"context"
	"errors"
	"testing"
	"time"

	"github.com/bwmarrin/discordgo"
	"github.com/guergeiro/discord-bots/pkg/domain/entity"
	"github.com/stretchr/testify/assert"
)

type fakeOriginalPresenter struct {
	shouldError bool
}

var birthdayInput = []entity.Birthday{
	{
		Id:   "1",
		Date: time.Date(2000, time.January, 15, 0, 0, 0, 0, time.UTC),
	},
	{
		Id:   "2",
		Date: time.Date(1995, time.March, 10, 0, 0, 0, 0, time.UTC),
	},
}

func (p fakeOriginalPresenter) Present(ctx context.Context, input []entity.Birthday, args ...any) error {
	if p.shouldError {
		return errors.New("Some error")
	}
	return nil
}

func TestNewAdminRetriggerPresenter(t *testing.T) {
	presenter := NewAdminRetriggerPresenter(fakeOriginalPresenter{})

	assert.NotNil(t, presenter)
}

func TestNoArgumentsAdminRetriggerPresenter(t *testing.T) {
	presenter := NewAdminRetriggerPresenter(fakeOriginalPresenter{})

	err := presenter.Present(context.Background(), birthdayInput)

	assert.Error(t, err)
}

func TestNoSessionArgumentAdminRetriggerPresenter(t *testing.T) {
	presenter := NewAdminRetriggerPresenter(fakeOriginalPresenter{})

	err := presenter.Present(context.Background(), birthdayInput, 123, 123)

	assert.Error(t, err)
}

func TestNoInteractionCreateArgumentAdminRetriggerPresenter(t *testing.T) {
	presenter := NewAdminRetriggerPresenter(fakeOriginalPresenter{})

	err := presenter.Present(
		context.Background(),
		birthdayInput,
		&discordgo.Session{},
		123,
	)

	assert.Error(t, err)
}

func TestErrorOriginalAdminRetriggerPresenter(t *testing.T) {
	presenter := NewAdminRetriggerPresenter(fakeOriginalPresenter{shouldError: true})

	err := presenter.Present(
		context.Background(),
		birthdayInput,
		&discordgo.Session{},
		&discordgo.InteractionCreate{},
	)

	assert.Error(t, err)
}

func TestErrorLenZeroAdminRetriggerPresenter(t *testing.T) {
	defer func() {
		r := recover()
		assert.NotNil(t, r)
	}()

	presenter := NewAdminRetriggerPresenter(fakeOriginalPresenter{})

	err := presenter.Present(
		context.Background(),
		[]entity.Birthday{},
		&discordgo.Session{},
		&discordgo.InteractionCreate{},
	)

	assert.Error(t, err)
}

func TestErrorAdminRetriggerPresenter(t *testing.T) {
	defer func() {
		r := recover()
		assert.NotNil(t, r)
	}()

	presenter := NewAdminRetriggerPresenter(fakeOriginalPresenter{})

	err := presenter.Present(
		context.Background(),
		birthdayInput,
		&discordgo.Session{},
		&discordgo.InteractionCreate{},
	)

	assert.Error(t, err)
}
