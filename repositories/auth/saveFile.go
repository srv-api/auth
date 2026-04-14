package auth

import (
	"fmt"
	"io"
	"os"
	"path/filepath"

	dto "github.com/srv-api/auth/dto/auth"
	"github.com/srv-api/auth/entity"
	detail "github.com/srv-api/detail/entity"
	res "github.com/srv-api/util/s/response"
)

func (r *authRepository) SaveFile(req dto.ProfilePictureRequest) (dto.ProfilePictureResponse, error) {
	// Save the file physically
	src, err := req.File.Open()
	if err != nil {
		return dto.ProfilePictureResponse{}, fmt.Errorf("failed to open source file: %w", err)
	}
	defer src.Close()

	// Ensure the destination directory exists
	dir := filepath.Dir(req.Destination)
	if err := os.MkdirAll(dir, os.ModePerm); err != nil {
		return dto.ProfilePictureResponse{}, fmt.Errorf("failed to create directory: %w", err)
	}

	// Create the destination file
	dst, err := os.Create(req.Destination)
	if err != nil {
		return dto.ProfilePictureResponse{}, fmt.Errorf("failed to create destination file: %w", err)
	}
	defer dst.Close()

	// Copy file content
	if _, err := io.Copy(dst, src); err != nil {
		return dto.ProfilePictureResponse{}, fmt.Errorf("failed to copy file content: %w", err)
	}

	// Prepare metadata for database
	fileRecord := entity.UploadedFile{
		FileName:  filepath.Base(req.Destination),
		FilePath:  req.Destination,
		DetailID:  req.DetailID,
		UserID:    req.UserID,
		ID:        req.ID,
		CreatedBy: req.CreatedBy,
	}

	// Save metadata to database
	if err := r.DB.Create(&fileRecord).Error; err != nil {
		return dto.ProfilePictureResponse{}, fmt.Errorf("failed to save file metadata to database: %w", err)
	}

	// Return response
	return dto.ProfilePictureResponse{
		FilePath: req.Destination,
	}, nil
}

func (r *authRepository) CheckMerchantDetail(DetailID string, merchantDetail *detail.UserDetail) error {
	err := r.DB.Where("id = ?", DetailID).First(merchantDetail).Error
	if err != nil {
		return err
	}

	// Periksa apakah semua kolom penting sudah terisi
	if merchantDetail.Latitude == 0 || merchantDetail.Longitude == 0 {
		return res.ErrorBuilder(&res.ErrorConstant.DetailNoProvide, nil)
	}

	return nil
}
