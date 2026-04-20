package entity

import (
	"time"

	"gorm.io/gorm"
)

type ProfilePicture struct {
	ID          string         `gorm:"primary_key" json:"id"`
	UserID      string         `gorm:"type:varchar(36);index" json:"user_id"`
	DetailID    string         `gorm:"type:varchar(36);index" json:"detail_id"`
	FileName    string         `gorm:"file_name" json:"file_name"`
	FilePath    string         `gorm:"file_path" json:"file_path"`
	DataAccount string         `gorm:"status,omitempty" json:"data_account"`
	CreatedBy   string         `gorm:"created_by" json:"created_by"`
	UpdatedBy   string         `gorm:"updated_by" json:"updated_by"`
	DeletedBy   string         `gorm:"deleted_by" json:"deleted_by"`
	CreatedAt   time.Time      `json:"created_at"`
	UpdatedAt   time.Time      `json:"updated_at"`
	DeletedAt   gorm.DeletedAt `gorm:"index" json:"deleted_at"`
}
