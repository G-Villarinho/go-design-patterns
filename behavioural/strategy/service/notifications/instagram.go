package notifications

import (
	"log/slog"

	"github.com/g-villarinho/design-patterns/behavioural/strategy/service/interfaces"
)

type InstagramNotificationStrategy struct {
}

func NewInstagramNotificationStrategy() interfaces.NotificationStrategy {
	return &InstagramNotificationStrategy{}
}

func (s *InstagramNotificationStrategy) SendNotification(channel string, destination string, message string) {
	logger := slog.With(
		slog.String("service", "notification"),
	)

	logger.Info("Notificação %s enviada para o Instagram", message, destination)
}
