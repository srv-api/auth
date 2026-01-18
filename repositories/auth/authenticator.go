package auth

import (
	dto "github.com/srv-cashpay/auth/dto/auth"
	"github.com/srv-cashpay/auth/entity"
)

func (r *authRepository) Authenticator(req dto.AuthenticatorRequest) (dto.AuthenticatorResponse, error) {

	create := entity.Authenticator{
		ID:       req.ID,
		Status:   req.Status,
		TokenApp: req.TokenApp,
		Otp:      req.Otp,
	}

	if err := r.DB.Save(&create).Error; err != nil {
		return dto.AuthenticatorResponse{}, err
	}

	response := dto.AuthenticatorResponse{
		ID:       req.ID,
		Status:   req.Status,
		TokenApp: create.TokenApp,
		Otp:      create.Otp,
	}

	return response, nil

}
