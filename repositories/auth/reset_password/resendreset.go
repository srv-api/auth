package repositories

import (
	"time"

	dto "github.com/srv-cashpay/auth/dto/auth"
	"github.com/srv-cashpay/auth/entity"

	"gorm.io/gorm"
)

func (u *verifyResetRepository) ResendReset(req dto.ResendResetRequest) (*entity.PasswordResetToken, error) {
	var user entity.PasswordResetToken
	// Fetch user by token from the database
	if err := u.DB.Where("token = ?", req.Token).First(&user).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, err
		}
		return nil, err
	}

	// Update OTP and ExpiredAt in the database
	user.Otp = req.Otp
	user.ExpiredAt = time.Now().Add(4 * time.Minute)
	if err := u.DB.Save(&user).Error; err != nil {
		return nil, err
	}
	return &user, nil
}
