package notifications

import (
	"log/slog"

	"github.com/g-villarinho/design-patterns/behavioural/strategy/service/interfaces"
)

type WhatsAppNotificationStrategy struct {
}

func NewWhatsAppNotificationStrategy() interfaces.NotificationStrategy {
	return &WhatsAppNotificationStrategy{}
}

func (s *WhatsAppNotificationStrategy) SendNotification(channel string, destination string, message string) {
	logger := slog.With(
		slog.String("service", "notification"),
	)

	logger.Info("Notificação %s enviada para o WhatsApp", message, destination)
}
