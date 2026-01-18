package repositories

import (
	"github.com/srv-cashpay/auth/entity"
)

func (u *verifyResetRepository) GetUserByEmail(email string) (*entity.AccessDoor, error) {
	var user entity.AccessDoor
	if err := u.DB.Where("email = ?", email).First(&user).Error; err != nil {
		return nil, err
	}
	return &user, nil
}
