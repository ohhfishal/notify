package notifier

import (
  "context"
  "log/slog"
)

type DiscordNotifier struct {
  API string
  AppID string
  Logger *slog.Logger
}

func (dn *DiscordNotifier) Notify(ctx context.Context, msg string) error {
  dn.Logger.Debug("starting", "message", msg)
  return nil
}
