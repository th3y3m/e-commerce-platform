package Repositories

import (
	"errors"
	"fmt"
	"strings"
	"th3y3m/e-commerce-platform/BusinessObjects"
	"th3y3m/e-commerce-platform/Interface"
	"th3y3m/e-commerce-platform/Util"
	"time"

	"github.com/sirupsen/logrus"
	"gorm.io/gorm"
)

type UserRepository struct {
	log *logrus.Logger
	db  *gorm.DB
}

func NewUserRepository(log *logrus.Logger, db *gorm.DB) Interface.IUserRepository {
	return &UserRepository{log: log, db: db}
}

func (r *UserRepository) GetPaginatedUserList(searchValue, sortBy string, pageIndex, pageSize int, status *bool) (Util.PaginatedList[BusinessObjects.User], error) {
	r.log.Infof("Fetching paginated user list with searchValue: %s, sortBy: %s, pageIndex: %d, pageSize: %d, status: %v", searchValue, sortBy, pageIndex, pageSize, status)
	// db, err := Util.ConnectToPostgreSQL()
	// if err != nil {
	// 	r.log.Error("Failed to connect to PostgreSQL:", err)
	// 	return Util.PaginatedList[BusinessObjects.User]{}, err
	// }

	var users []BusinessObjects.User

	// Scope the query to the User table
	query := r.db.Model(&BusinessObjects.User{})

	// Apply search filter
	if searchValue != "" {
		searchValueLower := strings.ToLower(searchValue)
		query = query.Where("LOWER(email) LIKE ?", "%"+searchValueLower+"%")
	}

	// Apply status filter only if status is provided
	if status != nil {
		query = query.Where("status = ?", *status)
	}

	// Apply sorting
	switch sortBy {
	case "email_asc":
		query = query.Order("email ASC")
	case "email_desc":
		query = query.Order("email DESC")
	case "username_asc":
		query = query.Order("username ASC")
	case "username_desc":
		query = query.Order("username DESC")
	case "fullname_asc":
		query = query.Order("full_name ASC")
	case "fullname_desc":
		query = query.Order("full_name DESC")
	case "phonenumber_asc":
		query = query.Order("phone_number ASC")
	case "phonenumber_desc":
		query = query.Order("phone_number DESC")
	case "address_asc":
		query = query.Order("address ASC")
	case "address_desc":
		query = query.Order("address DESC")
	case "usertype_asc":
		query = query.Order("user_type ASC")
	case "usertype_desc":
		query = query.Order("user_type DESC")
	case "createdat_asc":
		query = query.Order("created_at ASC")
	case "createdat_desc":
		query = query.Order("created_at DESC")
	case "status_asc":
		query = query.Order("status ASC")
	case "status_desc":
		query = query.Order("status DESC")
	default:
		query = query.Order("created_at DESC") // Default sorting
	}

	// Get total count of filtered users
	var totalCount int64
	if err := query.Count(&totalCount).Error; err != nil {
		r.log.Error("Failed to count users:", err)
		return Util.PaginatedList[BusinessObjects.User]{}, err
	}

	// Apply pagination
	offset := (pageIndex - 1) * pageSize
	if err := query.Offset(offset).Limit(pageSize).Find(&users).Error; err != nil {
		r.log.Error("Failed to fetch paginated users:", err)
		return Util.PaginatedList[BusinessObjects.User]{}, err
	}

	// Create paginated list
	paginatedList := Util.NewPaginatedList(users, totalCount, pageIndex, pageSize)

	r.log.Infof("Successfully fetched paginated user list with total count: %d", totalCount)
	return paginatedList, nil
}

func (r *UserRepository) GetAllUsers() ([]BusinessObjects.User, error) {
	r.log.Info("Fetching all users")
	// db, err := Util.ConnectToPostgreSQL()
	// if err != nil {
	// 	r.log.Error("Failed to connect to PostgreSQL:", err)
	// 	return nil, err
	// }

	var users []BusinessObjects.User
	if err := r.db.Find(&users).Error; err != nil {
		r.log.Error("Failed to fetch all users:", err)
		return nil, err
	}

	r.log.Info("Successfully fetched all users")
	return users, nil
}

func (r *UserRepository) GetUserByID(userID string) (BusinessObjects.User, error) {
	r.log.Infof("Fetching user by ID: %s", userID)
	// db, err := Util.ConnectToPostgreSQL()
	// if err != nil {
	// 	r.log.Error("Failed to connect to PostgreSQL:", err)
	// 	return BusinessObjects.User{}, err
	// }

	var user BusinessObjects.User
	if err := r.db.First(&user, "user_id = ?", userID).Error; err != nil {
		r.log.Error("Failed to fetch user by ID:", err)
		return BusinessObjects.User{}, err
	}

	r.log.Infof("Successfully fetched user by ID: %s", userID)
	return user, nil
}

func (r *UserRepository) CreateUser(user BusinessObjects.User) (BusinessObjects.User, error) {
	r.log.Infof("Creating new user with email: %s", user.Email)
	// db, err := Util.ConnectToPostgreSQL()
	// if err != nil {
	// 	r.log.Error("Failed to connect to PostgreSQL:", err)
	// 	return BusinessObjects.User{}, err
	// }

	if err := r.db.Create(&user).Error; err != nil {
		r.log.Error("Failed to create new user:", err)
		return BusinessObjects.User{}, err
	}

	r.log.Infof("Successfully created new user with email: %s", user.Email)
	return user, nil
}

func (r *UserRepository) UpdateUser(user BusinessObjects.User) error {
	r.log.Infof("Updating user with ID: %s", user.UserID)
	// db, err := Util.ConnectToPostgreSQL()
	// if err != nil {
	// 	r.log.Error("Failed to connect to PostgreSQL:", err)
	// 	return err
	// }

	if err := r.db.Save(&user).Error; err != nil {
		r.log.Error("Failed to update user:", err)
		return err
	}

	r.log.Infof("Successfully updated user with ID: %s", user.UserID)
	return nil
}

func (r *UserRepository) DeleteUser(userID string) error {
	r.log.Infof("Deleting user with ID: %s", userID)
	// db, err := Util.ConnectToPostgreSQL()
	// if err != nil {
	// 	r.log.Error("Failed to connect to PostgreSQL:", err)
	// 	return err
	// }

	if err := r.db.Delete(&BusinessObjects.User{}, "user_id = ?", userID).Error; err != nil {
		r.log.Error("Failed to delete user:", err)
		return err
	}

	r.log.Infof("Successfully deleted user with ID: %s", userID)
	return nil
}

func (r *UserRepository) GetUserByEmail(email string) (BusinessObjects.User, error) {
	r.log.Infof("Fetching user by email: %s", email)
	// db, err := Util.ConnectToPostgreSQL()
	// if err != nil {
	// 	r.log.Error("Failed to connect to PostgreSQL:", err)
	// 	return BusinessObjects.User{}, fmt.Errorf("database connection failed: %w", err)
	// }

	var user BusinessObjects.User
	if err := r.db.First(&user, "email = ?", email).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			r.log.Infof("No user found with email: %s", email)
			// No record found, return an empty user and nil error
			return BusinessObjects.User{}, nil
		}
		r.log.Error("Error querying user by email:", err)
		// Other errors
		return BusinessObjects.User{}, fmt.Errorf("error querying user by email: %w", err)
	}

	r.log.Infof("Successfully fetched user by email: %s", email)
	return user, nil
}

func (r *UserRepository) StoreToken(user *BusinessObjects.User, token string) error {
	r.log.Infof("Storing token for user ID: %s", user.UserID)
	// db, err := Util.ConnectToPostgreSQL()
	// if err != nil {
	// 	r.log.Error("Failed to connect to PostgreSQL:", err)
	// 	return err
	// }

	user.Token = token
	user.TokenExpires = time.Now().Add(time.Hour * 24) // Token valid for 24 hours

	// Save to database (pseudocode)
	if err := r.db.Save(user).Error; err != nil {
		r.log.Error("Failed to store token:", err)
		return err
	}

	r.log.Infof("Successfully stored token for user ID: %s", user.UserID)
	return nil
}

func (r *UserRepository) SetToken(user *BusinessObjects.User, token string) error {
	r.log.Infof("Setting token for user ID: %s", user.UserID)
	// db, err := Util.ConnectToPostgreSQL()
	// if err != nil {
	// 	r.log.Error("Failed to connect to PostgreSQL:", err)
	// 	return err
	// }

	user.Token = token
	user.TokenExpires = time.Now().Add(time.Hour * 24) // Token valid for 24 hours

	// Save to database (pseudocode)
	if err := r.db.Save(user).Error; err != nil {
		r.log.Error("Failed to set token:", err)
		return err
	}

	r.log.Infof("Successfully set token for user ID: %s", user.UserID)
	return nil
}

func (r *UserRepository) VerifyToken(token string) bool {
	r.log.Infof("Verifying token: %s", token)
	// db, err := Util.ConnectToPostgreSQL()
	// if err != nil {
	// 	r.log.Error("Failed to connect to PostgreSQL:", err)
	// 	return false
	// }

	var user BusinessObjects.User
	if err := r.db.Where("token = ?", token).First(&user).Error; err != nil {
		r.log.Error("Failed to verify token:", err)
		return false
	}

	// Check token expiry
	if time.Now().After(user.TokenExpires) {
		r.log.Infof("Token expired for user ID: %s", user.UserID)
		return false
	}

	r.log.Infof("Successfully verified token for user ID: %s", user.UserID)
	return true
}

func (r *UserRepository) GetUserByToken(token string) (BusinessObjects.User, error) {
	r.log.Infof("Fetching user by token: %s", token)
	// db, err := Util.ConnectToPostgreSQL()
	// if err != nil {
	// 	r.log.Error("Failed to connect to PostgreSQL:", err)
	// 	return BusinessObjects.User{}, err
	// }

	var user BusinessObjects.User
	if err := r.db.First(&user, "token = ?", token).Error; err != nil {
		r.log.Error("Failed to fetch user by token:", err)
		return BusinessObjects.User{}, err
	}

	r.log.Infof("Successfully fetched user by token: %s", token)
	return user, nil
}
