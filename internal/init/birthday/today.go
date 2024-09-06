package birthday

import (
	"context"
	"slices"
	"strings"

	"github.com/bwmarrin/discordgo"
	repository "github.com/guergeiro/discord-bots/internal/infra/birthday"
	"github.com/guergeiro/discord-bots/internal/infra/connection"
	controller "github.com/guergeiro/discord-bots/pkg/adapter/controller/birthday"
	usecase "github.com/guergeiro/discord-bots/pkg/application/usecase/birthday"
)

func TodayHandler(s *discordgo.Session, i *discordgo.InteractionCreate) {
	controller := controller.NewBirthdayTodayController(
		usecase.NewTodayBirthdayUseCase(
			repository.NewBirthdayPostgresRepository(
				connection.PostgresConn,
			),
		),
	)
	if response := controller.Handle(context.Background()); len(response) > 0 {

		response = slices.Insert(response, 0, "Today's birthdays:")
		s.InteractionRespond(i.Interaction, &discordgo.InteractionResponse{
			Type: discordgo.InteractionResponseChannelMessageWithSource,
			Data: &discordgo.InteractionResponseData{
				Content: strings.Join(response, "\n"),
				Flags:   discordgo.MessageFlagsEphemeral,
			},
		})
	} else {
		s.InteractionRespond(i.Interaction, &discordgo.InteractionResponse{
			Type: discordgo.InteractionResponseChannelMessageWithSource,
			Data: &discordgo.InteractionResponseData{
				Content: "No birthdays for today",
				Flags:   discordgo.MessageFlagsEphemeral,
			},
		})
	}

}

func TodayCron(s *discordgo.Session, channelId string) {
	controller := controller.NewBirthdayTodayController(
		usecase.NewTodayBirthdayUseCase(
			repository.NewBirthdayPostgresRepository(
				connection.PostgresConn,
			),
		),
	)
	if response := controller.Handle(context.Background()); len(response) > 0 {
		s.ChannelMessageSend(channelId, strings.Join(response, "\n"))
	}

}
