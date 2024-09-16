package Services

// import (
// 	"fmt"
// 	"net"
// 	"net/url"
// 	"os"
// 	"strings"
// 	"time"

// 	"th3y3m/e-commerce-platform/BusinessObjects"
// 	"th3y3m/e-commerce-platform/Repositories"
// 	"th3y3m/e-commerce-platform/Util"
// )

// type TransactionStatusModel struct {
// 	IsSuccessful bool
// 	RedirectUrl  string
// }

// type VnpayService struct {
// 	url        string
// 	returnUrl  string
// 	tmnCode    string
// 	hashSecret string
// }

// func NewVnpayService() *VnpayService {
// 	return &VnpayService{
// 		url:        "https://sandbox.vnTransaction.vn/Transactionv2/vpcpay.html",
// 		returnUrl:  "https://localhost:7173/VNpayAPI/Transactionconfirm",
// 		tmnCode:    "FKUXJX95",
// 		hashSecret: "0D3EAMNJYSY9INENB5JYP8XW2U8MD8WE",
// 	}
// }

// func (s *VnpayService) CreateTransactionUrl(amount float64, infor, orderinfor string) (string, error) {
// 	hostName, err := os.Hostname()
// 	if err != nil {
// 		return "", err
// 	}

// 	ipAddrs, err := net.LookupIP(hostName)
// 	if err != nil || len(ipAddrs) == 0 {
// 		return "", err
// 	}
// 	clientIPAddress := ipAddrs[0].String()

// 	pay := Util.NewPayLib()
// 	vnpAmount := amount * 100000
// 	pay.AddRequestData("vnp_Version", "2.1.0")
// 	pay.AddRequestData("vnp_Command", "pay")
// 	pay.AddRequestData("vnp_TmnCode", s.tmnCode)
// 	pay.AddRequestData("vnp_Amount", fmt.Sprintf("%.0f", vnpAmount))
// 	pay.AddRequestData("vnp_BankCode", "")
// 	pay.AddRequestData("vnp_CreateDate", time.Now().Format("20060102150405"))
// 	pay.AddRequestData("vnp_CurrCode", "VND")
// 	pay.AddRequestData("vnp_IpAddr", clientIPAddress)
// 	pay.AddRequestData("vnp_Locale", "vn")
// 	pay.AddRequestData("vnp_OrderInfo", infor)
// 	pay.AddRequestData("vnp_OrderType", "other")
// 	pay.AddRequestData("vnp_ReturnUrl", s.returnUrl)
// 	pay.AddRequestData("vnp_TxnRef", orderinfor)

// 	TransactionUrl := pay.CreateRequestUrl(s.url, s.hashSecret)
// 	return TransactionUrl, nil
// }

// func (s *VnpayService) ValidateTransactionResponse(queryString string) (*TransactionStatusModel, error) {
// 	defer func() {
// 		if r := recover(); r != nil {
// 		}
// 	}()

// 	timeZone, err := time.LoadLocation("Asia/Ho_Chi_Minh")
// 	if err != nil {
// 		return nil, err
// 	}
// 	nowUtc := time.Now().UTC()

// 	values, err := url.ParseQuery(queryString)
// 	if err != nil {
// 		return nil, err
// 	}

// 	order, err := Repositories.GetOrderByID(values.Get("vnp_TxnRef"))
// 	if err != nil {
// 		return nil, err
// 	}

// 	vnpResponseCode := values.Get("vnp_ResponseCode")
// 	vnpSecureHash := values.Get("vnp_SecureHash")
// 	pos := strings.Index(queryString, "&vnp_SecureHash")
// 	checkSignature := s.ValidateSignature(queryString[1:pos], vnpSecureHash, s.hashSecret)

// 	if order.OrderStatus == "true" && order.OrderID != "" {
// 		return &TransactionStatusModel{
// 			IsSuccessful: false,
// 			RedirectUrl:  "LINK_INVALID",
// 		}, nil
// 	}

// 	if checkSignature && s.tmnCode == values.Get("vnp_TmnCode") {
// 		orderid := values.Get("vnp_TxnRef")

// 		if vnpResponseCode == "00" && values.Get("vnp_TransactionStatus") == "00" {
// 			Transaction := &BusinessObjects.Transaction{
// 				TransactionID: "P" + Util.GenerateID(10),
// 				OrderID:       orderid,
// 				// PaymentAmount:    float64(values.Get("vnp_Amount")) / 100000,
// 				TransactionDate:  nowUtc.In(timeZone),
// 				PaymentStatus:    "Complete",
// 				PaymentSignature: values.Get("vnp_BankTranNo"),
// 				PaymentMethod:    "VNPay",
// 			}
// 			err := Repositories.CreateTransaction(*Transaction)
// 			if err != nil {
// 				return nil, err
// 			}

// 			order.OrderStatus = "Complete"
// 			err = Repositories.UpdateOrder(order)
// 			if err != nil {
// 				return nil, err
// 			}

// 			return &TransactionStatusModel{
// 				IsSuccessful: true,
// 				RedirectUrl:  fmt.Sprintf("https://localhost:3000/confirm?orderId=%s", values.Get("vnp_TxnRef")),
// 			}, nil
// 		} else {
// 			// amount := float64(values.Get("vnp_Amount"))
// 			if values.Get("vnp_BankTranNo") != "" || values.Get("vnp_TxnRef") != "" {
// 				order.OrderStatus = "cancel"
// 				err = Repositories.UpdateOrder(order)
// 				if err != nil {
// 					return nil, err
// 				}

// 				Transaction := &BusinessObjects.Transaction{
// 					TransactionID: "P" + Util.GenerateID(10),
// 					OrderID:       orderid,
// 					// PaymentAmount:   amount / 100000,
// 					TransactionDate: nowUtc.In(timeZone),
// 					PaymentStatus:   "Fail",
// 					PaymentMethod:   "VNPay",
// 				}
// 				err = Repositories.CreateTransaction(*Transaction)
// 				if err != nil {
// 					return nil, err
// 				}
// 			}
// 			return &TransactionStatusModel{
// 				IsSuccessful: false,
// 				RedirectUrl:  fmt.Sprintf("https://localhost:3000/reject?orderId=%s", values.Get("vnp_TxnRef")),
// 			}, nil
// 		}
// 	} else {
// 		return &TransactionStatusModel{
// 			IsSuccessful: false,
// 			RedirectUrl:  "LINK_INVALID",
// 		}, nil
// 	}
// }

// func (s *VnpayService) ValidateSignature(rspraw, inputHash, secretKey string) bool {
// 	myChecksum := Util.HmacSHA512(secretKey, rspraw)
// 	return strings.EqualFold(myChecksum, inputHash)
// }
