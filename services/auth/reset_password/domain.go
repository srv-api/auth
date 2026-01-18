package auth

import (
	dto "github.com/srv-cashpay/auth/dto/auth"
	"github.com/srv-cashpay/auth/entity"
	r "github.com/srv-cashpay/auth/repositories/auth/reset_password"
	m "github.com/srv-cashpay/middlewares/middlewares"
)

type ResetService interface {
	VerifyOtpReset(req dto.VerifyResetRequest) (*entity.PasswordResetToken, error)
	RequestResetPassword(req dto.ResetPasswordRequest) (dto.ResetPasswordResponse, error)
	ResetPassword(req dto.Reset) error
	ResendReset(req dto.ResendResetRequest) (*entity.PasswordResetToken, error)
}

type verifyService struct {
	Repo r.ResetRepository
	jwt  m.JWTService
}

func NewResetService(Repo r.ResetRepository, jwtS m.JWTService) ResetService {
	return &verifyService{
		Repo: Repo,
		jwt:  jwtS,
	}
}
