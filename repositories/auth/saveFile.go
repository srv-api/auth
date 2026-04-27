package auth

import (
	"fmt"
	"io"
	"os"
	"path/filepath"
	"time"

	dto "github.com/srv-api/auth/dto/auth"
	"github.com/srv-api/auth/entity"
	detail "github.com/srv-api/detail/entity"
	res "github.com/srv-api/util/s/response"
	"gorm.io/gorm/clause"
)

func (r *authRepository) SaveFile(req dto.ProfilePictureRequest) (dto.ProfilePictureResponse, error) {
	// Save file physically (sama seperti sebelumnya)
	src, err := req.File.Open()
	if err != nil {
		return dto.ProfilePictureResponse{}, fmt.Errorf("failed to open source file: %w", err)
	}
	defer src.Close()

	dir := filepath.Dir(req.Destination)
	if err := os.MkdirAll(dir, os.ModePerm); err != nil {
		return dto.ProfilePictureResponse{}, fmt.Errorf("failed to create directory: %w", err)
	}

	dst, err := os.Create(req.Destination)
	if err != nil {
		return dto.ProfilePictureResponse{}, fmt.Errorf("failed to create destination file: %w", err)
	}
	defer dst.Close()

	if _, err := io.Copy(dst, src); err != nil {
		return dto.ProfilePictureResponse{}, fmt.Errorf("failed to copy file content: %w", err)
	}

	// UPSERT: Insert or Update in one query
	profilePicture := entity.ProfilePicture{
		ID:        req.ID,
		UserID:    req.UserID,
		DetailID:  req.DetailID,
		FileName:  filepath.Base(req.Destination),
		FilePath:  req.Destination,
		CreatedBy: req.CreatedBy,
		CreatedAt: time.Now(),
		UpdatedBy: req.UpdatedBy,
		UpdatedAt: time.Now(),
	}

	err = r.DB.Clauses(clause.OnConflict{
		Columns:   []clause.Column{{Name: "user_id"}},
		DoUpdates: clause.AssignmentColumns([]string{"file_name", "file_path", "updated_by", "updated_at"}),
	}).Create(&profilePicture).Error

	if err != nil {
		return dto.ProfilePictureResponse{}, fmt.Errorf("failed to upsert profile picture: %w", err)
	}

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
