package repositories

import (
	"errors"
	"fmt"
	"time"

	"github.com/srv-cashpay/auth/entity"
)

func (u *verifyResetRepository) ValidatePasswordResetToken(token string) (string, error) {
	// Find the PasswordResetToken with the given token
	var passwordResetToken entity.PasswordResetToken
	if err := u.DB.Where("token = ?", token).First(&passwordResetToken).Error; err != nil {
		// Handle errors s(token not found, database error, etc.)
		return "", err
	}

	currentUTC := time.Now().UTC()
	if currentUTC.After(passwordResetToken.ExpiredAt.UTC()) {
		// Token is expired
		return "", errors.New("Password reset token has expired")
	}
	fmt.Println("Current Time:", currentUTC)
	fmt.Println("Expiration Time:", passwordResetToken.ExpiredAt.UTC())

	// Return the associated user ID if the token is valid
	return passwordResetToken.UserID, nil
}
