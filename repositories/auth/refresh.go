package auth

import (
	dto "github.com/srv-api/auth/dto/auth"
	"github.com/srv-api/auth/entity"
)

func (u *authRepository) RefreshToken(req dto.RefreshTokenRequest) (*entity.AccessDoor, error) {
	var existingUser entity.AccessDoor

	err := u.DB.Preload("Verified").Preload("Merchant").Where("id = ?", req.UserID).First(&existingUser).Error
	if err != nil {
		return nil, err
	}

	return &existingUser, nil
}
