package Interface

type IAuthenticationService interface {
	Login(email, password string) (string, error)
	RegisterCustomer(email, password string) error
	RegisterSeller(email, password string) error
	RegisterAdmin(email, password string) error
	VerifyUserEmail(token string) error
}
