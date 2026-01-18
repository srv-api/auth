package entity

import (
	"time"

	"github.com/srv-cashpay/merchant/entity"
	"gorm.io/gorm"
)

type UserMerchant struct {
	ID             string                `gorm:"primary_key;type:varchar(39)" json:"id"`
	FullName       string                `gorm:"full_name;type:varchar(70)" json:"full_name"`
	Whatsapp       string                `gorm:"uniqueIndex;type:varchar(200)" json:"whatsapp"`
	Email          string                `gorm:"uniqueIndex;type:varchar(150)" json:"email"`
	Password       string                `gorm:"password" json:"password"`
	GoogleID       string                `gorm:"google_id" json:"google_id"`
	Provider       string                `gorm:"provider" json:"provider"`
	AccessRoleID   string                `gorm:"access_role_id" json:"access_role_id"`
	LoginAttempts  int                   `gorm:"login_attempts" json:"login_attempts"`
	Suspended      bool                  `gorm:"suspended" json:"suspended"`
	LastAttempt    time.Time             `gorm:"last_attempt" json:"last_attempt"`
	Merchant       entity.MerchantDetail `json:"merchant" gorm:"foreignKey:UserID"`
	ProfilePicture ProfilePicture        `json:"profile_picture" gorm:"foreignKey:UserID"`
	Verified       UserVerified          `json:"verified" gorm:"foreignKey:UserID"`
	File           []File                `json:"file" gorm:"foreignKey:UserID"`
	Platform       string                `gorm:"platform" json:"platform"`
	CreatedBy      string                `gorm:"created_by" json:"created_by"`
	UpdatedBy      string                `gorm:"updated_by" json:"updated_by"`
	DeletedBy      string                `gorm:"deleted_by" json:"deleted_by"`
	CreatedAt      time.Time             `json:"created_at"`
	UpdatedAt      time.Time             `json:"updated_at"`
	DeletedAt      gorm.DeletedAt        `gorm:"index" json:"deleted_at"`
}
