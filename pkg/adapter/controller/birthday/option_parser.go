package birthday

import "github.com/bwmarrin/discordgo"

func parseName(
	options []*discordgo.ApplicationCommandInteractionDataOption,
) string {
	return options[0].Name
}
