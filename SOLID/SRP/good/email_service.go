package email

type EmailService struct{}

func NewEmailService() *EmailService {
    return &EmailService{}
}

func (e *EmailService) SendEmailNotification() {
    println("Sending email notification for invoice")
}
