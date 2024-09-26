package Interface

type IMailService interface {
	SendMail(to string, token string) error
	VerifyToken(token string) bool
}
