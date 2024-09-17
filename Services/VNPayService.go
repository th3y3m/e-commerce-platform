package Services

import (
	"fmt"
	"log"
	"net"
	"net/url"
	"os"
	"sort"
	"strconv"
	"strings"
	"time"

	"th3y3m/e-commerce-platform/BusinessObjects"
	"th3y3m/e-commerce-platform/Repositories"
	"th3y3m/e-commerce-platform/Util"

	"github.com/joho/godotenv"
)

var TimeZoneAsiaHoChiMinh, _ = time.LoadLocation("Asia/Ho_Chi_Minh")

func NewVnpayService() *VnpayService {
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatalf("Error loading .env file: %v", err)
	}
	return &VnpayService{
		url:        os.Getenv("VNPAY_URL"),
		returnUrl:  os.Getenv("VNPAY_RETURN_URL"),
		tmnCode:    os.Getenv("VNPAY_TMNCODE"),
		hashSecret: os.Getenv("VNPAY_HASH_SECRET"),
	}
}

type TransactionStatusModel struct {
	IsSuccessful bool
	RedirectUrl  string
}

type VnpayService struct {
	url        string
	returnUrl  string
	tmnCode    string
	hashSecret string
}

func (s *VnpayService) CreateVNPayUrl(amount float64, orderinfor string) (string, error) {
	hostName, err := os.Hostname()
	if err != nil {
		return "", err
	}

	ipAddrs, err := net.LookupIP(hostName)
	if err != nil || len(ipAddrs) == 0 {
		return "", err
	}
	clientIPAddress := ipAddrs[0].String()

	pay := Util.NewPayLib()
	vnpAmount := amount * 100000
	pay.AddRequestData("vnp_Version", "2.1.0")
	pay.AddRequestData("vnp_Command", "pay")
	pay.AddRequestData("vnp_TmnCode", s.tmnCode)
	pay.AddRequestData("vnp_Amount", fmt.Sprintf("%.0f", vnpAmount))
	pay.AddRequestData("vnp_CreateDate", time.Now().Format("20060102150405"))
	pay.AddRequestData("vnp_IpAddr", clientIPAddress)
	pay.AddRequestData("vnp_OrderInfo", "Customer")
	pay.AddRequestData("vnp_ReturnUrl", s.returnUrl)
	pay.AddRequestData("vnp_TxnRef", orderinfor)

	TransactionUrl := pay.CreateRequestUrl(s.url, s.hashSecret)
	return TransactionUrl, nil
}

func (s *VnpayService) ValidateVNPayResponse(queryString url.Values) (*TransactionStatusModel, error) {

	vnpSecureHash := queryString.Get("vnp_SecureHash")
	vnpAmount := queryString.Get("vnp_Amount")
	queryString.Del("vnp_SecureHash")
	queryString.Del("vnp_SecureHashType")

	rawData := make([]string, 0, len(queryString))
	for key, val := range queryString {
		rawData = append(rawData, key+"="+strings.Join(val, ""))
	}
	sort.Strings(rawData)
	rawQueryString := strings.Join(rawData, "&")

	if !s.ValidateSignature(rawQueryString, vnpSecureHash, s.hashSecret) {
		return &TransactionStatusModel{IsSuccessful: false, RedirectUrl: "LINK_INVALID"}, nil
	}

	order, err := Repositories.GetOrderByID(queryString.Get("vnp_TxnRef"))
	if err != nil {
		return nil, err
	}

	if order.OrderStatus == "Complete" {
		return &TransactionStatusModel{
			IsSuccessful: false,
			RedirectUrl:  "LINK_INVALID",
		}, nil
	}

	vnpResponseCode := queryString.Get("vnp_ResponseCode")
	if vnpResponseCode == "00" && queryString.Get("vnp_TransactionStatus") == "00" {
		order.OrderStatus = "Complete"
		err = Repositories.UpdateOrder(order)
		if err != nil {
			return nil, err
		}

		paymentAmount, err := strconv.ParseFloat(vnpAmount, 64)
		if err != nil {
			return nil, fmt.Errorf("invalid payment amount: %v", err)
		}

		Transaction := &BusinessObjects.Transaction{
			TransactionID:    "TRA" + Util.GenerateID(10),
			OrderID:          order.OrderID,
			TransactionDate:  time.Now().In(TimeZoneAsiaHoChiMinh),
			PaymentAmount:    paymentAmount,
			PaymentStatus:    "Complete",
			PaymentSignature: queryString.Get("vnp_BankTranNo"),
			PaymentMethod:    "VNPay",
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
		RedirectUrl:  fmt.Sprintf("https://localhost:3000/reject?orderId=%s", queryString.Get("vnp_TxnRef")),
	}, nil
}

func (s *VnpayService) ValidateSignature(rspraw, inputHash, secretKey string) bool {
	return Util.HmacSHA512(secretKey, rspraw) == inputHash
}
