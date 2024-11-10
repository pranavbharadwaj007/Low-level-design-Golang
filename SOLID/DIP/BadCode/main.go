package main

import "fmt"

// EmailService handles email notifications
type EmailService struct{}

func (e *EmailService) SendEmail(message string) {
    fmt.Printf("Sending email: %s\n", message)
}

// SMSService handles SMS notifications
type SMSService struct{}

func (s *SMSService) SendSMS(message string) {
    fmt.Printf("Sending SMS: %s\n", message)
}

// NotificationService directly depends on concrete implementations
type NotificationService struct {
    emailService *EmailService
    smsService  *SMSService
}

func NewNotificationService() *NotificationService {
    return &NotificationService{
        emailService: &EmailService{},
        smsService:  &SMSService{},
    }
}

func (n *NotificationService) NotifyByEmail(msg string) {
    n.emailService.SendEmail(msg)
}

func (n *NotificationService) NotifyBySMS(msg string) {
    n.smsService.SendSMS(msg)
}

// Usage example for bad implementation
func BadExample() {
    service := NewNotificationService()
    service.NotifyByEmail("Your order is confirmed")
    service.NotifyBySMS("Your OTP is 1234")
}
