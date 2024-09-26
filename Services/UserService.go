package Services

import (
	"th3y3m/e-commerce-platform/BusinessObjects"
	"th3y3m/e-commerce-platform/Interface"
	"th3y3m/e-commerce-platform/Util"
	"time"

	"github.com/google/uuid"
	"github.com/sirupsen/logrus"
)

type UserService struct {
	userRepository Interface.IUserRepository
	log            *logrus.Logger
}

func NewUserService(userRepository Interface.IUserRepository, log *logrus.Logger) Interface.IUserService {
	return &UserService{
		userRepository: userRepository,
		log:            log,
	}
}

func (r *UserService) GetAllUsers() ([]BusinessObjects.User, error) {
	return r.userRepository.GetAllUsers()
}

func (r *UserService) GetPaginatedUserList(searchValue, sortBy string, pageIndex, pageSize int, status *bool) (Util.PaginatedList[BusinessObjects.User], error) {
	return r.userRepository.GetPaginatedUserList(searchValue, sortBy, pageIndex, pageSize, status)
}

func (r *UserService) GetUserByID(id string) (BusinessObjects.User, error) {
	return r.userRepository.GetUserByID(id)
}

func (r *UserService) CreateUser(email, password, role string) (BusinessObjects.User, error) {
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
		ImageURL:     `https://firebasestorage.googleapis.com/v0/b/sendo-a5204.appspot.com/o/users%2FOIP.jpeg?alt=media&token=438c9b2b-2ed0-4252-a3e1-24045ea76b7e`,
		CreatedAt:    time.Now(),
		Status:       true,
	}

	newUser, err := r.userRepository.CreateUser(user)
	if err != nil {
		return BusinessObjects.User{}, err
	}

	return newUser, nil
}

func (r *UserService) BanUser(id string) error {
	user, err := r.userRepository.GetUserByID(id)
	if err != nil {
		return err
	}

	user.Status = false
	if err := r.userRepository.UpdateUser(user); err != nil {
		return err
	}
	return nil
}

func (r *UserService) UnBanUser(id string) error {
	user, err := r.userRepository.GetUserByID(id)
	if err != nil {
		return err
	}

	user.Status = true
	if err := r.userRepository.UpdateUser(user); err != nil {
		return err
	}
	return nil
}

func (r *UserService) UpdateProfile(id, fullname, phonenumber, address, ImageURL string) error {
	user, err := r.userRepository.GetUserByID(id)
	if err != nil {
		return err
	}

	user.FullName = fullname
	user.PhoneNumber = phonenumber
	user.Address = address
	user.ImageURL = ImageURL

	if err := r.userRepository.UpdateUser(user); err != nil {
		return err
	}
	return nil
}
