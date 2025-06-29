package interfaces

type NotificationStrategy interface {
	SendNotification(channel string, destination string, message string)
}
