package auth

import (
	"fmt"
	"io"
	"log"
	"os"
	"path/filepath"

	dto "github.com/srv-api/auth/dto/auth"
	"github.com/srv-api/auth/entity"
	detail "github.com/srv-api/detail/entity"
	res "github.com/srv-api/util/s/response"
)

func (r *authRepository) SaveFile(req dto.ProfilePictureRequest) (dto.ProfilePictureResponse, error) {
	log.Println("=== Repository SaveFile ===")
	log.Printf("Destination: %s", req.Destination)

	// Save the file physically
	src, err := req.File.Open()
	if err != nil {
		log.Printf("Failed to open source file: %v", err)
		return dto.ProfilePictureResponse{}, fmt.Errorf("failed to open source file: %w", err)
	}
	defer src.Close()

	// Ensure the destination directory exists
	dir := filepath.Dir(req.Destination)
	log.Printf("Creating directory: %s", dir)
	if err := os.MkdirAll(dir, os.ModePerm); err != nil {
		log.Printf("Failed to create directory: %v", err)
		return dto.ProfilePictureResponse{}, fmt.Errorf("failed to create directory: %w", err)
	}

	// Create the destination file
	dst, err := os.Create(req.Destination)
	if err != nil {
		log.Printf("Failed to create destination file: %v", err)
		return dto.ProfilePictureResponse{}, fmt.Errorf("failed to create destination file: %w", err)
	}
	defer dst.Close()

	// Copy file content
	bytesWritten, err := io.Copy(dst, src)
	if err != nil {
		log.Printf("Failed to copy file content: %v", err)
		return dto.ProfilePictureResponse{}, fmt.Errorf("failed to copy file content: %w", err)
	}
	log.Printf("File copied successfully, bytes: %d", bytesWritten)

	// Prepare metadata for database
	fileRecord := entity.UploadedFile{
		FileName:  filepath.Base(req.Destination),
		FilePath:  req.Destination,
		DetailID:  req.DetailID,
		UserID:    req.UserID,
		ID:        req.ID,
		CreatedBy: req.CreatedBy,
	}
	log.Printf("Saving metadata: %+v", fileRecord)

	// Save metadata to database
	if err := r.DB.Create(&fileRecord).Error; err != nil {
		log.Printf("Failed to save metadata: %v", err)
		return dto.ProfilePictureResponse{}, fmt.Errorf("failed to save file metadata to database: %w", err)
	}
	log.Println("Metadata saved successfully")

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
