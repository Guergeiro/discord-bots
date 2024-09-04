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

func RemoveHandler(s *discordgo.Session, i *discordgo.InteractionCreate) {
	id := i.Member.User.ID

	controller := controller.NewBirthdayRemoveController(
		usecase.NewRemoveBirthdayUseCase(
			repository.NewBirthdayPostgresRepository(
				connection.PostgresConn,
			),
		),
	)
	response := controller.Handle(context.Background(), id)
	s.InteractionRespond(i.Interaction, &discordgo.InteractionResponse{
		Type: discordgo.InteractionResponseChannelMessageWithSource,
		Data: &discordgo.InteractionResponseData{
			Content: strings.Join(response, "\n"),
			Flags: discordgo.MessageFlagsEphemeral,
		},
	})
}
