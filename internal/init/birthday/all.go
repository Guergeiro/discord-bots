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

func AllHandler(s *discordgo.Session, i *discordgo.InteractionCreate) {
	controller := controller.NewBirthdayAllController(
		usecase.NewAllBirthdayUseCase(
			repository.NewBirthdayPostgresRepository(
				connection.PostgresConn,
			),
		),
	)
	response := controller.Handle(context.Background())
	response = slices.Insert(response, 0, "These are all the birthdays:")
	s.InteractionRespond(i.Interaction, &discordgo.InteractionResponse{
		Type: discordgo.InteractionResponseChannelMessageWithSource,
		Data: &discordgo.InteractionResponseData{
			Content: strings.Join(response, "\n"),
			Flags: discordgo.MessageFlagsEphemeral,
		},
	})
}
