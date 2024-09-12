package Repositories

import (
	"errors"
	"fmt"
	"strings"
	"th3y3m/e-commerce-platform/BusinessObjects"
	"th3y3m/e-commerce-platform/Util"
	"time"

	"gorm.io/gorm"
)

func GetPaginatedUserList(searchValue, sortBy string, pageIndex, pageSize int, status *bool) (Util.PaginatedList[BusinessObjects.User], error) {
	db, err := Util.ConnectToPostgreSQL()
	if err != nil {
		return Util.PaginatedList[BusinessObjects.User]{}, err
	}

	var users []BusinessObjects.User

	// Scope the query to the User table
	query := db.Model(&BusinessObjects.User{})

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
		return Util.PaginatedList[BusinessObjects.User]{}, err
	}

	// Apply pagination
	offset := (pageIndex - 1) * pageSize
	if err := query.Offset(offset).Limit(pageSize).Find(&users).Error; err != nil {
		return Util.PaginatedList[BusinessObjects.User]{}, err
	}

	// Create paginated list
	paginatedList := Util.NewPaginatedList(users, totalCount, pageIndex, pageSize)

	return paginatedList, nil
}

func GetAllUsers() ([]BusinessObjects.User, error) {
	db, err := Util.ConnectToPostgreSQL()
	if err != nil {
		return nil, err
	}

	var users []BusinessObjects.User
	if err := db.Find(&users).Error; err != nil {
		return nil, err
	}

	return users, nil
}

func GetUserByID(userID string) (BusinessObjects.User, error) {
	db, err := Util.ConnectToPostgreSQL()
	if err != nil {
		return BusinessObjects.User{}, err
	}

	var user BusinessObjects.User
	if err := db.First(&user, "user_id = ?", userID).Error; err != nil {
		return BusinessObjects.User{}, err
	}

	return user, nil
}

func CreateUser(user BusinessObjects.User) (BusinessObjects.User, error) {
	db, err := Util.ConnectToPostgreSQL()
	if err != nil {
		return BusinessObjects.User{}, err
	}

	if err := db.Create(&user).Error; err != nil {
		return BusinessObjects.User{}, err
	}

	return user, nil
}

func UpdateUser(user BusinessObjects.User) error {
	db, err := Util.ConnectToPostgreSQL()
	if err != nil {
		return err
	}

	if err := db.Save(&user).Error; err != nil {
		return err
	}

	return nil
}

func DeleteUser(userID string) error {
	db, err := Util.ConnectToPostgreSQL()
	if err != nil {
		return err
	}

	if err := db.Delete(&BusinessObjects.User{}, "user_id = ?", userID).Error; err != nil {
		return err
	}

	return nil
}

func GetUserByEmail(email string) (BusinessObjects.User, error) {
	db, err := Util.ConnectToPostgreSQL()
	if err != nil {
		return BusinessObjects.User{}, fmt.Errorf("database connection failed: %w", err)
	}

	var user BusinessObjects.User
	if err := db.First(&user, "email = ?", email).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			// No record found, return an empty user and nil error
			return BusinessObjects.User{}, nil
		}
		// Other errors
		return BusinessObjects.User{}, fmt.Errorf("error querying user by email: %w", err)
	}

	return user, nil
}

func StoreToken(user *BusinessObjects.User, token string) error {
	db, err := Util.ConnectToPostgreSQL()
	if err != nil {
		return err
	}

	user.Token = token
	user.TokenExpires = time.Now().Add(time.Hour * 24) // Token valid for 24 hours

	// Save to database (pseudocode)
	return db.Save(user).Error
}

func SetToken(user *BusinessObjects.User, token string) error {
	db, err := Util.ConnectToPostgreSQL()
	if err != nil {
		return err
	}

	user.Token = token
	user.TokenExpires = time.Now().Add(time.Hour * 24) // Token valid for 24 hours

	// Save to database (pseudocode)
	return db.Save(user).Error
}

func VerifyToken(token string) bool {
	db, err := Util.ConnectToPostgreSQL()
	if err != nil {
		return false
	}

	var user BusinessObjects.User
	if err := db.Where("token = ?", token).First(&user).Error; err != nil {
		return false
	}

	// Check token expiry
	if time.Now().After(user.TokenExpires) {
		return false
	}

	return true
}
