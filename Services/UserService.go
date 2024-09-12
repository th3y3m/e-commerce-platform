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

func CreateUser(email, password, role string) (BusinessObjects.User, error) {
	passwordHash, err := Util.HashPassword(password)
	if err != nil {
		return BusinessObjects.User{}, err
	}

	user := BusinessObjects.User{
		UserID:       uuid.New().String(),
		Email:        email,
		Username:     email,
		PasswordHash: passwordHash,
		UserType:     role,
		CreatedAt:    time.Now(),
		Status:       true,
	}

	newUser, err := Repositories.CreateUser(user)
	if err != nil {
		return BusinessObjects.User{}, err
	}

	return newUser, nil
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

func UnBanUser(id string) error {
	user, err := GetUserByID(id)
	if err != nil {
		return err
	}

	user.Status = true
	if err := Repositories.UpdateUser(user); err != nil {
		return err
	}
	return nil
}

func UpdateProfile(id, fullname, phonenumber, address string) error {
	user, err := GetUserByID(id)
	if err != nil {
		return err
	}

	user.FullName = fullname
	user.PhoneNumber = phonenumber
	user.Address = address

	if err := Repositories.UpdateUser(user); err != nil {
		return err
	}
	return nil
}
