package service

import (
	"strings"

	"github.com/g-villarinho/design-patterns/behavioural/strategy/service/interfaces"
	"github.com/g-villarinho/design-patterns/behavioural/strategy/service/notifications"
)

type NotificationService struct {
	strategies map[string]interfaces.NotificationStrategy
}

func NewNotificationService() *NotificationService {
	strategies := map[string]interfaces.NotificationStrategy{
		"discord":   notifications.NewDiscordNotificationStrategy(),
		"instagram": notifications.NewInstagramNotificationStrategy(),
		"twitter":   notifications.NewTwitterNotificationStrategy(),
		"whatsapp":  notifications.NewWhatsAppNotificationStrategy(),
		"email":     notifications.NewEmailNotificationStrategy(),
	}

	return &NotificationService{
		strategies: strategies,
	}
}

func (s *NotificationService) Notify(channel string, destination string, message string) {
	channel = strings.ToLower(channel)

	strategy, exists := s.strategies[channel]
	if !exists {
		unsupportedStrategy := notifications.NewUnsupportedNotificationStrategy()
		unsupportedStrategy.SendNotification(channel, destination, message)
		return
	}

	strategy.SendNotification(channel, destination, message)
}
