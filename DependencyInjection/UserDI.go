package DependencyInjection

import (
	"th3y3m/e-commerce-platform/Interface"
	"th3y3m/e-commerce-platform/Repositories"
	"th3y3m/e-commerce-platform/Services"
	"th3y3m/e-commerce-platform/Util"
	"time"

	"github.com/go-co-op/gocron"
	"github.com/sirupsen/logrus"
)

func NewNewAuthenticationServiceProvider() Interface.IAuthenticationService {
	userRepository := NewUserRepositoryProvider()
	userService := NewUserServiceProvider()
	log := logrus.New()
	scheduler := gocron.NewScheduler(time.UTC)
	deleteJobs := make(map[string]*gocron.Job)
	scheduler.StartAsync()
	return Services.NewAuthenticationService(userRepository, userService, log, scheduler, deleteJobs)
}

func NewUserRepositoryProvider() Interface.IUserRepository {
	log := logrus.New()
	db, err := Util.ConnectToPostgreSQL()
	if err != nil {
		log.Error("Failed to connect to PostgreSQL:", err)
		return nil
	}
	return Repositories.NewUserRepository(log, db)
}

func NewUserServiceProvider() Interface.IUserService {
	log := logrus.New()
	userRepository := NewUserRepositoryProvider()
	return Services.NewUserService(userRepository, log)
}

func NewOAuthServiceProvider() Interface.IOAuthService {
	userRepository := NewUserRepositoryProvider()
	return Services.NewOAuthService(userRepository)
}

func NewMailServiceProvider() Interface.IMailService {
	userRepository := NewUserRepositoryProvider()
	return Services.NewMailService(userRepository)
}
