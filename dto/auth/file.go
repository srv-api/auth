package auth

import (
	"mime/multipart"
	"time"
)

type GalleryUploadRequest struct {
	UserID    string                  `json:"user_id"`
	DetailID  string                  `json:"detail_id"`
	CreatedBy string                  `json:"created_by"`
	UpdatedBy string                  `json:"updated_by"`
	Files     []*multipart.FileHeader `form:"images"`
}

type SingleGalleryRequest struct {
	ID          string
	UserID      string
	DetailID    string
	CreatedBy   string
	UpdatedBy   string
	File        *multipart.FileHeader
	Destination string
}

type GalleryResponse struct {
	ID       string `json:"id"`
	FileName string `json:"file_name"`
	FilePath string `json:"file_path"`
	UserID   string `json:"user_id"`
	DetailID string `json:"detail_id"`
}

type MultipleGalleryResponse struct {
	SuccessCount int               `json:"success_count"`
	FailedCount  int               `json:"failed_count"`
	Files        []GalleryResponse `json:"success_files"`
	Errors       []string          `json:"errors,omitempty"`
}

type GetGalleryResponse struct {
	ID          string    `json:"id"`
	UserID      string    `json:"user_id"`
	DetailID    string    `json:"detail_id"`
	FileName    string    `json:"file_name"`
	FilePath    string    `json:"file_path"`
	DataAccount string    `json:"data_account"`
	CreatedAt   time.Time `json:"created_at"`
}
