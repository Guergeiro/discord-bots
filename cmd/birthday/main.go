package main

import (
	"fmt"
	"log"
	"os"
	"os/signal"
	"syscall"

	"github.com/bwmarrin/discordgo"
	"github.com/guergeiro/discord-bots/internal/env"
	"github.com/guergeiro/discord-bots/internal/init/birthday"
)

func main() {
	DISCORD_TOKEN, err := env.Get("DISCORD_TOKEN")
	if err != nil {
		log.Println(err.Error())
		return
	}

	// Create a new Discord session using the provided bot token.
	session, err := discordgo.New("Bot " + DISCORD_TOKEN)
	if err != nil {
		log.Println("Error creating Discord session", err)
		return
	}

	session.AddHandler(func(s *discordgo.Session, r *discordgo.Ready) {
		log.Printf(
			"Logged in as: %v#%v\n",
			s.State.User.Username,
			s.State.User.Discriminator,
		)
	})

	err = session.Open()
	if err != nil {
		log.Println("Error opening connection", err)
		return
	}
	// Cleanly close down the Discord session.
	defer session.Close()

	cron, err := birthday.CreateCron(session)
	if err != nil {
		log.Println("Error creating cron", err)
		return
	}
	defer cron.Close()

	cmd, err := birthday.CreateCommand(session)
	if err != nil {
		log.Println("Error creating command", err)
		return
	}
	defer cmd.Close()
	handler := cmd.Handler()
	session.AddHandler(handler)

	sc := make(chan os.Signal, 1)
	signal.Notify(sc, syscall.SIGINT, syscall.SIGTERM, os.Interrupt)
	fmt.Println("Bot is now running. Press CTRL-C to exit.")
	<-sc
}
