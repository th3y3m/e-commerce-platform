package Services

import (
	"bytes"
	"log"
	"net/smtp"
	"os"
	"path/filepath"
	"strings"
	"text/template"
	"th3y3m/e-commerce-platform/BusinessObjects"
	"th3y3m/e-commerce-platform/Interface"

	"github.com/joho/godotenv"
)

type MailService struct {
	repository        Interface.IUserRepository
	productRepository Interface.IProductRepository
}

func NewMailService(repository Interface.IUserRepository, productRepository Interface.IProductRepository) Interface.IMailService {
	return &MailService{
		repository:        repository,
		productRepository: productRepository,
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

func (m *MailService) SendOrderDetails(Customer BusinessObjects.User,
	Order BusinessObjects.Order,
	OrderDetails []BusinessObjects.OrderDetail) error {

	err := godotenv.Load(".env")
	if err != nil {
		log.Printf("Error loading .env file: %v", err)
		return err
	}

	from, password := os.Getenv("EMAIL"), os.Getenv("PASSWORD")
	smtpHost, smtpPort := os.Getenv("SMTP_HOST"), os.Getenv("SMTP_PORT")

	// Populate the form data
	type OrderDetailWithProduct struct {
		OrderDetail BusinessObjects.OrderDetail
		Product     BusinessObjects.Product
	}

	var orderDetailsWithProduct []OrderDetailWithProduct
	for _, od := range OrderDetails {
		product, err := m.productRepository.GetProductByID(od.ProductID)
		if err != nil {
			log.Printf("Failed to get product details for product ID %s: %v", od.ProductID, err)
			return err
		}
		orderDetailsWithProduct = append(orderDetailsWithProduct, OrderDetailWithProduct{
			OrderDetail: od,
			Product:     product,
		})
	}

	form := struct {
		Customer                BusinessObjects.User
		Order                   BusinessObjects.Order
		OrderDetailsWithProduct []OrderDetailWithProduct
	}{
		Customer:                Customer,
		Order:                   Order,
		OrderDetailsWithProduct: orderDetailsWithProduct,
	}

	htmlPath := filepath.Join("Services", "Notifycation.html")

	htmlTemplate, err := os.ReadFile(htmlPath)
	if err != nil {
		log.Printf("Failed to read HTML template: %v", err)
		return err
	}

	tmpl, err := template.New("email").Funcs(template.FuncMap{
		"multiply": func(a, b interface{}) float64 {
			var af, bf float64
			switch v := a.(type) {
			case int:
				af = float64(v)
			case float64:
				af = v
			}
			switch v := b.(type) {
			case int:
				bf = float64(v)
			case float64:
				bf = v
			}
			return af * bf
		},
		"toFloat64": func(i interface{}) float64 {
			switch v := i.(type) {
			case int:
				return float64(v)
			case float64:
				return v
			default:
				return 0
			}
		},
	}).Parse(string(htmlTemplate))
	if err != nil {
		log.Printf("Failed to parse HTML template: %v", err)
		return err
	}

	var htmlContent bytes.Buffer
	err = tmpl.Execute(&htmlContent, form)
	if err != nil {
		log.Printf("Failed to execute HTML template: %v", err)
		return err
	}

	// Prepare email content
	subject := "Subject: Order Details\n"
	mime := "MIME-version: 1.0;\nContent-Type: text/html; charset=\"UTF-8\";\n\n"
	msg := []byte(subject + mime + htmlContent.String())

	// SMTP authentication
	auth := smtp.PlainAuth("", from, password, smtpHost)

	// Send email
	err = smtp.SendMail(smtpHost+":"+smtpPort, auth, from, []string{form.Customer.Email}, msg)
	if err != nil {
		log.Printf("Failed to send email to %s: %v", form.Customer.Email, err)
		return err
	}

	return nil
}

// VerifyToken checks if the token is valid
func (m *MailService) VerifyToken(token string) bool {
	return m.repository.VerifyToken(token)
}
