package cmd

import (
  "context"
  "fmt"
  "log/slog"
  "io"

  "github.com/ohhfishal/notify/notifier"
)

type Discord struct {
  API string `env:"API" help:"API Token to access Discord"`
  AppID string `env:"APP_ID" help:"Discord App ID"`
}

func (d Discord) Run(ctx context.Context, logger *slog.Logger, stdin io.Reader) error {
  notifier := notifier.DiscordNotifier {
    API: d.API,
    AppID: d.AppID,
    Logger: logger,
  }

  data, err := io.ReadAll(stdin)
  if err != nil {
    return fmt.Errorf("failed to read stdin: %w", err)
  }
  return notifier.Notify(ctx, string(data))
}
