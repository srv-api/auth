package repositories

import (
	"time"

	dto "github.com/srv-cashpay/auth/dto/auth"
	"github.com/srv-cashpay/auth/entity"

	"gorm.io/gorm"
)

type ResetRepository interface {
	VerifyOtpReset(req dto.VerifyResetRequest) (*entity.PasswordResetToken, error)
	SavePasswordResetToken(userID string, token string, otp string, expiryDuration time.Duration) (*entity.PasswordResetToken, error)
	GetUserByEmail(email string) (*entity.AccessDoor, error)
	ValidatePasswordResetToken(token string) (string, error)
	UpdateUserPassword(userID string, newPassword string) error
	ResendReset(req dto.ResendResetRequest) (*entity.PasswordResetToken, error)
}

type verifyResetRepository struct {
	DB *gorm.DB
}

func NewResetRepository(DB *gorm.DB) ResetRepository {
	return &verifyResetRepository{
		DB: DB,
	}
}
