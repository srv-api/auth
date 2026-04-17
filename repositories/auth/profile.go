package auth

import (
	dto "github.com/srv-api/auth/dto/auth"
	"github.com/srv-api/auth/entity"
	util "github.com/srv-api/util/s"
)

func (u *authRepository) Profile(req dto.ProfileRequest) (dto.ProfileResponse, error) {
	var existingUser entity.AccessDoor

	if err := u.DB.Preload("ProfilePicture").Where("id = ?", req.UserID).First(&existingUser).Error; err != nil {
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
	baseURL := "http://103.150.227.223:2356"
	profilePicture := ""
	if existingUser.ProfilePicture.FilePath != "" {
		profilePicture = baseURL + existingUser.ProfilePicture.FilePath
	}

	resp := dto.ProfileResponse{
		ID:             existingUser.ID,
		FullName:       existingUser.FullName,
		Whatsapp:       encryptedWhatsapp,
		Email:          encryptedEmail,
		ProfilePicture: profilePicture,
	}

	return resp, nil
}
