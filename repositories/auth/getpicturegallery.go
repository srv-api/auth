package auth

import (
	dto "github.com/srv-api/auth/dto/auth"
	"github.com/srv-api/auth/entity"
)

func (b *authRepository) GetPictureGallery(req dto.GetPictureGalleryRequest) (*dto.GetGalleryResponse, error) {
	tr := entity.File{
		FilePath: req.FilePath,
	}

	if err := b.DB.Where("file_path = ?", tr.FilePath).Take(&tr).Error; err != nil {
		return nil, err
	}

	response := &dto.GetGalleryResponse{
		FileName: tr.FileName,
		FilePath: tr.FilePath,
	}

	return response, nil
}
