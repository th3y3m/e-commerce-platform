package Services

import (
	"errors"
	"th3y3m/e-commerce-platform/Repositories"
	"th3y3m/e-commerce-platform/Util"
)

func Login(email, password string) (string, error) {
	if email == "" || password == "" {
		return "", errors.New("email and password are required")
	}

	user, err := Repositories.GetUserByEmail(email)

	if err != nil {
		return "", err
	}

	if Util.CheckPasswordHash(password, user.PasswordHash) {
		return "", errors.New("invalid password")
	}

	token, err := Util.GenerateToken(user)

	if err != nil {
		return "", err
	}

	return token, nil
}

func RegisterCustomer(email, password string) error {
	if email == "" || password == "" {
		return errors.New("email and password are required")
	}

	hash, err := Util.HashPassword(password)

	if err != nil {
		return err
	}

	err = CreateUser(email, hash, "customer")

	if err != nil {
		return err
	}

	return nil
}
