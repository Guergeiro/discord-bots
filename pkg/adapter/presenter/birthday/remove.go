package birthday

import (
	"context"
	"errors"
	"fmt"

	"github.com/bwmarrin/discordgo"
)

type BirthdayRemovePresenter struct{}

func NewBirthdayRemovePresenter() *BirthdayRemovePresenter {
	return &BirthdayRemovePresenter{}
}

func (p *BirthdayRemovePresenter) Present(
	ctx context.Context,
	input string,
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
			Content: fmt.Sprintf("%s\n", input),
			Flags:   discordgo.MessageFlagsEphemeral,
		},
	})
}
