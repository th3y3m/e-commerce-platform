package Repositories_test

import (
	"testing"
	"th3y3m/e-commerce-platform/BusinessObjects"
	"th3y3m/e-commerce-platform/Repositories"

	"github.com/sirupsen/logrus"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"gorm.io/gorm"
)

// MockDB implements the Database interface using testify/mock
type MockDB struct {
	mock.Mock
}

// Mock Find method
func (m *MockDB) Find(dest interface{}, conds ...interface{}) *gorm.DB {
	// Correctly call m.Called by spreading conds
	args := m.Called(append([]interface{}{dest}, conds...)...)

	// Simulate populating `dest` with the mock data
	// if users, ok := args.Get(0).([]BusinessObjects.User); ok {
	// 	*dest.(*[]BusinessObjects.User) = users
	// }

	// return &gorm.DB{}
	return args.Get(0).(*gorm.DB)
}

// Mock First method
func (m *MockDB) First(dest interface{}, conds ...interface{}) *gorm.DB {
	args := m.Called(dest, conds)
	return args.Get(0).(*gorm.DB)
}

// Mock Create method
func (m *MockDB) Create(value interface{}) *gorm.DB {
	args := m.Called(value)
	return args.Get(0).(*gorm.DB)
}

// Mock Save method
func (m *MockDB) Save(value interface{}) *gorm.DB {
	args := m.Called(value)
	return args.Get(0).(*gorm.DB)
}

// Mock Delete method
func (m *MockDB) Delete(value interface{}, conds ...interface{}) *gorm.DB {
	args := m.Called(append([]interface{}{value}, conds...)...)
	return args.Get(0).(*gorm.DB)
}

// Mock Where method
func (m *MockDB) Where(query interface{}, args ...interface{}) *gorm.DB {
	args1 := m.Called(query, args)
	return args1.Get(0).(*gorm.DB)
}

// Mock Count method
func (m *MockDB) Count(count *int64) *gorm.DB {
	args := m.Called(count)
	return args.Get(0).(*gorm.DB)
}

func (m *MockDB) Model(value interface{}) *gorm.DB {
	args := m.Called(value)
	return args.Get(0).(*gorm.DB)
}

func TestGetAllUsers(t *testing.T) {
	// Set up mock logger and DB
	log := logrus.New()
	mockDB := new(MockDB)

	// Create mock return value
	mockUsers := []BusinessObjects.User{
		{UserID: "1", Email: "test1@example.com"},
		{UserID: "2", Email: "test2@example.com"},
	}

	// Mock the Find function to simulate finding users
	mockDB.On("Find", mock.Anything, mock.Anything).Return(&gorm.DB{}).Run(func(args mock.Arguments) {
		// Simulate GORM populating the users
		arg := args.Get(0).(*[]BusinessObjects.User)
		*arg = mockUsers
	})

	// Create UserRepository with mock DB
	repo := Repositories.NewUserRepository(log, mockDB)

	// Call GetAllUsers
	users, err := repo.GetAllUsers()

	// Assertions
	assert.NoError(t, err)
	assert.Equal(t, 2, len(users))
	assert.Equal(t, "test1@example.com", users[0].Email)
	assert.Equal(t, "test2@example.com", users[1].Email)

	// Verify that all expectations were met
	mockDB.AssertExpectations(t)
}

func TestGetUserByID(t *testing.T) {
	// Set up mock logger and DB
	log := logrus.New()
	mockDB := new(MockDB)

	// Create mock return value for the specific user
	mockUser := BusinessObjects.User{
		UserID: "1",
		Email:  "test1@example.com",
	}

	// Mock the First function to simulate finding a user by ID
	// We expect the query as a slice []interface{}{"user_id = ?", "1"}
	mockDB.On("First", mock.Anything, []interface{}{"user_id = ?", "1"}).Return(&gorm.DB{}).Run(func(args mock.Arguments) {
		// Simulate GORM populating the user
		arg := args.Get(0).(*BusinessObjects.User)
		*arg = mockUser
	})

	// Create UserRepository with mock DB
	repo := Repositories.NewUserRepository(log, mockDB)

	// Call GetUserByID
	user, err := repo.GetUserByID("1")

	// Assertions
	assert.NoError(t, err)
	assert.Equal(t, "1", user.UserID)
	assert.Equal(t, "test1@example.com", user.Email)

	// Verify that all expectations were met
	mockDB.AssertExpectations(t)
}

func TestGetUserByIDUsingDDT(t *testing.T) {
	// Set up mock logger and DB
	log := logrus.New()
	mockDB := new(MockDB)

	// Define test cases in a table-driven format
	testCases := []struct {
		name       string               // Name of the test case
		userID     string               // Input: user ID
		mockUser   BusinessObjects.User // The user we expect to be returned
		shouldFail bool                 // Whether the test should fail (if user not found)
	}{
		{
			name:   "User Found - ID 1",
			userID: "1",
			mockUser: BusinessObjects.User{
				UserID: "1",
				Email:  "test1@example.com",
			},
			shouldFail: false,
		},
		{
			name:   "User Found - ID 2",
			userID: "2",
			mockUser: BusinessObjects.User{
				UserID: "2",
				Email:  "test2@example.com",
			},
			shouldFail: false,
		},
		{
			name:       "User Not Found",
			userID:     "999",
			mockUser:   BusinessObjects.User{},
			shouldFail: true, // Expect an error when the user isn't found
		},
	}

	// Set up mock responses based on the test cases
	for _, tc := range testCases {
		if tc.shouldFail {
			// For the test case where we expect a failure, mock an error response
			mockDB.On("First", mock.Anything, []interface{}{"user_id = ?", tc.userID}).Return(&gorm.DB{Error: gorm.ErrRecordNotFound})
		} else {
			// For test cases where a user is found, mock the user response
			mockDB.On("First", mock.Anything, []interface{}{"user_id = ?", tc.userID}).Return(&gorm.DB{}).Run(func(args mock.Arguments) {
				arg := args.Get(0).(*BusinessObjects.User)
				*arg = tc.mockUser
			})
		}
	}

	// Iterate over test cases
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			// Create UserRepository with mock DB
			repo := Repositories.NewUserRepository(log, mockDB)

			// Call GetUserByID
			user, err := repo.GetUserByID(tc.userID)

			// Assertions based on whether an error was expected
			if tc.shouldFail {
				assert.Error(t, err)
				assert.Equal(t, BusinessObjects.User{}, user) // Ensure no user is returned
			} else {
				assert.NoError(t, err)
				assert.Equal(t, tc.mockUser.UserID, user.UserID)
				assert.Equal(t, tc.mockUser.Email, user.Email)
			}
		})
	}

	// Verify that all expectations were met
	mockDB.AssertExpectations(t)
}

func TestCreateUser(t *testing.T) {
	// Set up mock logger and DB
	log := logrus.New()
	mockDB := new(MockDB)

	// Define test cases
	testCases := []struct {
		name       string               // Name of the test case
		user       BusinessObjects.User // Input: user to create
		shouldFail bool                 // Whether the test should fail
	}{
		{
			name: "Create User Success",
			user: BusinessObjects.User{
				UserID: "1",
				Email:  "test1@example.com",
			},
			shouldFail: false,
		},
		{
			name: "Create User Failure",
			user: BusinessObjects.User{
				UserID: "2",
				Email:  "test2@example.com",
			},
			shouldFail: true, // Expect failure
		},
	}

	// Set up mock responses based on the test cases
	for _, tc := range testCases {
		if tc.shouldFail {
			mockDB.On("Create", &tc.user).Return(&gorm.DB{Error: gorm.ErrInvalidData})
		} else {
			mockDB.On("Create", &tc.user).Return(&gorm.DB{})
		}
	}

	// Iterate over test cases
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			// Create UserRepository with mock DB
			repo := Repositories.NewUserRepository(log, mockDB)

			// Call CreateUser
			user, err := repo.CreateUser(tc.user)

			// Assertions based on whether an error was expected
			if tc.shouldFail {
				assert.Error(t, err)
				assert.Equal(t, BusinessObjects.User{}, user)
			} else {
				assert.NoError(t, err)
				assert.Equal(t, tc.user.Email, user.Email)
			}
		})
	}

	// Verify that all expectations were met
	mockDB.AssertExpectations(t)
}

func TestUpdateUser(t *testing.T) {
	// Set up mock logger and DB
	log := logrus.New()
	mockDB := new(MockDB)

	// Define test cases
	testCases := []struct {
		name       string               // Name of the test case
		user       BusinessObjects.User // Input: user to update
		shouldFail bool                 // Whether the test should fail
	}{
		{
			name: "Update User Success",
			user: BusinessObjects.User{
				UserID: "1",
				Email:  "test1@example.com",
			},
			shouldFail: false,
		},
		{
			name: "Update User Failure",
			user: BusinessObjects.User{
				UserID: "2",
				Email:  "test2@example.com",
			},
			shouldFail: true, // Expect failure
		},
	}

	// Set up mock responses based on the test cases
	for _, tc := range testCases {
		if tc.shouldFail {
			mockDB.On("Save", &tc.user).Return(&gorm.DB{Error: gorm.ErrInvalidData})
		} else {
			mockDB.On("Save", &tc.user).Return(&gorm.DB{})
		}
	}

	// Iterate over test cases
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			// Create UserRepository with mock DB
			repo := Repositories.NewUserRepository(log, mockDB)

			// Call UpdateUser
			err := repo.UpdateUser(tc.user)

			// Assertions based on whether an error was expected
			if tc.shouldFail {
				assert.Error(t, err)
			} else {
				assert.NoError(t, err)
			}
		})
	}

	// Verify that all expectations were met
	mockDB.AssertExpectations(t)
}

func TestDeleteUser(t *testing.T) {
	// Set up mock logger and DB
	log := logrus.New()
	mockDB := new(MockDB)

	// Define test cases
	testCases := []struct {
		userID string
	}{
		{userID: "1"},
		{userID: "2"},
	}

	for _, tc := range testCases {
		// Set up mock expectations
		mockDB.On("Delete", &BusinessObjects.User{}, "user_id = ?", tc.userID).Return(&gorm.DB{})

		// Create UserRepository with mock DB
		repo := Repositories.NewUserRepository(log, mockDB)

		// Call DeleteUser
		err := repo.DeleteUser(tc.userID)

		// Assertions
		assert.NoError(t, err)
		mockDB.AssertExpectations(t)
	}
}

func TestGetUserByEmail(t *testing.T) {
	MockDB := new(MockDB)

	user := BusinessObjects.User{
		UserID: "1",
		Email:  "test1@gmail.com",
	}

	MockDB.On("First", mock.Anything, mock.Anything).Return(&gorm.DB{}).Run(func(args mock.Arguments) {
		arg := args.Get(0).(*BusinessObjects.User)
		*arg = user
	})

	repo := Repositories.NewUserRepository(logrus.New(), MockDB)

	user, err := repo.GetUserByEmail(user.Email)

	assert.NoError(t, err)
	assert.Equal(t, "test1@gmail.com", user.Email)
	assert.Equal(t, "1", user.UserID)
}

func TestStoreToken(t *testing.T) {
	MockDB := new(MockDB)

	user := BusinessObjects.User{
		UserID: "1",
	}

	MockDB.On("Save", &user).Return(&gorm.DB{})
	repo := Repositories.NewUserRepository(logrus.New(), MockDB)

	err := repo.StoreToken(&user, "token")

	assert.NoError(t, err)

	assert.Equal(t, "token", user.Token)
}

func TestSetToken(t *testing.T) {
	MockDB := new(MockDB)

	user := BusinessObjects.User{
		UserID: "1",
		Token:  "123",
	}

	MockDB.On("Save", &user).Return(&gorm.DB{})
	repo := Repositories.NewUserRepository(logrus.New(), MockDB)

	err := repo.SetToken(&user, "token")

	assert.NoError(t, err)

	assert.Equal(t, "token", user.Token)
}

func TestGetUserByToken(t *testing.T) {
	MockDB := new(MockDB)

	user := BusinessObjects.User{
		UserID: "1",
		Token:  "token",
	}

	MockDB.On("First", mock.Anything, mock.Anything).Return(&gorm.DB{}).Run(func(args mock.Arguments) {
		arg := args.Get(0).(*BusinessObjects.User)
		*arg = user
	})

	repo := Repositories.NewUserRepository(logrus.New(), MockDB)

	user, err := repo.GetUserByToken("token")

	assert.NoError(t, err)
	assert.Equal(t, "1", user.UserID)
	assert.Equal(t, "token", user.Token)
}
