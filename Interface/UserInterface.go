package Interface

import (
	"th3y3m/e-commerce-platform/BusinessObjects"
	"th3y3m/e-commerce-platform/Util"
)

type IUserRepository interface {
	GetPaginatedUserList(searchValue, sortBy string, pageIndex, pageSize int, status *bool) (Util.PaginatedList[BusinessObjects.User], error)
	GetAllUsers() ([]BusinessObjects.User, error)
	GetUserByID(userID string) (BusinessObjects.User, error)
	CreateUser(user BusinessObjects.User) (BusinessObjects.User, error)
	UpdateUser(user BusinessObjects.User) error
	DeleteUser(userID string) error
	GetUserByEmail(email string) (BusinessObjects.User, error)
	StoreToken(user *BusinessObjects.User, token string) error
	SetToken(user *BusinessObjects.User, token string) error
	VerifyToken(token string) bool
	GetUserByToken(token string) (BusinessObjects.User, error)
}

type IUserService interface {
	GetAllUsers() ([]BusinessObjects.User, error)
	GetPaginatedUserList(searchValue, sortBy string, pageIndex, pageSize int, status *bool) (Util.PaginatedList[BusinessObjects.User], error)
	GetUserByID(userID string) (BusinessObjects.User, error)
	CreateUser(email, password, role string) (BusinessObjects.User, error)
	BanUser(id string) error
	UnBanUser(id string) error
	UpdateProfile(id, fullname, phonenumber, address, ImageURL string) error
}
