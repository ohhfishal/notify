package handler

import (
	"errors"
	"fmt"
	"strings"

	"github.com/disgoorg/disgo/discord"
	"github.com/disgoorg/disgo/handler"
	"github.com/ohhfishal/schedule/cmd"
)

func GetSchedule(event *handler.CommandEvent) error {
	var stdout strings.Builder
	err := cmd.Run(event.Ctx, &stdout, []string{"get"})
	if err != nil {
		return fmt.Errorf(`running get: %w`, err)
	}
	return event.CreateMessage(discord.MessageCreate{
		Content: stdout.String(),
	})
}

func NewEvent(event *handler.CommandEvent) error {
	var stdout strings.Builder
	data := event.SlashCommandInteractionData()

	// Required fields
	name := data.String("name")
	if name == `` {
		return errors.New(`missing name`)
	}
	date := data.String("date")
	if date == `` {
		return errors.New(`missing date`)
	}
	args := []string{"new", name, date}

	// Optional fields
	if time := data.String("time"); time != `` {
		args = append(args, time)
	}
	if description := data.String("description"); description != `` {
		args = append(args, "--description", description)
	}
	if description := data.String("recurrence"); description != `` {
		args = append(args, "--recurrence", description)
	}

	err := cmd.Run(event.Ctx, &stdout, args)
	if err != nil {
		return fmt.Errorf(`running cmd: %w`, err)
	}
	return event.CreateMessage(discord.MessageCreate{
		Content: stdout.String(),
	})
}
