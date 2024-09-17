package Services

import (
	"encoding/json"
	"errors"
	"fmt"
	"log"
	"net/url"
	"os"
	"strconv"
	"time"

	"th3y3m/e-commerce-platform/BusinessObjects"
	"th3y3m/e-commerce-platform/Repositories"
	"th3y3m/e-commerce-platform/Util"

	"github.com/joho/godotenv"
)

func NewMoMoService() (*MoMoService, error) {
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatalf("Error loading .env file: %v", err)
	}

	return &MoMoService{
		endpoint:    os.Getenv("MOMO_ENDPOINT"),
		secretKey:   os.Getenv("MOMO_SECRET_KEY"),
		accessKey:   os.Getenv("MOMO_ACCESS_KEY"),
		returnUrl:   os.Getenv("MOMO_RETURN_URL"),
		notifyUrl:   os.Getenv("MOMO_NOTIFY_URL"),
		partnerCode: os.Getenv("MOMO_PARTNER_CODE"),
		requestType: os.Getenv("MOMO_REQUEST_TYPE"),
		extraData:   os.Getenv("MOMO_EXTRA_DATA"),
	}, nil
}

type MoMoService struct {
	endpoint    string
	secretKey   string
	accessKey   string
	returnUrl   string
	notifyUrl   string
	partnerCode string
	requestType string
	extraData   string
}

// CreatePaymentUrl generates a payment URL for the given amount and order details.
func (s *MoMoService) CreateMoMoUrl(amount float64, orderId string) (string, error) {
	requestId := Util.GenerateID(10)
	orderInfo := "Customer"
	formattedAmount := int64(amount * 1000) // Convert to VND

	// Create raw signature string
	rawHash := fmt.Sprintf("accessKey=%s&amount=%d&extraData=%s&ipnUrl=%s&orderId=%s&orderInfo=%s&partnerCode=%s&redirectUrl=%s&requestId=%s&requestType=%s",
		s.accessKey, formattedAmount, s.extraData, s.notifyUrl, orderId, orderInfo, s.partnerCode, s.returnUrl, requestId, s.requestType)
	signature := Util.HmacSHA256(s.secretKey, rawHash)

	// Build request payload
	paymentRequest := map[string]interface{}{
		"partnerCode": s.partnerCode,
		"partnerName": "MoMo",
		"storeId":     "MoMoStore",
		"requestId":   requestId,
		"amount":      formattedAmount,
		"orderId":     orderId,
		"orderInfo":   orderInfo,
		"redirectUrl": s.returnUrl,
		"ipnUrl":      s.notifyUrl,
		"extraData":   s.extraData,
		"requestType": s.requestType,
		"signature":   signature,
		"lang":        "en",
	}

	// Send POST request to MoMo API
	response, err := Util.SendHttpRequest(s.endpoint, paymentRequest)
	if err != nil {
		return "", err
	}

	// Parse response and extract payment URL
	var jsonResponse map[string]interface{}
	if err := json.Unmarshal([]byte(response), &jsonResponse); err != nil {
		return "", err
	}

	if payUrl, ok := jsonResponse["payUrl"].(string); ok {
		return payUrl, nil
	}

	if message, ok := jsonResponse["message"].(string); ok {
		return "", fmt.Errorf("error creating payment URL: %s", message)
	}

	return "", errors.New("unexpected response from MoMo API")
}

func (s *MoMoService) ValidateMoMoResponse(queryString url.Values) (*TransactionStatusModel, error) {

	orderId := queryString.Get("orderId")
	resultCode := queryString.Get("resultCode")
	amount := queryString.Get("amount")
	signature := queryString.Get("signature")

	order, err := Repositories.GetOrderByID(orderId)
	if err != nil {
		return nil, err
	}

	if order.OrderID == "" || order.OrderStatus == "Complete" {
		return &TransactionStatusModel{
			IsSuccessful: false,
			RedirectUrl:  "https://localhost:3000/reject",
		}, nil
	}

	if resultCode == "0" {
		order.OrderStatus = "Complete"
		err = Repositories.UpdateOrder(order)
		if err != nil {
			return nil, err
		}
		paymentAmount, err := strconv.ParseFloat(amount, 64)
		if err != nil {
			return nil, fmt.Errorf("invalid payment amount: %v", err)
		}
		Transaction := &BusinessObjects.Transaction{
			TransactionID:    "TRA" + Util.GenerateID(10),
			OrderID:          order.OrderID,
			PaymentAmount:    paymentAmount,
			TransactionDate:  time.Now().In(TimeZoneAsiaHoChiMinh),
			PaymentStatus:    "Complete",
			PaymentSignature: signature,
			PaymentMethod:    "MoMo",
		}
		err = Repositories.CreateTransaction(*Transaction)
		if err != nil {
			return nil, err
		}

		return &TransactionStatusModel{
			IsSuccessful: true,
			RedirectUrl:  fmt.Sprintf("https://localhost:3000/confirm?orderId=%s", order.OrderID),
		}, nil
	}

	order.OrderStatus = "Cancelled"
	err = Repositories.UpdateOrder(order)
	if err != nil {
		return nil, err
	}

	return &TransactionStatusModel{
		IsSuccessful: false,
		RedirectUrl:  fmt.Sprintf("https://localhost:3000/reject?orderId=%s", orderId),
	}, nil
}
