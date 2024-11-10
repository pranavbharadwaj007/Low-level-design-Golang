package main

import "fmt"

// NotificationChannel defines the interface for all notification methods
type NotificationChannel interface {
    Send(msg string)
}

// EmailService implements NotificationChannel
type EmailService struct{}

func (e *EmailService) Send(msg string) {
    fmt.Printf("Sending Email: %s\n", msg)
}

// SMSService implements NotificationChannel
type SMSService struct{}

func (s *SMSService) Send(msg string) {
    fmt.Printf("Sending SMS: %s\n", msg)
}

// NotificationService depends on the NotificationChannel interface
type NotificationService struct {
    channel NotificationChannel
}

func NewNotificationService(channel NotificationChannel) *NotificationService {
    return &NotificationService{
        channel: channel,
    }
}

func (n *NotificationService) Notify(msg string) {
    n.channel.Send(msg)
}

// Usage example for good implementation
func GoodExample() {
    // Using email notification
    emailService := &EmailService{}
    emailNotification := NewNotificationService(emailService)
    emailNotification.Notify("Your order has been shipped")

    // Using SMS notification
    smsService := &SMSService{}
    smsNotification := NewNotificationService(smsService)
    smsNotification.Notify("OTP 1234")

    // Easy to add new notification channels without modifying NotificationService
}

// Example of adding a new notification channel
type PushNotificationService struct{}

func (p *PushNotificationService) Send(msg string) {
    fmt.Printf("Sending Push Notification: %s\n", msg)
}

// Main function to demonstrate both implementations
func main() {

    fmt.Println("\nGood Implementation:")
    // Good implementation example
    emailNotification := NewNotificationService(&EmailService{})
    emailNotification.Notify("Order #123 confirmed")

    smsNotification := NewNotificationService(&SMSService{})
    smsNotification.Notify("Welcome to our service!")

    // Adding new notification type is easy
    pushNotification := NewNotificationService(&PushNotificationService{})
    pushNotification.Notify("New message received")
}
