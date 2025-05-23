package birthday

import (
	"context"
	"errors"
	"fmt"
	"slices"
	"strings"

	"github.com/bwmarrin/discordgo"
	"github.com/guergeiro/discord-bots/internal/iter"
	"github.com/guergeiro/discord-bots/pkg/domain/entity"
)

type BirthdayAllPresenter struct{}

func NewBirthdayAllPresenter() *BirthdayAllPresenter {
	return &BirthdayAllPresenter{}
}

func (p *BirthdayAllPresenter) Present(
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
				Content: "No birthdays exist\n",
				Flags:   discordgo.MessageFlagsEphemeral,
			},
		})
	}
	output := slices.Collect(
		iter.Map(
			slices.Values(input),
			func(birthday entity.Birthday) string {
				return fmt.Sprintf(
					"<@%s> - %s", birthday.Id, birthday.PrettyBirthday(),
				)
			},
		),
	)
	header := []string{
		"These are all the birthdays:",
	}
	finalMessage := append(header, output...)
	return s.InteractionRespond(i.Interaction, &discordgo.InteractionResponse{
		Type: discordgo.InteractionResponseChannelMessageWithSource,
		Data: &discordgo.InteractionResponseData{
			Content: strings.Join(finalMessage, "\n"),
			Flags:   discordgo.MessageFlagsEphemeral,
		},
	})
}
