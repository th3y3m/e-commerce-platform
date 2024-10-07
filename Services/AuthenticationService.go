package Services

import (
	"errors"
	"fmt"
	"time"

	"th3y3m/e-commerce-platform/Interface"
	"th3y3m/e-commerce-platform/Util"

	"github.com/go-co-op/gocron"
	"github.com/sirupsen/logrus"
)

type AuthenticationService struct {
	userRepository Interface.IUserRepository
	userService    Interface.IUserService
	mailService    Interface.IMailService
	log            *logrus.Logger
	scheduler      *gocron.Scheduler
	deleteJobs     map[string]*gocron.Job
}

func NewAuthenticationService(userRepository Interface.IUserRepository, userService Interface.IUserService, log *logrus.Logger, scheduler *gocron.Scheduler, deleteJobs map[string]*gocron.Job, mailService Interface.IMailService) Interface.IAuthenticationService {
	return &AuthenticationService{
		userRepository: userRepository,
		userService:    userService,
		log:            log,
		scheduler:      scheduler,
		deleteJobs:     deleteJobs,
		mailService:    mailService,
	}
}

// func init() {
// 	// Start the scheduler in the background
// 	scheduler.StartAsync()
// }

func (a *AuthenticationService) Login(email, password string) (string, error) {
	if email == "" || password == "" {
		return "", errors.New("email and password are required")
	}

	user, err := a.userRepository.GetUserByEmail(email)
	if err != nil {
		return "", err
	}

	if !Util.CheckPasswordHash(user.PasswordHash, password) {
		return "", errors.New("invalid password")
	}

	token, err := Util.GenerateToken(user)
	if err != nil {
		return "", err
	}

	return token, nil
}

func (a *AuthenticationService) RegisterCustomer(email, password string) error {
	if email == "" || password == "" {
		return errors.New("email and password are required")
	}

	user, err := a.userRepository.GetUserByEmail(email)
	if err != nil {
		return fmt.Errorf("error checking user existence: %w", err)
	}

	// If the user already exists, prevent registration
	if user.Email != "" {
		return errors.New("user already exists")
	}

	newUser, err := a.userService.CreateUser(email, password, "customer")
	if err != nil {
		return fmt.Errorf("error creating user: %w", err)
	}

	token, err := Util.GenerateToken(newUser)
	if err != nil {
		return fmt.Errorf("error generating token: %w", err)
	}

	if err := a.userRepository.StoreToken(&newUser, token); err != nil {
		return fmt.Errorf("error storing token: %w", err)
	}

	if err := a.mailService.SendMail(newUser.Email, token); err != nil {
		return fmt.Errorf("error sending mail: %w", err)
	}

	// Schedule user deletion in 1 minute
	job, err := a.scheduler.Every(15).Minutes().StartAt(time.Now().Add(1 * time.Minute)).Do(func() {
		err := a.userRepository.DeleteUser(newUser.UserID)
		if err != nil {
			fmt.Printf("Error deleting user: %v\n", err)
		} else {
			fmt.Printf("User with ID %v deleted due to email verification timeout.\n", newUser.UserID)
		}
	})
	if err != nil {
		return fmt.Errorf("error scheduling user deletion: %w", err)
	}

	a.deleteJobs[newUser.UserID] = job

	return nil
}
func (a *AuthenticationService) RegisterSeller(email, password string) error {
	if email == "" || password == "" {
		return errors.New("email and password are required")
	}

	user, err := a.userRepository.GetUserByEmail(email)
	if err != nil {
		return fmt.Errorf("error checking user existence: %w", err)
	}

	// If the user already exists, prevent registration
	if user.Email != "" {
		return errors.New("user already exists")
	}

	newUser, err := a.userService.CreateUser(email, password, "seller")
	if err != nil {
		return fmt.Errorf("error creating user: %w", err)
	}

	token, err := Util.GenerateToken(newUser)
	if err != nil {
		return fmt.Errorf("error generating token: %w", err)
	}

	if err := a.userRepository.StoreToken(&newUser, token); err != nil {
		return fmt.Errorf("error storing token: %w", err)
	}

	if err := a.mailService.SendMail(newUser.Email, token); err != nil {
		return fmt.Errorf("error sending mail: %w", err)
	}

	// Schedule user deletion in 1 minute
	job, err := a.scheduler.Every(15).Minutes().StartAt(time.Now().Add(1 * time.Minute)).Do(func() {
		err := a.userRepository.DeleteUser(newUser.UserID)
		if err != nil {
			fmt.Printf("Error deleting user: %v\n", err)
		} else {
			fmt.Printf("User with ID %v deleted due to email verification timeout.\n", newUser.UserID)
		}
	})
	if err != nil {
		return fmt.Errorf("error scheduling user deletion: %w", err)
	}

	// Store the job reference so we can cancel it later
	fmt.Printf("Scheduled delete task for user %v.\n", newUser.UserID)
	a.deleteJobs[newUser.UserID] = job

	return nil
}
func (a *AuthenticationService) RegisterAdmin(email, password string) error {
	if email == "" || password == "" {
		return errors.New("email and password are required")
	}

	user, err := a.userRepository.GetUserByEmail(email)
	if err != nil {
		return fmt.Errorf("error checking user existence: %w", err)
	}

	// If the user already exists, prevent registration
	if user.Email != "" {
		return errors.New("user already exists")
	}

	newUser, err := a.userService.CreateUser(email, password, "admin")
	if err != nil {
		return fmt.Errorf("error creating user: %w", err)
	}

	token, err := Util.GenerateToken(newUser)
	if err != nil {
		return fmt.Errorf("error generating token: %w", err)
	}

	if err := a.userRepository.StoreToken(&newUser, token); err != nil {
		return fmt.Errorf("error storing token: %w", err)
	}

	if err := a.mailService.SendMail(newUser.Email, token); err != nil {
		return fmt.Errorf("error sending mail: %w", err)
	}

	// Schedule user deletion in 1 minute
	job, err := a.scheduler.Every(15).Minutes().StartAt(time.Now().Add(1 * time.Minute)).Do(func() {
		err := a.userRepository.DeleteUser(newUser.UserID)
		if err != nil {
			fmt.Printf("Error deleting user: %v\n", err)
		} else {
			fmt.Printf("User with ID %v deleted due to email verification timeout.\n", newUser.UserID)
		}
	})
	if err != nil {
		return fmt.Errorf("error scheduling user deletion: %w", err)
	}

	// Store the job reference so we can cancel it later
	fmt.Printf("Scheduled delete task for user %v.\n", newUser.UserID)
	a.deleteJobs[newUser.UserID] = job

	return nil
}

func (a *AuthenticationService) VerifyUserEmail(token string) error {
	if !a.mailService.VerifyToken(token) {
		return errors.New("token expired or invalid")
	}

	// Extract the user ID from the token (using your existing JWT decode logic)
	userID, err := Util.DecodeJWT(token)
	if err != nil {
		return fmt.Errorf("error decoding token: %w", err)
	}

	// Cancel the scheduled delete user task
	if job, exists := a.deleteJobs[userID]; exists {
		// Remove the job from the scheduler by reference
		fmt.Printf("Cancel Job: %v\n", job)
		a.scheduler.RemoveByReference(job)
		delete(a.deleteJobs, userID) // Remove the job from the map
		fmt.Printf("Scheduled delete task for user %v canceled.\n", userID)
	} else {
		fmt.Printf("No scheduled delete task found for user %v.\n", userID)
		return errors.New("no scheduled delete task found")
	}

	return nil
}

// func RegisterCustomer(email, password string) error {
// 	if email == "" || password == "" {
// 		return errors.New("email and password are required")
// 	}

// 	user, err := Repositories.GetUserByEmail(email)
// 	if err != nil {
// 		return fmt.Errorf("error checking user existence: %w", err)
// 	}

// 	// If the user already exists, prevent registration
// 	if user.Email != "" {
// 		return errors.New("user already exists")
// 	}

// 	hash, err := Util.HashPassword(password)
// 	if err != nil {
// 		return fmt.Errorf("error hashing password: %w", err)
// 	}

// 	newUser, err := CreateUser(email, hash, "customer")
// 	if err != nil {
// 		return fmt.Errorf("error creating user: %w", err)
// 	}

// 	token, err := Util.GenerateToken(newUser)
// 	if err != nil {
// 		return fmt.Errorf("error generating token: %w", err)
// 	}

// 	if err := Repositories.StoreToken(&newUser, token); err != nil {
// 		return fmt.Errorf("error storing token: %w", err)
// 	}

// 	if err := SendMail(newUser.Email, newUser.Token); err != nil {
// 		return fmt.Errorf("error sending mail: %w", err)
// 	}

// 	return nil
// }
