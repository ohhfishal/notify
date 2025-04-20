package cmd

import (
	"context"
	"fmt"
	"log/slog"
	"os"
	"os/signal"
	"syscall"

	"github.com/disgoorg/disgo"
	"github.com/disgoorg/disgo/bot"
	"github.com/disgoorg/disgo/events"
	"github.com/disgoorg/disgo/gateway"
	"github.com/disgoorg/disgo/handler"
	"github.com/disgoorg/disgo/handler/middleware"
	"github.com/disgoorg/snowflake/v2"
	h "github.com/ohhfishal/notify/handler"
)

type Discord struct {
	Token   string `env:"TOKEN" help:"API Token to access Discord"`
	AppID   string `env:"APP_ID" help:"Discord App ID"`
	GuildID string `env:"GUILD_ID" help:"Discord App ID"`
}

func (d Discord) Run(ctx context.Context, logger *slog.Logger) error {
	r := handler.New()
	r.Use(middleware.Logger)
	r.Group(func(r handler.Router) {
		r.Command("/ping", h.Ping)
		r.Command("/xkcd", h.XKCD)
	})
	r.Group(func(r handler.Router) {
		r.Use(middleware.Print("schedule"))
		r.Command("/today", h.GetSchedule)
		r.Command("/new", h.NewEvent)
	})
	r.NotFound(h.NotFound)

	logger.Debug(`creating client`)
	client, err := disgo.New(d.Token,
		bot.WithDefaultGateway(),
		bot.WithGatewayConfigOpts(
			gateway.WithIntents(
				gateway.IntentDirectMessages,
				gateway.IntentGuilds,
				gateway.IntentGuildMessages,
				gateway.IntentDirectMessages,
			),
		),
		bot.WithEventListeners(r),
		bot.WithEventListenerFunc(func(e *events.Ready) {
			logger.Info("Ready")
		}),
	)
	if err != nil {
		return fmt.Errorf(`building bot: %w`, err)
	}
	defer client.Close(context.TODO())

	id, err := snowflake.Parse(d.GuildID)
	if err != nil {
		return fmt.Errorf(`parsing app id: %w`, err)
	}

	if err = handler.SyncCommands(client, h.Commands, []snowflake.ID{id}); err != nil {
		return fmt.Errorf(`syncing commands: %w`, err)
	}

	logger.Debug(`starting`)
	if err = client.OpenGateway(context.TODO()); err != nil {
		return fmt.Errorf(`running: %w`, err)
	}

	s := make(chan os.Signal, 1)
	signal.Notify(s, syscall.SIGINT, syscall.SIGTERM, os.Interrupt)
	<-s
	return nil
}
