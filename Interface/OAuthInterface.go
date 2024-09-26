package Interface

import "github.com/markbates/goth"

type IOAuthService interface {
	HandleOAuthUser(user goth.User) (string, error)
}
