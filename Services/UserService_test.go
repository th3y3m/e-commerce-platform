package Services

import (
	"testing"
	"th3y3m/e-commerce-platform/BusinessObjects"
	"th3y3m/e-commerce-platform/Util"
	"th3y3m/e-commerce-platform/mocks"

	"github.com/google/uuid"
	"github.com/sirupsen/logrus"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestGetById_Success(t *testing.T) {
	// Setup
	mockRepo := &mocks.IUserRepository{}
	log := logrus.New()
	userService := UserService{
		userRepository: mockRepo,
		log:            log,
	}

	userID := "1"
	expectedUser := BusinessObjects.User{
		UserID: "1",
		Email:  "test@gmail.com",
	}
	mockRepo.On("GetUserByID", userID).Return(expectedUser, nil)

	// Act
	result, err := userService.GetUserByID(userID)

	// Assert
	assert.NoError(t, err)
	assert.NotNil(t, result)
	assert.Equal(t, expectedUser, result)
	mockRepo.AssertExpectations(t)
}

func TestGetById_Error(t *testing.T) {
	// Setup
	mockRepo := &mocks.IUserRepository{}
	log := logrus.New()
	userService := UserService{
		userRepository: mockRepo,
		log:            log,
	}

	userID := "1"
	expectedError := assert.AnError
	mockRepo.On("GetUserByID", userID).Return(BusinessObjects.User{}, expectedError)

	// Act
	result, err := userService.GetUserByID(userID)

	// Assert
	assert.Error(t, err)
	assert.Equal(t, expectedError, err)
	assert.Equal(t, BusinessObjects.User{}, result)
	mockRepo.AssertExpectations(t)
}

func TestCreateUser_Success(t *testing.T) {
	// Setup
	mockRepo := &mocks.IUserRepository{}
	log := logrus.New()
	userService := UserService{
		userRepository: mockRepo,
		log:            log,
	}

	email := "test@gmail.com"
	password := "password"
	role := "user"
	passwordHash, _ := Util.HashPassword(password) // Hash the password

	expectedUser := BusinessObjects.User{
		UserID:       uuid.New().String(),
		Email:        email,
		PasswordHash: passwordHash,
		UserType:     role,
	}

	mockRepo.On("CreateUser", mock.AnythingOfType("BusinessObjects.User")).Return(expectedUser, nil)
	// Act
	result, err := userService.CreateUser(email, password, role)

	// Assert
	assert.NoError(t, err)
	assert.NotNil(t, result)
	assert.Equal(t, expectedUser.Email, result.Email)
	mockRepo.AssertExpectations(t)
}

func TestCreateUser_Error(t *testing.T) {
	// Setup
	mockRepo := &mocks.IUserRepository{}
	log := logrus.New()
	userService := UserService{
		userRepository: mockRepo,
		log:            log,
	}

	email := "test@gmail.com"
	password := "password"
	role := "user"
	expectedError := assert.AnError

	mockRepo.On("CreateUser", mock.AnythingOfType("BusinessObjects.User")).Return(BusinessObjects.User{}, expectedError)
	// Act
	result, err := userService.CreateUser(email, password, role)

	// Assert
	assert.Error(t, err)
	assert.Equal(t, expectedError, err)
	assert.Equal(t, BusinessObjects.User{}, result)
	mockRepo.AssertExpectations(t)
}

func TestBanUser_Success(t *testing.T) {
	// Setup
	mockRepo := &mocks.IUserRepository{}
	log := logrus.New()
	userService := UserService{
		userRepository: mockRepo,
		log:            log,
	}

	userID := "1"
	user := BusinessObjects.User{
		UserID: "1",
		Email:  "test@gmail.com",
		Status: true,
	}
	expectedUser := user
	expectedUser.Status = false

	mockRepo.On("GetUserByID", userID).Return(user, nil)
	mockRepo.On("UpdateUser", expectedUser).Return(nil)

	// Act
	err := userService.BanUser(userID)

	// Assert
	assert.NoError(t, err)
	mockRepo.AssertExpectations(t)
}

func TestBanUser_Error(t *testing.T) {
	// Setup
	mockRepo := &mocks.IUserRepository{}
	log := logrus.New()
	userService := UserService{
		userRepository: mockRepo,
		log:            log,
	}

	userID := "1"
	user := BusinessObjects.User{
		UserID: "1",
		Email:  "test@gmail.com",
		Status: true,
	}
	expectedError := assert.AnError

	mockRepo.On("GetUserByID", userID).Return(user, nil)
	mockRepo.On("UpdateUser", mock.AnythingOfType("BusinessObjects.User")).Return(expectedError)

	// Act
	err := userService.BanUser(userID)

	// Assert
	assert.Error(t, err)
	assert.Equal(t, expectedError, err)
	mockRepo.AssertExpectations(t)
}

func TestGetAllUsers_Success(t *testing.T) {
	// Setup
	mockRepo := &mocks.IUserRepository{}
	log := logrus.New()
	userService := UserService{
		userRepository: mockRepo,
		log:            log,
	}

	expectedUsers := []BusinessObjects.User{
		{
			UserID: "1",
			Email:  "test@gmail.com",
		},
		{
			UserID: "2",
			Email:  "test2@gmail.com",
		},
	}

	mockRepo.On("GetAllUsers").Return(expectedUsers, nil)

	// Act
	result, err := userService.GetAllUsers()

	// Assert
	assert.NoError(t, err)
	assert.NotNil(t, result)
	assert.Equal(t, expectedUsers, result)
	mockRepo.AssertExpectations(t)
}

func TestGetAllUsers_Error(t *testing.T) {
	// Setup
	mockRepo := &mocks.IUserRepository{}
	log := logrus.New()
	userService := UserService{
		userRepository: mockRepo,
		log:            log,
	}

	expectedError := assert.AnError

	mockRepo.On("GetAllUsers").Return([]BusinessObjects.User{}, expectedError)

	// Act
	result, err := userService.GetAllUsers()

	// Assert
	assert.Error(t, err)
	assert.Equal(t, expectedError, err)
	assert.Equal(t, []BusinessObjects.User{}, result)
	mockRepo.AssertExpectations(t)
}

func TestGetPaginatedUserList_Success(t *testing.T) {
	// Setup
	mockRepo := &mocks.IUserRepository{}
	log := logrus.New()
	userService := UserService{
		userRepository: mockRepo,
		log:            log,
	}

	searchValue := "test"
	sortBy := "email"
	pageIndex := 1
	pageSize := 10
	status := true

	expectedUsers := Util.PaginatedList[BusinessObjects.User]{
		Items: []BusinessObjects.User{
			{
				UserID: "1",
				Email:  "test@gmail.com",
			},
			{
				UserID: "2",
				Email:  "test2@gmail.com",
			},
		},
		TotalCount: 2,
	}

	mockRepo.On("GetPaginatedUserList", searchValue, sortBy, pageIndex, pageSize, &status).Return(expectedUsers, nil)

	// Act
	result, err := userService.GetPaginatedUserList(searchValue, sortBy, pageIndex, pageSize, &status)

	// Assert
	assert.NoError(t, err)
	assert.NotNil(t, result)
	assert.Equal(t, expectedUsers, result)
	mockRepo.AssertExpectations(t)
}

func TestGetPaginatedUserList_Error(t *testing.T) {
	// Setup
	mockRepo := &mocks.IUserRepository{}
	log := logrus.New()
	userService := UserService{
		userRepository: mockRepo,
		log:            log,
	}

	searchValue := "test"
	sortBy := "email"
	pageIndex := 1
	pageSize := 10
	status := true

	expectedError := assert.AnError

	mockRepo.On("GetPaginatedUserList", searchValue, sortBy, pageIndex, pageSize, &status).Return(Util.PaginatedList[BusinessObjects.User]{}, expectedError)

	// Act
	result, err := userService.GetPaginatedUserList(searchValue, sortBy, pageIndex, pageSize, &status)

	// Assert
	assert.Error(t, err)
	assert.Equal(t, expectedError, err)
	assert.Equal(t, Util.PaginatedList[BusinessObjects.User]{}, result)
	mockRepo.AssertExpectations(t)
}

func TestUpdateProfile_Success(t *testing.T) {
	// Setup
	mockRepo := &mocks.IUserRepository{}
	log := logrus.New()
	userService := UserService{
		userRepository: mockRepo,
		log:            log,
	}

	userID := "1"
	user := BusinessObjects.User{
		UserID:      "1",
		Email:       "test@gmail.com",
		Username:    "test",
		FullName:    "test",
		PhoneNumber: "123456789",
		Address:     "test",
		ImageURL:    "test",
	}

	mockRepo.On("GetUserByID", userID).Return(user, nil)
	mockRepo.On("UpdateUser", user).Return(nil)

	// Act
	err := userService.UpdateProfile(userID, user.FullName, user.PhoneNumber, user.Address, user.ImageURL)

	// Assert
	assert.NoError(t, err)
	mockRepo.AssertExpectations(t)
}

func TestUpdateProfile_Error(t *testing.T) {
	// Setup
	mockRepo := &mocks.IUserRepository{}
	log := logrus.New()
	userService := UserService{
		userRepository: mockRepo,
		log:            log,
	}

	userID := "1"
	user := BusinessObjects.User{
		UserID:      "1",
		Email:       "test@gmail.com",
		Username:    "test",
		FullName:    "test",
		PhoneNumber: "123456789",
		Address:     "test",
		ImageURL:    "test",
	}
	expectedError := assert.AnError

	mockRepo.On("GetUserByID", userID).Return(user, nil)
	mockRepo.On("UpdateUser", user).Return(expectedError)

	// Act
	err := userService.UpdateProfile(userID, user.FullName, user.PhoneNumber, user.Address, user.ImageURL)

	// Assert
	assert.Error(t, err)
	assert.Equal(t, expectedError, err)
	mockRepo.AssertExpectations(t)
}

func TestUnBanUser_Success(t *testing.T) {
	// Setup
	mockRepo := &mocks.IUserRepository{}
	log := logrus.New()
	userService := UserService{
		userRepository: mockRepo,
		log:            log,
	}

	userID := "1"
	user := BusinessObjects.User{
		UserID: "1",
		Email:  "test@gmail.com",
		Status: false,
	}

	expectedUser := user
	expectedUser.Status = true

	mockRepo.On("GetUserByID", userID).Return(user, nil)
	mockRepo.On("UpdateUser", expectedUser).Return(nil)

	// Act
	err := userService.UnBanUser(userID)

	// Assert
	assert.NoError(t, err)
	mockRepo.AssertExpectations(t)
}

func TestUnBanUser_Error(t *testing.T) {
	// Setup
	mockRepo := &mocks.IUserRepository{}
	log := logrus.New()
	userService := UserService{
		userRepository: mockRepo,
		log:            log,
	}

	userID := "1"
	user := BusinessObjects.User{
		UserID: "1",
		Email:  "test@gmail.com",
		Status: false,
	}
	expectedError := assert.AnError

	mockRepo.On("GetUserByID", userID).Return(user, nil)
	mockRepo.On("UpdateUser", mock.AnythingOfType("BusinessObjects.User")).Return(expectedError)

	// Act
	err := userService.UnBanUser(userID)

	// Assert
	assert.Error(t, err)
	assert.Equal(t, expectedError, err)
	mockRepo.AssertExpectations(t)
}
