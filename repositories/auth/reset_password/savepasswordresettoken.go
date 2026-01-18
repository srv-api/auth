package repositories

import (
	"time"

	"github.com/srv-cashpay/auth/entity"
)

func (u *verifyResetRepository) SavePasswordResetToken(userID string, token string, otp string, expiryDuration time.Duration) (*entity.PasswordResetToken, error) {
	// Create a PasswordResetToken entity

	passwordResetToken := &entity.PasswordResetToken{
		UserID:    userID,
		Token:     token,
		ExpiredAt: time.Now().Add(4 * time.Minute),
		Otp:       otp,
	}

	// Save the token in the database
	if err := u.DB.Create(passwordResetToken).Error; err != nil {
		return nil, err
	}

	return passwordResetToken, nil
}
