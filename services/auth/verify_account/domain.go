package auth

import (
	dto "github.com/srv-cashpay/auth/dto/auth"
	"github.com/srv-cashpay/auth/entity"
	r "github.com/srv-cashpay/auth/repositories/auth/verify_account"
	m "github.com/srv-cashpay/middlewares/middlewares"
)

type VerifyService interface {
	VerifyUserByToken(req dto.VerificationRequest) (*entity.UserVerified, error)
	ResendVerifyUserByToken(req dto.ResendVerificationRequest) (*entity.UserVerified, error)
}

type verifyService struct {
	Repo r.DomainRepository
	jwt  m.JWTService
}

func NewVerifyService(Repo r.DomainRepository, jwtS m.JWTService) VerifyService {
	return &verifyService{
		Repo: Repo,
		jwt:  jwtS,
	}
}
