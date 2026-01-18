package repositories

import (
	"fmt"

	dto "github.com/srv-cashpay/auth/dto/auth"
	"github.com/srv-cashpay/auth/entity"

	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

func (u *verifyResetRepository) VerifyOtpReset(req dto.VerifyResetRequest) (*entity.PasswordResetToken, error) {
	var user entity.PasswordResetToken
	if err := u.DB.Clauses(clause.OnConflict{
		Columns:   []clause.Column{{Name: "token"}},
		DoUpdates: clause.Assignments(map[string]interface{}{"verified": true}),
	}).Where("token = ?", req.Token).Where("otp = ?", req.Otp).First(&user).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, fmt.Errorf("User not found with the given verification token")
		}
		return nil, err
	}

	return &user, nil
}
