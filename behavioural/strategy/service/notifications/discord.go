package notifications

import (
	"log/slog"

	"github.com/g-villarinho/design-patterns/behavioural/strategy/service/interfaces"
)

type DiscordNotificationStrategy struct {
}

func NewDiscordNotificationStrategy() interfaces.NotificationStrategy {
	return &DiscordNotificationStrategy{}
}

func (s *DiscordNotificationStrategy) SendNotification(channel string, destination string, message string) {
	logger := slog.With(
		slog.String("service", "notification"),
	)

	logger.Info("Notificação %s enviada para o Discord", message, destination)
}
