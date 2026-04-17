package auth

import (
	"fmt"
	"io"
	"log"
	"os"
	"path/filepath"
	"time"

	"github.com/google/uuid"
	dto "github.com/srv-api/auth/dto/auth"
	"github.com/srv-api/auth/entity"
	"gorm.io/gorm"
)

// SaveGalleryFile menyimpan satu file gallery ke database
func (r *authRepository) SaveGalleryFile(req dto.SingleGalleryRequest) (dto.GalleryResponse, error) {
	log.Println("=== Repository SaveGalleryFile ===")
	log.Printf("File: %s", req.File.Filename)

	// Save the file physically
	src, err := req.File.Open()
	if err != nil {
		log.Printf("Failed to open source file: %v", err)
		return dto.GalleryResponse{}, fmt.Errorf("failed to open source file: %w", err)
	}
	defer src.Close()

	// Ensure the destination directory exists
	dir := filepath.Dir(req.Destination)
	log.Printf("Creating directory: %s", dir)
	if err := os.MkdirAll(dir, os.ModePerm); err != nil {
		log.Printf("Failed to create directory: %v", err)
		return dto.GalleryResponse{}, fmt.Errorf("failed to create directory: %w", err)
	}

	// Create the destination file
	dst, err := os.Create(req.Destination)
	if err != nil {
		log.Printf("Failed to create destination file: %v", err)
		return dto.GalleryResponse{}, fmt.Errorf("failed to create destination file: %w", err)
	}
	defer dst.Close()

	// Copy file content
	bytesWritten, err := io.Copy(dst, src)
	if err != nil {
		log.Printf("Failed to copy file content: %v", err)
		return dto.GalleryResponse{}, fmt.Errorf("failed to copy file content: %w", err)
	}
	log.Printf("File copied successfully, bytes: %d", bytesWritten)

	// Generate ID if not provided
	id := req.ID
	if id == "" {
		id = uuid.New().String()
	}

	// Prepare metadata for database
	now := time.Now()
	fileRecord := entity.File{
		ID:          id,
		FileName:    req.File.Filename,
		FilePath:    req.Destination,
		DetailID:    req.DetailID,
		UserID:      req.UserID,
		DataAccount: "active",
		CreatedBy:   req.CreatedBy,
		UpdatedBy:   req.UpdatedBy,
		CreatedAt:   now,
		UpdatedAt:   now,
	}
	log.Printf("Saving metadata: %+v", fileRecord)

	// Save metadata to database
	if err := r.DB.Create(&fileRecord).Error; err != nil {
		log.Printf("Failed to save metadata: %v", err)
		return dto.GalleryResponse{}, fmt.Errorf("failed to save file metadata to database: %w", err)
	}
	log.Println("Metadata saved successfully")

	return dto.GalleryResponse{
		ID:       fileRecord.ID,
		FileName: fileRecord.FileName,
		FilePath: req.Destination,
		UserID:   fileRecord.UserID,
		DetailID: fileRecord.DetailID,
	}, nil
}

// GetUserGallery mengambil semua gallery user
func (r *authRepository) GetUserGallery(userID string) ([]entity.File, error) {
	var files []entity.File

	err := r.DB.
		Where("user_id = ? AND data_account = ?", userID, "active").
		Order("created_at DESC").
		Find(&files).Error

	if err != nil {
		return nil, err
	}

	return files, nil
}

// DeleteGalleryFile menghapus file gallery (soft delete)
func (r *authRepository) DeleteGalleryFile(fileID, userID string) error {
	// Soft delete with updating data_account and deleted_at
	result := r.DB.
		Model(&entity.File{}).
		Where("id = ? AND user_id = ?", fileID, userID).
		Updates(map[string]interface{}{
			"data_account": "deleted",
			"deleted_at":   time.Now(),
			"deleted_by":   userID,
		})

	if result.Error != nil {
		return result.Error
	}

	if result.RowsAffected == 0 {
		return gorm.ErrRecordNotFound
	}

	return nil
}

// GetGalleryByID mengambil satu file gallery berdasarkan ID
func (r *authRepository) GetGalleryByID(fileID, userID string) (*entity.File, error) {
	var file entity.File

	err := r.DB.
		Where("id = ? AND user_id = ? AND data_account = ?", fileID, userID, "active").
		First(&file).Error

	if err != nil {
		return nil, err
	}

	return &file, nil
}

// DeleteAllUserGallery menghapus semua gallery user (soft delete)
func (r *authRepository) DeleteAllUserGallery(userID string) error {
	result := r.DB.
		Model(&entity.File{}).
		Where("user_id = ? AND data_account = ?", userID, "active").
		Updates(map[string]interface{}{
			"data_account": "deleted",
			"deleted_at":   time.Now(),
		})

	return result.Error
}
