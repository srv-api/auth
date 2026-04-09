package repositories

import (
	"time"

	dto "github.com/srv-api/auth/dto/auth"
	"github.com/srv-api/auth/entity"
)

func (u *verifyRepository) UpdateUserVerificationStatus(user *dto.VerificationResponse) error {
	// Update status verified di tabel UserVerified
	now := time.Now()

	err := u.DB.Model(&entity.UserVerified{}).
		Where("id = ?", user.ID).
		Updates(map[string]interface{}{
			"verified":        true,
			"status_account":  true,
			"account_expired": now.AddDate(0, 6, 0),
		}).Error

	if err != nil {
		return err
	}

	// ✅ Update status suspended di tabel AccessDoor
	err = u.DB.Model(&entity.AccessDoor{}).
		Where("id = ?", user.UserID).
		Update("suspended", false).Error

	return err
}
