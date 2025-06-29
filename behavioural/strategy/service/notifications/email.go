package notifications

import (
	"log/slog"

	"github.com/g-villarinho/design-patterns/behavioural/strategy/service/interfaces"
)

type EmailNotificationStrategy struct {
}

func NewEmailNotificationStrategy() interfaces.NotificationStrategy {
	return &EmailNotificationStrategy{}
}

func (s *EmailNotificationStrategy) SendNotification(channel string, destination string, message string) {
	logger := slog.With(
		slog.String("service", "notification"),
	)

	logger.Info("Notificação %s enviada para o Email", message, destination)
}
