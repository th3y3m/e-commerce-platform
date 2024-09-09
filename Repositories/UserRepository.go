package Repositories

import (
	"log"
	"th3y3m/e-commerce-platform/BusinessObjects"
	"th3y3m/e-commerce-platform/Util"
)

func GetAllUsers() ([]BusinessObjects.User, error) {
	db, err := Util.ConnectToSQLServer()
	if err != nil {
		log.Fatalf("Error connecting to SQL Server: %v", err)
		return nil, err
	}
	defer db.Close()

	rows, err := db.Query("SELECT * FROM Users")
	if err != nil {
		log.Fatalf("Error querying the database: %v", err)
		return nil, err
	}
	defer rows.Close()

	users := []BusinessObjects.User{}

	for rows.Next() {
		var user BusinessObjects.User
		err := rows.Scan(&user.UserID, &user.Username, &user.PasswordHash, &user.Email, &user.FullName, &user.PhoneNumber, &user.Address, &user.UserType, &user.CreatedAt)
		if err != nil {
			log.Fatalf("Error scanning row: %v", err)
			return nil, err
		}
		users = append(users, user)
	}

	return users, nil
}

func GetUserByID(userID string) (BusinessObjects.User, error) {
	db, err := Util.ConnectToSQLServer()
	if err != nil {
		log.Fatalf("Error connecting to SQL Server: %v", err)
		return BusinessObjects.User{}, err
	}
	defer db.Close()

	var user BusinessObjects.User

	err = db.QueryRow("SELECT * FROM Users WHERE UserID = ?", userID).Scan(&user.UserID, &user.Username, &user.PasswordHash, &user.Email, &user.FullName, &user.PhoneNumber, &user.Address, &user.UserType, &user.CreatedAt)
	if err != nil {
		log.Fatalf("Error querying the database: %v", err)
		return BusinessObjects.User{}, err
	}

	return user, nil
}

func CreateUser(user BusinessObjects.User) error {
	db, err := Util.ConnectToSQLServer()
	if err != nil {
		log.Fatalf("Error connecting to SQL Server: %v", err)
		return err
	}
	defer db.Close()

	_, err = db.Exec("INSERT INTO Users (UserID, Username, PasswordHash, Email, FullName, PhoneNumber, Address, UserType, CreatedAt) VALUES (?, ?, ?, ?, ?, ?, ?, ?, ?)", user.UserID, user.Username, user.PasswordHash, user.Email, user.FullName, user.PhoneNumber, user.Address, user.UserType, user.CreatedAt)
	if err != nil {
		log.Fatalf("Error inserting into the database: %v", err)
		return err
	}

	return nil
}

func UpdateUser(user BusinessObjects.User) error {
	db, err := Util.ConnectToSQLServer()
	if err != nil {
		log.Fatalf("Error connecting to SQL Server: %v", err)
		return err
	}
	defer db.Close()

	_, err = db.Exec("UPDATE Users SET Username = ?, PasswordHash = ?, Email = ?, FullName = ?, PhoneNumber = ?, Address = ?, UserType = ?, CreatedAt = ? WHERE UserID = ?", user.Username, user.PasswordHash, user.Email, user.FullName, user.PhoneNumber, user.Address, user.UserType, user.CreatedAt, user.UserID)
	if err != nil {
		log.Fatalf("Error updating the database: %v", err)
		return err
	}

	return nil
}

func DeleteUser(userID string) error {
	db, err := Util.ConnectToSQLServer()
	if err != nil {
		log.Fatalf("Error connecting to SQL Server: %v", err)
		return err
	}
	defer db.Close()

	_, err = db.Exec("DELETE FROM Users WHERE UserID = ?", userID)
	if err != nil {
		log.Fatalf("Error deleting from the database: %v", err)
		return err
	}

	return nil
}
