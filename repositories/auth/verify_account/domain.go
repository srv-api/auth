package repositories

import (
	"sync"

	dto "github.com/srv-cashpay/auth/dto/auth"
	"github.com/srv-cashpay/auth/entity"

	"gorm.io/gorm"
)

type DomainRepository interface {
	UpdateUserVerificationStatus(user *entity.UserVerified) error
	VerifyUserByToken(req dto.VerificationRequest) (*entity.UserVerified, error)
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
