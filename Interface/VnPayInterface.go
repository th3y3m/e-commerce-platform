package Interface

import (
	"net/url"
	"th3y3m/e-commerce-platform/BusinessObjects"
)

type IVnPayService interface {
	CreateVNPayUrl(amount float64, orderinfor string) (string, error)
	ValidateVNPayResponse(queryString url.Values) (*BusinessObjects.TransactionStatusModel, error)
	ValidateSignature(rspraw, inputHash, secretKey string) bool
}
