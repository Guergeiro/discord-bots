package birthday

import (
	"context"
	"log"

	"github.com/bwmarrin/discordgo"
	"github.com/guergeiro/discord-bots/internal/env"
	repository "github.com/guergeiro/discord-bots/internal/infra/birthday"
	"github.com/guergeiro/discord-bots/internal/infra/connection"
	"github.com/guergeiro/discord-bots/pkg/adapter/controller"
	controller_birthday "github.com/guergeiro/discord-bots/pkg/adapter/controller/birthday"
	presenter_birthday "github.com/guergeiro/discord-bots/pkg/adapter/presenter/birthday"
	usecase "github.com/guergeiro/discord-bots/pkg/application/usecase/birthday"
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

	controller := controller.NewControllerBuilder().
		Add(
			controller_birthday.NewBirthdayAnnouncerController(
				usecase.NewTodayBirthdayUseCase(
					repository.NewBirthdayPostgresRepository(
						connection.PostgresConn,
					),
				),
				presenter_birthday.NewBirthdayAnnouncerPresenter(CHANNEL_ID),
			),
		).
		Build()

	c := cron.New()
	c.AddFunc("0 9 * * *", func() {
		if err := controller.Handle(context.Background(), session); err != nil {
			log.Println(err)
		}
	})
	c.Start()
	return &Cron{
		cron: c,
	}, nil
}
