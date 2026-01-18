package auth

import (
	dto "github.com/srv-cashpay/auth/dto/auth"
	"github.com/srv-cashpay/auth/entity"
	util "github.com/srv-cashpay/util/s"
)

func (u *authRepository) Profile(req dto.ProfileRequest) (dto.ProfileResponse, error) {
	var existingUser entity.AccessDoor

	if err := u.DB.Where("id = ?", req.UserID).Find(&existingUser).Error; err != nil {
		return dto.ProfileResponse{}, err
	}

	// Encrypt the email
	encryptedEmail, err := util.Decrypt(existingUser.Email)
	if err != nil {
		return dto.ProfileResponse{}, err
	}

	// Encrypt the email
	encryptedWhatsapp, err := util.Decrypt(existingUser.Whatsapp)
	if err != nil {
		return dto.ProfileResponse{}, err
	}

	resp := dto.ProfileResponse{
		ID:       existingUser.ID,
		FullName: existingUser.FullName,
		Whatsapp: encryptedWhatsapp,
		Email:    encryptedEmail,
	}

	return resp, nil
}
