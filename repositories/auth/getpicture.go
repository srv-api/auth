package auth

import (
	dto "github.com/srv-api/auth/dto/auth"
	"github.com/srv-api/auth/entity"
)

func (b *authRepository) GetPicture(req dto.GetProfilePictureRequest) (*dto.GetProfilePictureResponse, error) {
	tr := entity.ProfilePicture{
		FileName: req.FileName,
	}

	if err := b.DB.Where("file_name = ?", tr.FileName).Take(&tr).Error; err != nil {
		return nil, err
	}

	response := &dto.GetProfilePictureResponse{
		FileName: tr.FileName,
		FilePath: tr.FilePath,
	}

	return response, nil
}
