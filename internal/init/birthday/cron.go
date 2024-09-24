package birthday

import (
	"github.com/bwmarrin/discordgo"
	"github.com/guergeiro/discord-bots/internal/env"
	"github.com/robfig/cron/v3"
)

type Cron struct {
	cron *cron.Cron
}

func (c Cron) Close() error {
	ctx := c.cron.Stop()
	<-ctx.Done()
	return ctx.Err()
}

func CreateCron(session *discordgo.Session) (*Cron, error) {
	CHANNEL_ID, err := env.Get("CHANNEL_ID")
	if err != nil {
		return nil, err
	}

	c := cron.New()
	c.AddFunc("0 9 * * *", func() {
		TodayCron(session, CHANNEL_ID)
	})
	c.Start()
	return &Cron{
		cron: c,
	}, nil
}
