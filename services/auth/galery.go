package auth

import (
	"errors"
	"fmt"
	"log"
	"mime/multipart"
	"os"
	"path/filepath"
	"strings"
	"sync"
	"time"

	"github.com/google/uuid"
	dto "github.com/srv-api/auth/dto/auth"
	datail "github.com/srv-api/detail/entity"

	"gorm.io/gorm"
)

// Gallery untuk upload multiple files
func (s *authService) Gallery(req dto.GalleryUploadRequest) (dto.MultipleGalleryResponse, error) {
	log.Println("=== Service Gallery (Multiple Upload) ===")
	log.Printf("Number of files: %d", len(req.Files))

	// Validasi MerchantDetail
	var merchantDetail datail.UserDetail
	log.Printf("Checking merchant detail for ID: %s", req.DetailID)

	err := s.Repo.CheckMerchantDetail(req.DetailID, &merchantDetail)
	if err != nil {
		log.Printf("CheckMerchantDetail failed: %v", err)
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return dto.MultipleGalleryResponse{}, fmt.Errorf("merchant detail not found for detail_id: %s", req.DetailID)
		}
		return dto.MultipleGalleryResponse{}, err
	}
	log.Printf("Merchant detail found: %+v", merchantDetail)

	// Create uploads directory if not exists
	destinationDir := "uploads/gallery"
	if err := os.MkdirAll(destinationDir, os.ModePerm); err != nil {
		log.Printf("Failed to create directory: %v", err)
		return dto.MultipleGalleryResponse{}, fmt.Errorf("failed to create uploads directory: %w", err)
	}

	// Process each file
	var wg sync.WaitGroup
	var mu sync.Mutex
	successFiles := []dto.GalleryResponse{}
	errors := []string{}

	// Limit concurrent uploads to 5
	semaphore := make(chan struct{}, 5)

	for i, file := range req.Files {
		wg.Add(1)
		go func(index int, f *multipart.FileHeader) {
			defer wg.Done()

			// Acquire semaphore
			semaphore <- struct{}{}
			defer func() { <-semaphore }()

			// Process single file
			resp, err := s.processSingleGalleryFile(dto.SingleGalleryRequest{
				UserID:    req.UserID,
				DetailID:  req.DetailID,
				CreatedBy: req.CreatedBy,
				UpdatedBy: req.UpdatedBy,
				File:      f,
			})

			mu.Lock()
			defer mu.Unlock()

			if err != nil {
				errors = append(errors, fmt.Sprintf("File %s: %v", f.Filename, err))
			} else {
				successFiles = append(successFiles, resp)
			}
		}(i, file)
	}

	wg.Wait()

	log.Printf("Upload completed: %d success, %d failed", len(successFiles), len(errors))

	return dto.MultipleGalleryResponse{
		SuccessCount: len(successFiles),
		FailedCount:  len(errors),
		Files:        successFiles,
		Errors:       errors,
	}, nil
}

// processSingleGalleryFile proses satu file gallery
func (s *authService) processSingleGalleryFile(req dto.SingleGalleryRequest) (dto.GalleryResponse, error) {
	log.Printf("Processing file: %s", req.File.Filename)

	// Validate file size (max 5MB)
	const maxFileSize = 5 * 1024 * 1024
	if req.File.Size > maxFileSize {
		return dto.GalleryResponse{}, fmt.Errorf("file size exceeds 5MB limit (max 5MB)")
	}

	// Validate file type
	allowedExtensions := []string{".jpg", ".jpeg", ".png", ".gif", ".webp"}
	ext := strings.ToLower(filepath.Ext(req.File.Filename))

	isAllowed := false
	for _, allowed := range allowedExtensions {
		if ext == allowed {
			isAllowed = true
			break
		}
	}

	if !isAllowed {
		return dto.GalleryResponse{}, fmt.Errorf("invalid file type: only JPG, JPEG, PNG, GIF, WEBP are allowed")
	}

	// Generate unique file name
	newFileName := fmt.Sprintf("%d_%s%s", time.Now().UnixNano(), uuid.New().String(), ext)
	fullPath := filepath.Join("uploads/gallery", newFileName)

	singleReq := dto.SingleGalleryRequest{
		ID:          uuid.New().String(),
		UserID:      req.UserID,
		DetailID:    req.DetailID,
		CreatedBy:   req.CreatedBy,
		UpdatedBy:   req.UpdatedBy,
		File:        req.File,
		Destination: fullPath,
	}

	// Save file and metadata
	response, err := s.Repo.SaveGalleryFile(singleReq)
	if err != nil {
		return dto.GalleryResponse{}, err
	}

	log.Printf("File processed successfully: %s", newFileName)
	return response, nil
}

// GetUserGallery mengambil semua gallery user
func (s *authService) GetUserGallery(userID string) ([]dto.GetGalleryResponse, error) {
	files, err := s.Repo.GetUserGallery(userID)
	if err != nil {
		return nil, err
	}

	var responses []dto.GetGalleryResponse
	for _, file := range files {
		responses = append(responses, dto.GetGalleryResponse{
			ID:          file.ID,
			UserID:      file.UserID,
			DetailID:    file.DetailID,
			FileName:    file.FileName,
			FilePath:    "http://103.150.227.223:2356/" + file.FilePath,
			DataAccount: file.DataAccount,
			CreatedAt:   file.CreatedAt,
		})
	}

	return responses, nil
}

// DeleteGalleryFile menghapus file gallery
func (s *authService) DeleteGalleryFile(fileID, userID string) error {
	return s.Repo.DeleteGalleryFile(fileID, userID)
}
