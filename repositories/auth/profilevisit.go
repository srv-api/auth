package auth

import (
	dto "github.com/srv-api/auth/dto/auth"
	"github.com/srv-api/auth/entity"
	util "github.com/srv-api/util/s"
)

func (u *authRepository) ProfileVisit(req dto.ProfileVisitRequest) (dto.ProfileResponse, error) {
	var existingUser entity.AccessDoor

	if err := u.DB.Preload("ProfilePicture").Preload("File").Where("id = ?", req.ID).First(&existingUser).Error; err != nil {
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
	baseURL := "http://103.150.227.223:2356/profile/"
	galleryURL := "http://103.150.227.223:2356/box/"

	profilePicture := ""
	if existingUser.ProfilePicture.FilePath != "" {
		profilePicture = baseURL + existingUser.ProfilePicture.FilePath
	}
	// Convert gallery to DTO
	gallery := make([]dto.Gallery, 0)
	for _, img := range existingUser.File {
		gallery = append(gallery, dto.Gallery{
			ID:       img.ID,
			FileName: img.FileName,
			FilePath: galleryURL + img.FilePath,
		})
	}

	resp := dto.ProfileResponse{
		ID:             existingUser.ID,
		FullName:       existingUser.FullName,
		Whatsapp:       encryptedWhatsapp,
		Email:          encryptedEmail,
		Gender:         existingUser.Gender,
		ProfilePicture: profilePicture,
		Gallery:        gallery,
	}

	return resp, nil
}
