package birthday

import (
	"context"
	"errors"
	"github.com/bwmarrin/discordgo"
	"github.com/guergeiro/discord-bots/pkg/adapter/presenter"
	"github.com/guergeiro/discord-bots/pkg/domain/entity"
)

type AdminRetriggerPresenter struct {
	original presenter.Presenter[[]entity.Birthday]
}

func NewAdminRetriggerPresenter(
	original presenter.Presenter[[]entity.Birthday],
) *AdminRetriggerPresenter {
	return &AdminRetriggerPresenter{
		original: original,
	}
}

func (p *AdminRetriggerPresenter) Present(
	ctx context.Context,
	input []entity.Birthday,
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
	if len(input) == 0 {
		return s.InteractionRespond(i.Interaction, &discordgo.InteractionResponse{
			Type: discordgo.InteractionResponseChannelMessageWithSource,
			Data: &discordgo.InteractionResponseData{
				Content: "No birthdays for today\n",
				Flags:   discordgo.MessageFlagsEphemeral,
			},
		})
	}
	if err := p.original.Present(ctx, input, s); err != nil {
		return err
	}
	return s.InteractionRespond(i.Interaction, &discordgo.InteractionResponse{
		Type: discordgo.InteractionResponseChannelMessageWithSource,
		Data: &discordgo.InteractionResponseData{
			Content: "Message sent\n",
			Flags:   discordgo.MessageFlagsEphemeral,
		},
	})
}
