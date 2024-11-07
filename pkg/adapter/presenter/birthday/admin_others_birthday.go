package birthday

import (
	"context"
	"errors"
	"fmt"

	"github.com/bwmarrin/discordgo"
	"github.com/guergeiro/discord-bots/pkg/domain/entity"
)

type BirthdayAdminOthersBirthdayPresenter struct{}

func NewBirthdayAdminOthersBirthdayPresenter() *BirthdayAdminOthersBirthdayPresenter {
	return &BirthdayAdminOthersBirthdayPresenter{}
}

func (p *BirthdayAdminOthersBirthdayPresenter) Present(
	ctx context.Context,
	input entity.Birthday,
	args ...any,
) error {
	if len(args) != 2 {
		return errors.New("Requires exactly 2 arguments")
	}
	s, ok := args[0].(*discordgo.Session)
	if !ok {
		return errors.New("Can't cast 1st argument to *discordgo.Session")
	}
	i, ok := args[1].(*discordgo.InteractionCreate)
	if !ok {
		return errors.New(
			"Can't cast 2nd argument to *discordgo.InteractionCreate",
		)
	}
	return s.InteractionRespond(i.Interaction, &discordgo.InteractionResponse{
		Type: discordgo.InteractionResponseChannelMessageWithSource,
		Data: &discordgo.InteractionResponseData{
			Content: fmt.Sprintf(
				"<@%s> - %s\n", input.Id, input.PrettyBirthday(),
			),
			Flags: discordgo.MessageFlagsEphemeral,
		},
	})
}
