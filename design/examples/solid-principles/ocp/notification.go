package ocp

// BAD: Violates OCP - needs modification for new notification types
type NotificationServiceBad struct{}

func (s *NotificationServiceBad) Send(notificationType string, message string) {
	if notificationType == "email" {
		// send email
	} else if notificationType == "sms" {
		// send sms
	} else if notificationType == "push" {
		// send push notification
	}
	// Adding new type requires modifying this method!
}

// GOOD: Follows OCP - open for extension, closed for modification
type Notifier interface {
	Send(message string) error
}

type NotificationService struct {
	notifiers []Notifier
}

func (s *NotificationService) Notify(message string) {
	for _, notifier := range s.notifiers {
		notifier.Send(message)
	}
}

// New notification types can be added without modifying existing code
type EmailNotifier struct{}

func (n *EmailNotifier) Send(message string) error {
	// Send email
	return nil
}

type SMSNotifier struct{}

func (n *SMSNotifier) Send(message string) error {
	// Send SMS
	return nil
}

type PushNotifier struct{}

func (n *PushNotifier) Send(message string) error {
	// Send push notification
	return nil
}

// Can add SlackNotifier, TeamsNotifier, etc. without changing NotificationService
