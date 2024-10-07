package Services_test

import (
	"testing"
	"th3y3m/e-commerce-platform/BusinessObjects"
	"th3y3m/e-commerce-platform/Services"
	"th3y3m/e-commerce-platform/Util"
	"th3y3m/e-commerce-platform/mocks"
	"time"

	"github.com/go-co-op/gocron"
	"github.com/google/uuid"
	"github.com/sirupsen/logrus"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestLogin_Success(t *testing.T) {
	// Arrange
	userRepository := &mocks.IUserRepository{}
	userService := &mocks.IUserService{}
	mailService := &mocks.IMailService{}
	log := logrus.New()
	scheduler := gocron.NewScheduler(time.UTC)
	deleteJobs := make(map[string]*gocron.Job)
	authService := Services.NewAuthenticationService(userRepository, userService, log, scheduler, deleteJobs, mailService)

	email := "test@gmail.com"
	password := "password"

	passwordHash, err := Util.HashPassword(password)
	if err != nil {
		t.Error(err)
	}

	user := BusinessObjects.User{
		Email:        email,
		PasswordHash: passwordHash,
		UserID:       uuid.New().String(),
		UserType:     "user",
	}

	userRepository.On("GetUserByEmail", email).Return(user, nil)

	// Act
	token, err := authService.Login(email, password)

	// Assert
	assert.NoError(t, err)
	assert.NotEmpty(t, token)
	userRepository.AssertExpectations(t)
}

func TestLogin_Fail_EmptyEmail(t *testing.T) {
	// Arrange
	userRepository := &mocks.IUserRepository{}
	userService := &mocks.IUserService{}
	mailService := &mocks.IMailService{}
	log := logrus.New()
	scheduler := gocron.NewScheduler(time.UTC)
	deleteJobs := make(map[string]*gocron.Job)
	authService := Services.NewAuthenticationService(userRepository, userService, log, scheduler, deleteJobs, mailService)

	email := ""
	password := "password"

	// Act
	token, err := authService.Login(email, password)

	// Assert
	assert.Error(t, err)
	assert.Empty(t, token)
}

func TestLogin_Fail_EmptyPassword(t *testing.T) {
	// Arrange
	userRepository := &mocks.IUserRepository{}
	userService := &mocks.IUserService{}
	mailService := &mocks.IMailService{}
	log := logrus.New()
	scheduler := gocron.NewScheduler(time.UTC)
	deleteJobs := make(map[string]*gocron.Job)
	authService := Services.NewAuthenticationService(userRepository, userService, log, scheduler, deleteJobs, mailService)

	email := "test@gmail.com"
	password := ""

	// Act
	token, err := authService.Login(email, password)

	// Assert
	assert.Error(t, err)
	assert.Empty(t, token)
}

func TestLogin_Fail_UserNotFound(t *testing.T) {
	// Arrange
	userRepository := &mocks.IUserRepository{}
	userService := &mocks.IUserService{}
	mailService := &mocks.IMailService{}
	log := logrus.New()
	scheduler := gocron.NewScheduler(time.UTC)
	deleteJobs := make(map[string]*gocron.Job)
	authService := Services.NewAuthenticationService(userRepository, userService, log, scheduler, deleteJobs, mailService)

	email := "test@gmail.com"
	password := "password"

	userRepository.On("GetUserByEmail", email).Return(BusinessObjects.User{}, nil)

	// Act
	token, err := authService.Login(email, password)

	// Assert
	assert.Error(t, err)
	assert.Empty(t, token)
	userRepository.AssertExpectations(t)
}

func TestLogin_Fail_InvalidPassword(t *testing.T) {
	// Arrange
	userRepository := &mocks.IUserRepository{}
	userService := &mocks.IUserService{}
	mailService := &mocks.IMailService{}
	log := logrus.New()
	scheduler := gocron.NewScheduler(time.UTC)
	deleteJobs := make(map[string]*gocron.Job)
	authService := Services.NewAuthenticationService(userRepository, userService, log, scheduler, deleteJobs, mailService)

	email := "test@gmail.com"
	password := "password"

	passwordHash, err := Util.HashPassword("wrongpassword")
	if err != nil {
		t.Error(err)
	}

	user := BusinessObjects.User{
		Email:        email,
		PasswordHash: passwordHash,
		UserID:       uuid.New().String(),
		UserType:     "user",
	}

	userRepository.On("GetUserByEmail", email).Return(user, nil)

	// Act
	token, err := authService.Login(email, password)

	// Assert
	assert.Error(t, err)
	assert.Empty(t, token)
	userRepository.AssertExpectations(t)
}

func TestRegister_Success(t *testing.T) {
	// Arrange
	userRepository := &mocks.IUserRepository{}
	userService := &mocks.IUserService{}
	mailService := &mocks.IMailService{}
	log := logrus.New()
	scheduler := gocron.NewScheduler(time.UTC)
	deleteJobs := make(map[string]*gocron.Job)
	authService := Services.NewAuthenticationService(userRepository, userService, log, scheduler, deleteJobs, mailService)

	email := "test@gmail.com"
	password := "password"

	user := BusinessObjects.User{
		Email:        email,
		PasswordHash: "",
		UserID:       uuid.New().String(),
		UserType:     "customer",
	}

	userRepository.On("GetUserByEmail", email).Return(BusinessObjects.User{}, nil)
	userService.On("CreateUser", email, password, "customer").Return(user, nil)
	userRepository.On("StoreToken", &user, mock.AnythingOfType("string")).Return(nil)
	mailService.On("SendMail", user.Email, mock.AnythingOfType("string")).Return(nil)

	// Act
	err := authService.RegisterCustomer(email, password)

	// Assert
	assert.NoError(t, err)
	userRepository.AssertExpectations(t)
	userService.AssertExpectations(t)
	mailService.AssertExpectations(t)
}

func TestRegister_Fail_EmptyEmail(t *testing.T) {
	// Arrange
	userRepository := &mocks.IUserRepository{}
	userService := &mocks.IUserService{}
	mailService := &mocks.IMailService{}
	log := logrus.New()
	scheduler := gocron.NewScheduler(time.UTC)
	deleteJobs := make(map[string]*gocron.Job)
	authService := Services.NewAuthenticationService(userRepository, userService, log, scheduler, deleteJobs, mailService)

	email := ""
	password := "password"

	// Act
	err := authService.RegisterCustomer(email, password)

	// Assert
	assert.Error(t, err)
}

func TestRegister_Fail_EmptyPassword(t *testing.T) {
	// Arrange
	userRepository := &mocks.IUserRepository{}
	userService := &mocks.IUserService{}
	mailService := &mocks.IMailService{}
	log := logrus.New()
	scheduler := gocron.NewScheduler(time.UTC)
	deleteJobs := make(map[string]*gocron.Job)
	authService := Services.NewAuthenticationService(userRepository, userService, log, scheduler, deleteJobs, mailService)

	email := "test@gmail.com"
	password := ""

	// Act
	err := authService.RegisterCustomer(email, password)

	// Assert
	assert.Error(t, err)
}

func TestRegister_Fail_UserAlreadyExists(t *testing.T) {
	// Arrange
	userRepository := &mocks.IUserRepository{}
	userService := &mocks.IUserService{}
	mailService := &mocks.IMailService{}
	log := logrus.New()
	scheduler := gocron.NewScheduler(time.UTC)
	deleteJobs := make(map[string]*gocron.Job)
	authService := Services.NewAuthenticationService(userRepository, userService, log, scheduler, deleteJobs, mailService)

	email := "test@gmail.com"
	password := "password"

	user := BusinessObjects.User{
		Email:        email,
		PasswordHash: "",
		UserID:       uuid.New().String(),
		UserType:     "customer",
	}

	userRepository.On("GetUserByEmail", email).Return(user, nil)

	// Act
	err := authService.RegisterCustomer(email, password)

	// Assert
	assert.Error(t, err)
}
