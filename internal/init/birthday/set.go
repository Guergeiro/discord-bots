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

func SetHandler(s *discordgo.Session, i *discordgo.InteractionCreate) {
	date := i.ApplicationCommandData().Options[0].Options[0].Value
	id := i.Member.User.ID

	controller := controller.NewBirthdaySetController(
		usecase.NewSetBirthdayUseCase(
			repository.NewBirthdayPostgresRepository(
				connection.PostgresConn,
			),
		),
	)
	response := controller.Handle(context.Background(), id, date)
	s.InteractionRespond(i.Interaction, &discordgo.InteractionResponse{
		Type: discordgo.InteractionResponseChannelMessageWithSource,
		Data: &discordgo.InteractionResponseData{
			Content: strings.Join(response, "\n"),
			Flags:   discordgo.MessageFlagsEphemeral,
		},
	})
}
