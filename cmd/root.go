package cmd

import (
	"context"
	"fmt"
	"io"
	"log/slog"

	"github.com/alecthomas/kong"
)

type Root struct {
	Log struct {
		Level string `enum:"debug,info,warn,error,disable" default:"debug"`
	} `embed:"" prefix:"logging."`

	Discord Discord `cmd:"" help:"Use the Discord API to send a message to a user."`
}

func Run(ctx context.Context, stdin io.Reader, stdout io.Writer, args []string) error {
	var root Root
	parser, err := kong.New(
		&root,
		kong.BindTo(ctx, new(context.Context)),
		kong.BindTo(stdin, new(io.Reader)),
	)

	if err != nil {
		return err
	}
	parser.Stdout = stdout

	parsed, err := parser.Parse(args)
	if err != nil {
		return fmt.Errorf(`parsing args: %w`, err)
	}

	logger := slog.New(slog.NewJSONHandler(stdout, &slog.HandlerOptions{
		Level: Level(root.Log.Level),
	})).With("cmd", parsed.Command())

	if err := parsed.Run(logger); err != nil {
		return err
	}
	return nil
}

func Level(level string) slog.Level {
	switch level {
	case "debug":
		return slog.LevelDebug
	case "warn":
		return slog.LevelWarn
	case "error":
		return slog.LevelError
	case "disable":
		return slog.LevelError + 10
	case "info":
		fallthrough
	default:
		return slog.LevelInfo
	}
}
