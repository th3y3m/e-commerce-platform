package Interface

import (
	"th3y3m/e-commerce-platform/BusinessObjects"
)

type IMailService interface {
	SendMail(to string, token string) error
	VerifyToken(token string) bool
	SendOrderDetails(Customer BusinessObjects.User,
		Order BusinessObjects.Order,
		OrderDetails []BusinessObjects.OrderDetail) error
}
