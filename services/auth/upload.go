package auth

import (
	"errors"
	"fmt"
	"log"
	"os"
	"path/filepath"
	"strings"
	"time"

	dto "github.com/srv-api/auth/dto/auth"
	"github.com/srv-api/detail/entity"
	"gorm.io/gorm"
)

func (s *authService) Upload(req dto.ProfilePictureRequest) (dto.ProfilePictureResponse, error) {
	log.Println("=== Service Upload ===")
	log.Printf("Request: ID=%s, DetailID=%s, UserID=%s", req.ID, req.DetailID, req.UserID)

	// Validasi MerchantDetail
	var merchantDetail entity.UserDetail
	log.Printf("Checking merchant detail for ID: %s", req.DetailID)

	err := s.Repo.CheckMerchantDetail(req.DetailID, &merchantDetail)
	if err != nil {
		log.Printf("CheckMerchantDetail failed: %v", err)
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return dto.ProfilePictureResponse{}, fmt.Errorf("merchant detail not found for detail_id: %s", req.DetailID)
		}
		return dto.ProfilePictureResponse{}, err
	}
	log.Printf("Merchant detail found: %+v", merchantDetail)

	// Validate file type
	allowedExtensions := []string{".jpg", ".jpeg", ".png"}
	ext := strings.ToLower(filepath.Ext(req.File.Filename))
	log.Printf("File extension: %s", ext)

	isAllowed := false
	for _, allowed := range allowedExtensions {
		if ext == allowed {
			isAllowed = true
			break
		}
	}

	if !isAllowed {
		log.Printf("Invalid file type: %s", ext)
		return dto.ProfilePictureResponse{}, errors.New("invalid file type: only JPG and PNG are allowed")
	}

	// Generate unique file name
	newFileName := fmt.Sprintf("%d%s", time.Now().UnixNano(), ext)
	destinationDir := "uploads"
	log.Printf("Creating directory: %s", destinationDir)

	if err := os.MkdirAll(destinationDir, os.ModePerm); err != nil {
		log.Printf("Failed to create directory: %v", err)
		return dto.ProfilePictureResponse{}, fmt.Errorf("failed to create uploads directory: %w", err)
	}

	fullPath := filepath.Join(destinationDir, newFileName)
	req.Destination = fullPath
	log.Printf("Destination path: %s", fullPath)

	// Save file and metadata
	log.Println("Calling Repo.SaveFile...")
	response, err := s.Repo.SaveFile(req)
	if err != nil {
		log.Printf("SaveFile failed: %v", err)
		return dto.ProfilePictureResponse{}, err
	}

	log.Printf("SaveFile success, response: %+v", response)
	return response, nil
}
