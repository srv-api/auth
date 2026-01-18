package repositories

import (
	"github.com/srv-cashpay/auth/entity"
)

func (u *verifyResetRepository) UpdateUserPassword(userID string, newPassword string) error {
	var user entity.AccessDoor

	if err := u.DB.Where("id = ?", userID).First(&user).Error; err != nil {
		return err
	}

	user.Password = newPassword
	if err := u.DB.Save(&user).Error; err != nil {
		return err
	}

	return nil
}
