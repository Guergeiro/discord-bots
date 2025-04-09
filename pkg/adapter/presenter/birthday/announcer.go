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

type BirthdayAnnouncerPresenter struct {
	channelId string
}

func NewBirthdayAnnouncerPresenter(
	channelId string,
) *BirthdayAnnouncerPresenter {
	return &BirthdayAnnouncerPresenter{
		channelId: channelId,
	}
}

func (p *BirthdayAnnouncerPresenter) Present(
	ctx context.Context,
	input []entity.Birthday,
	args ...any,
) error {
	if len(args) != 1 {
		return errors.New("Requires exactly 1 argument")
	}
	s, ok := args[0].(*discordgo.Session)
	if !ok {
		return errors.New("Can't cast 1st argument to *discordgo.Session")
	}
	if len(input) == 0 {
		// no birthdays
		return nil
	}
	header := []string{
		"Hey @everyone!",
		"These are today's birthday mabecos",
	}
	output := slices.Collect(
		iter.Map(
			slices.Values(input),
			func(birthday entity.Birthday) string {
				return fmt.Sprintf("<@%s>", birthday.Id)
			},
		),
	)
	finalMessage := append(header, output...)
	_, err := s.ChannelMessageSendComplex(p.channelId, &discordgo.MessageSend{
		Content: strings.Join(finalMessage, "\n"),
		AllowedMentions: &discordgo.MessageAllowedMentions{
			Parse: []discordgo.AllowedMentionType{
				discordgo.AllowedMentionTypeEveryone,
			},
		},
	})
	return err
}
