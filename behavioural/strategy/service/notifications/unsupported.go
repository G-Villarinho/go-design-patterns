package notifications

import (
	"log/slog"

	"github.com/g-villarinho/design-patterns/behavioural/strategy/service/interfaces"
)

type UnsupportedNotificationStrategy struct {
}

func NewUnsupportedNotificationStrategy() interfaces.NotificationStrategy {
	return &UnsupportedNotificationStrategy{}
}

func (s *UnsupportedNotificationStrategy) SendNotification(channel string, destination string, message string) {
	logger := slog.With(
		slog.String("service", "notification"),
	)

	logger.Error("Canal de notificação não suportado", "channel", channel)
}
