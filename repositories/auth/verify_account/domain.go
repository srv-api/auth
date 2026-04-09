package repositories

import (
	"sync"

	dto "github.com/srv-api/auth/dto/auth"
	"github.com/srv-api/auth/entity"

	"gorm.io/gorm"
)

type DomainRepository interface {
	UpdateUserVerificationStatus(user *dto.VerificationResponse) error
	VerifyUserByToken(req dto.VerificationRequest) (*dto.VerificationResponse, error)
	ResendVerifyUserByToken(req dto.ResendVerificationRequest) (*entity.UserVerified, error)
}

type verifyRepository struct {
	DB    *gorm.DB
	mu    sync.Mutex
	users map[string]*entity.AccessDoor
}

func NewVerifyRepository(DB *gorm.DB) DomainRepository {
	return &verifyRepository{
		DB: DB,
	}
}
