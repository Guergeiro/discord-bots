package birthday

import (
	"github.com/bwmarrin/discordgo"
	"github.com/guergeiro/discord-bots/internal/env"
)

type Command struct {
	session *discordgo.Session
	command *discordgo.ApplicationCommand
	guildId string
}

func (c Command) Close() error {
	return c.session.ApplicationCommandDelete(
		c.session.State.User.ID,
		c.guildId,
		c.command.ID,
	)
}

func (c Command) Handler() (func(*discordgo.Session, *discordgo.InteractionCreate), error) {
	CHANNEL_ID, err := env.Get("CHANNEL_ID")
	if err != nil {
		return nil, err
	}
	return func(s *discordgo.Session, i *discordgo.InteractionCreate) {
		options := i.ApplicationCommandData().Options
		switch options[0].Name {
		case "all":
			AllHandler(s, i)
			return
		case "set":
			SetHandler(s, i)
			return
		case "remove":
			RemoveHandler(s, i)
			return
		case "today":
			TodayHandler(s, i)
			return
		case "retrigger":
			if i.Interaction.Member.User.ID != "255102189616365568" {
				s.InteractionRespond(i.Interaction, &discordgo.InteractionResponse{
					Type: discordgo.InteractionResponseChannelMessageWithSource,
					Data: &discordgo.InteractionResponseData{
						Content: "Only <@255102189616365568> is able to use this command",
						Flags:   discordgo.MessageFlagsEphemeral,
					},
				})
				return
			}
			TodayCron(s, CHANNEL_ID)
			s.InteractionRespond(i.Interaction, &discordgo.InteractionResponse{
				Type: discordgo.InteractionResponseChannelMessageWithSource,
				Data: &discordgo.InteractionResponseData{
					Content: "Command sent my lord",
					Flags:   discordgo.MessageFlagsEphemeral,
				},
			})
			return
		}
	}, nil
}

func CreateCommand(session *discordgo.Session) (*Command, error) {
	GUILD_ID, err := env.Get("GUILD_ID")
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
				Name:        "retrigger",
				Description: "Admin command to today's birthdays announcement",
				Type:        discordgo.ApplicationCommandOptionSubCommand,
			},
		},
	})
	if err != nil {
		return nil, err
	}
	return &Command{
		session: session,
		command: cmd,
		guildId: GUILD_ID,
	}, nil
}
