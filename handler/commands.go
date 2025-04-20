package handler

import (
	"fmt"
	"time"

	"github.com/disgoorg/disgo/discord"
	"github.com/disgoorg/disgo/handler"
)

var (
	Commands = []discord.ApplicationCommandCreate{
		discord.SlashCommandCreate{
			Name:        "ping",
			Description: "Replies with pong",
		},
		discord.SlashCommandCreate{
			Name:        "xkcd",
			Description: "Get a random XKCD comic",
		},
		discord.SlashCommandCreate{
			Name:        "today",
			Description: "Get all events you have today",
		},
		discord.SlashCommandCreate{
			Name:        "new",
			Description: "Create a new command",
			Options: []discord.ApplicationCommandOption{
				discord.ApplicationCommandOptionString{
					Name:        "name",
					Description: "Name of the event",
					Required:    true,
				},
				discord.ApplicationCommandOptionString{
					Name: "date",
					Description: fmt.Sprintf(
						"Date of the event (%s)", time.DateOnly,
					),
					Required: true,
				},
				discord.ApplicationCommandOptionString{
					Name:        "time",
					Description: "Time of the event (15:04)",
				},
				discord.ApplicationCommandOptionString{
					Name:        "description",
					Description: "Description of the event",
				},
				discord.ApplicationCommandOptionString{
					Name:        "recurrence",
					Description: "RRULE of the event (Advanced usecase in dev)",
				},
			},
		},
	}
)

func Ping(event *handler.CommandEvent) error {
	return event.CreateMessage(discord.MessageCreate{
		Content: "pong",
	})
}

func NotFound(event *handler.InteractionEvent) error {
	return event.CreateMessage(
		discord.MessageCreate{
			Content: "not found",
		},
	)
}
