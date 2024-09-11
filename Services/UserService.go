package Services

import (
	"th3y3m/e-commerce-platform/BusinessObjects"
	"th3y3m/e-commerce-platform/Repositories"
	"th3y3m/e-commerce-platform/Util"
	"time"

	"github.com/google/uuid"
)

func GetAllUsers() ([]BusinessObjects.User, error) {
	return Repositories.GetAllUsers()
}

func GetPaginatedUserList(searchValue, sortBy string, pageIndex, pageSize int, status *bool) (Util.PaginatedList[BusinessObjects.User], error) {
	return Repositories.GetPaginatedUserList(searchValue, sortBy, pageIndex, pageSize, status)
}

func GetUserByID(id string) (BusinessObjects.User, error) {
	return Repositories.GetUserByID(id)
}

func CreateUser(email, username, password, role string) error {
	passwordHash, err := Util.HashPassword(password)
	if err != nil {
		return err
	}

	user := BusinessObjects.User{
		UserID:       uuid.New().String(),
		Email:        email,
		Username:     username,
		PasswordHash: passwordHash,
		UserType:     role,
		CreatedAt:    time.Now(),
		Status:       true,
	}

	if err := Repositories.CreateUser(user); err != nil {
		return err
	}

	return nil
}

func BanUser(id string) error {
	user, err := GetUserByID(id)
	if err != nil {
		return err
	}

	user.Status = false
	if err := Repositories.UpdateUser(user); err != nil {
		return err
	}
	return nil
}
