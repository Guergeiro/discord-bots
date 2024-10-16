package birthday

import (
	"context"
	"strings"

	"github.com/bwmarrin/discordgo"
	repository "github.com/guergeiro/discord-bots/internal/infra/birthday"
	"github.com/guergeiro/discord-bots/internal/infra/connection"
	controller "github.com/guergeiro/discord-bots/pkg/adapter/controller/birthday"
	usecase "github.com/guergeiro/discord-bots/pkg/application/usecase/birthday"
)

func TodayCron(s *discordgo.Session, channelId string) {
	controller := controller.NewBirthdayTodayController(
		usecase.NewTodayBirthdayUseCase(
			repository.NewBirthdayPostgresRepository(
				connection.PostgresConn,
			),
		),
	)
	if response := controller.Handle(context.Background()); len(response) > 0 {
		header := []string{
			"Hey @everyone!",
			"These are today's birthday mabecos",
		}
		finalMessage := append(header, response...)
		s.ChannelMessageSendComplex(channelId, &discordgo.MessageSend{
			Content: strings.Join(finalMessage, "\n"),
			AllowedMentions: &discordgo.MessageAllowedMentions{
				Parse: []discordgo.AllowedMentionType{
					discordgo.AllowedMentionTypeEveryone,
				},
			},
		})
	}
}
