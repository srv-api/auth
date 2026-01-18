package entity

import (
	"time"
)

type PasswordResetToken struct {
	ID        uint      `gorm:"primaryKey"`
	UserID    string    `gorm:"type:varchar(36);index" json:"user_id"`
	Token     string    `gorm:"token" json:"token"`
	Otp       string    `gorm:"otp" json:"otp"`
	CreatedAt time.Time `gorm:"autoCreateTime"`
	ExpiredAt time.Time `gorm:"expired_at" json:"expired_at"`
}
