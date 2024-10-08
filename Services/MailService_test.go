package Services

import (
	"testing"
	"th3y3m/e-commerce-platform/mocks"

	"github.com/stretchr/testify/assert"
)

func TestVerifyToken_Success(t *testing.T) {
	// Arrange
	userRepository := &mocks.IUserRepository{}
	productRepository := &mocks.IProductRepository{}
	MailService := NewMailService(userRepository, productRepository)

	token := "test"

	userRepository.On("VerifyToken", token).Return(true)

	// Act
	check := MailService.VerifyToken(token)

	// Assert
	assert.Equal(t, check, true)
	userRepository.AssertExpectations(t)
}
