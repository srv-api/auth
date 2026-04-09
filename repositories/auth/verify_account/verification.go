package repositories

import (
	"time"

	dto "github.com/srv-api/auth/dto/auth"
	"github.com/srv-api/auth/entity"
)

func (u *verifyRepository) UpdateUserVerificationStatus(user *dto.VerificationResponse) error {
	// Update status verified dan account_expired
	now := time.Now()

	err := u.DB.Model(&entity.UserVerified{}).
		Where("id = ?", user.ID).
		Updates(map[string]interface{}{
			"verified":        true,
			"status_account":  true,
			"account_expired": now.AddDate(0, 6, 0), // 6 bulan
		}).Error

	if err != nil {
		return err
	}

	// ✅ Update juga status di tabel UserMerchant/AccessDoor
	err = u.DB.Model(&entity.UserMerchant{}).
		Where("id = ?", user.UserID).
		Update("suspended", false).Error

	return err
}
