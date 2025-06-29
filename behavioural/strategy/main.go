package main

import (
	"fmt"

	"github.com/g-villarinho/design-patterns/behavioural/strategy/service"
)

func main() {
	notificationService := service.NewNotificationService()

	channels := []string{"discord", "instagram", "twitter", "whatsapp", "email", "telegram"}

	for _, channel := range channels {
		fmt.Printf("Enviando notificação via %s...\n", channel)
		notificationService.Notify(channel, "usuario@exemplo.com", "Olá! Esta é uma mensagem de teste.")
		fmt.Println()
	}
}
