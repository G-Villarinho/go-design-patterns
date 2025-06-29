package notifications

import (
	"log/slog"

	"github.com/g-villarinho/design-patterns/behavioural/strategy/service/interfaces"
)

type TwitterNotificationStrategy struct {
}

func NewTwitterNotificationStrategy() interfaces.NotificationStrategy {
	return &TwitterNotificationStrategy{}
}

func (s *TwitterNotificationStrategy) SendNotification(channel string, destination string, message string) {
	logger := slog.With(
		slog.String("service", "notification"),
	)

	logger.Info("Notificação %s enviada para o Twitter", message, destination)
}
