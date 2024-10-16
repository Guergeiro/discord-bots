package birthday

import (
	"context"
	"log"
	"strings"

	"github.com/bwmarrin/discordgo"
	"github.com/guergeiro/discord-bots/internal/env"
	repository "github.com/guergeiro/discord-bots/internal/infra/birthday"
	"github.com/guergeiro/discord-bots/internal/infra/connection"
	"github.com/guergeiro/discord-bots/pkg/adapter/controller"
	controller_birthday "github.com/guergeiro/discord-bots/pkg/adapter/controller/birthday"
	usecase "github.com/guergeiro/discord-bots/pkg/application/usecase/birthday"
)

type Command struct {
	session   *discordgo.Session
	command   *discordgo.ApplicationCommand
	guildId   string
	channelId string
	adminId   string
}

func (c Command) Close() error {
	return c.session.ApplicationCommandDelete(
		c.session.State.User.ID,
		c.guildId,
		c.command.ID,
	)
}

func (c Command) Handler() func(
	*discordgo.Session,
	*discordgo.InteractionCreate,
) {

	controller := controller.NewControllerBuilder[[]string]().
		Add(
			controller_birthday.NewBirthdayAllController(
				usecase.NewAllBirthdayUseCase(
					repository.NewBirthdayPostgresRepository(
						connection.PostgresConn,
					),
				),
			),
		).
		Add(
			controller_birthday.NewBirthdaySetController(
				usecase.NewSetBirthdayUseCase(
					repository.NewBirthdayPostgresRepository(
						connection.PostgresConn,
					),
				),
			),
		).
		Add(
			controller_birthday.NewBirthdayRemoveController(
				usecase.NewRemoveBirthdayUseCase(
					repository.NewBirthdayPostgresRepository(
						connection.PostgresConn,
					),
				),
			),
		).
		Add(
			controller_birthday.NewBirthdayTodayController(
				usecase.NewTodayBirthdayUseCase(
					repository.NewBirthdayPostgresRepository(
						connection.PostgresConn,
					),
				),
			),
		).
		Add(
			controller_birthday.NewBirthdayAdminController(
				controller.NewControllerBuilder[[]string]().
					Add(
						controller_birthday.NewBirthdayAdminRetriggerController(
							usecase.NewTodayBirthdayUseCase(
								repository.NewBirthdayPostgresRepository(
									connection.PostgresConn,
								),
							),
						),
					).
					Add(
						controller_birthday.NewBirthdayAdminOthersBirthdayController(
							usecase.NewSetBirthdayUseCase(
								repository.NewBirthdayPostgresRepository(
									connection.PostgresConn,
								),
							),
						),
					).
					Build(),
			),
		).
		Build()

	return func(s *discordgo.Session, i *discordgo.InteractionCreate) {
		response := controller.Handle(context.Background(), i)
		log.Println(response)

		s.InteractionRespond(i.Interaction, &discordgo.InteractionResponse{
			Type: discordgo.InteractionResponseChannelMessageWithSource,
			Data: &discordgo.InteractionResponseData{
				Content: strings.Join(response, "\n"),
				Flags:   discordgo.MessageFlagsEphemeral,
			},
		})
	}
}

func CreateCommand(session *discordgo.Session) (*Command, error) {
	GUILD_ID, err := env.Get("GUILD_ID")
	if err != nil {
		return nil, err
	}
	CHANNEL_ID, err := env.Get("CHANNEL_ID")
	if err != nil {
		return nil, err
	}
	ADMIN_ID, err := env.Get("ADMIN_ID")
	if err != nil {
		return nil, err
	}
	cmd, err := session.ApplicationCommandCreate(session.State.User.ID, GUILD_ID, &discordgo.ApplicationCommand{
		Name:        "birth",
		Description: "All birth related options",
		Options: []*discordgo.ApplicationCommandOption{
			{
				Name:        "set",
				Description: "Set your own birthday",
				Type:        discordgo.ApplicationCommandOptionSubCommand,
				Options: []*discordgo.ApplicationCommandOption{
					{
						Type:        discordgo.ApplicationCommandOptionString,
						Name:        "date",
						Description: "YYYY-MM-DD",
						Required:    true,
					},
				},
			},
			{
				Name:        "remove",
				Description: "Remove your own birthday",
				Type:        discordgo.ApplicationCommandOptionSubCommand,
			},
			{
				Name:        "today",
				Description: "Displays today's birthdays",
				Type:        discordgo.ApplicationCommandOptionSubCommand,
			},
			{
				Name:        "all",
				Description: "Displays everyone's birthday",
				Type:        discordgo.ApplicationCommandOptionSubCommand,
			},
			{
				Name:        "admin",
				Description: "Subcommands group",
				Type:        discordgo.ApplicationCommandOptionSubCommandGroup,
				Options: []*discordgo.ApplicationCommandOption{
					{
						Name:        "retrigger",
						Description: "Rettriger today's annoucement",
						Type:        discordgo.ApplicationCommandOptionSubCommand,
					},
					{
						Name:        "others-birthday",
						Description: "Set other users birthdays",
						Type:        discordgo.ApplicationCommandOptionSubCommand,
						Options: []*discordgo.ApplicationCommandOption{
							{
								Type:        discordgo.ApplicationCommandOptionString,
								Name:        "date",
								Description: "YYYY-MM-DD",
								Required:    true,
							},
							{
								Type:        discordgo.ApplicationCommandOptionString,
								Name:        "user",
								Description: "User ID",
								Required:    true,
							},
						},
					},
				},
			},
		},
	})
	if err != nil {
		return nil, err
	}
	return &Command{
		session:   session,
		command:   cmd,
		guildId:   GUILD_ID,
		channelId: CHANNEL_ID,
		adminId:   ADMIN_ID,
	}, nil
}
