package Services

import (
	"log"
	"net/smtp"
	"os"
	"path/filepath"
	"strings"
	"th3y3m/e-commerce-platform/Interface"

	"github.com/joho/godotenv"
)

type MailService struct {
	repository Interface.IUserRepository
}

func NewMailService(repository Interface.IUserRepository) Interface.IMailService {
	return &MailService{
		repository: repository,
	}
}

// SendMail sends the email to the user
func (m *MailService) SendMail(to string, token string) error {
	// Load environment variables
	err := godotenv.Load(".env")
	if err != nil {
		log.Printf("Error loading .env file: %v", err)
		return err
	}

	// Retrieve environment variables
	from, password := os.Getenv("EMAIL"), os.Getenv("PASSWORD")
	smtpHost, smtpPort := os.Getenv("SMTP_HOST"), os.Getenv("SMTP_PORT")

	// Construct the file path for the HTML template
	htmlPath := filepath.Join("Services", "Confirmation.html") // Adjust path as needed

	// Read the HTML template
	htmlTemplate, err := os.ReadFile(htmlPath)
	if err != nil {
		log.Printf("Failed to read HTML template: %v", err)
		return err
	}

	// Replace the {{TOKEN}} placeholder with the actual token
	htmlContent := strings.Replace(string(htmlTemplate), "{{TOKEN}}", token, 1)

	// Set up the email headers and body
	subject := "Subject: Verify your email\n"
	msg := []byte(subject + "MIME-version: 1.0;\nContent-Type: text/html; charset=\"UTF-8\";\n\n" + htmlContent)

	// Set up SMTP authentication
	auth := smtp.PlainAuth("", from, password, smtpHost)

	// Send the email
	err = smtp.SendMail(smtpHost+":"+smtpPort, auth, from, []string{to}, msg)
	if err != nil {
		log.Printf("Failed to send email to %s: %v", to, err)
		return err
	}

	return nil
}

// VerifyToken checks if the token is valid
func (m *MailService) VerifyToken(token string) bool {
	return m.repository.VerifyToken(token)
}
