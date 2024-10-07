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
