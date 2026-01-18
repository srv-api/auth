package auth

import (
	"fmt"

	dto "github.com/srv-cashpay/auth/dto/auth"
	util "github.com/srv-cashpay/util/s"
)

func (u *authService) UpdateProfile(req dto.UpdateProfileRequest) (dto.UpdateProfileResponse, error) {
	// Prepare updated request
	request := dto.UpdateProfileRequest{
		ID:        req.ID,
		FullName:  req.FullName,
		Email:     req.Email,
		Whatsapp:  req.Whatsapp,
		UpdatedBy: req.UpdatedBy,
	}

	// Encrypt email if provided
	if req.Email != "" {
		encryptedEmail, err := util.Encrypt(req.Email)
		if err != nil {
			return dto.UpdateProfileResponse{}, fmt.Errorf("failed to encrypt email: %w", err)
		}
		request.Email = string(encryptedEmail)
	}

	// Encrypt email if provided
	if req.Whatsapp != "" {
		encryptedWhatsapp, err := util.Encrypt(req.Whatsapp)
		if err != nil {
			return dto.UpdateProfileResponse{}, fmt.Errorf("failed to encrypt Whatsapp: %w", err)
		}
		request.Whatsapp = string(encryptedWhatsapp)
	}

	// Encrypt password if provided
	if req.Password != "" {
		encryptedPassword, err := util.GenerateFromPassword(req.Password)
		if err != nil {
			return dto.UpdateProfileResponse{}, fmt.Errorf("failed to encrypt password: %w", err)
		}
		request.Password = string(encryptedPassword)
	}

	// Update profile in repository
	product, err := u.Repo.UpdateProfile(request)
	if err != nil {
		return dto.UpdateProfileResponse{}, fmt.Errorf("failed to update profile: %w", err)
	}

	// Prepare response
	response := dto.UpdateProfileResponse{
		ID:        product.ID,
		FullName:  product.FullName,
		Email:     product.Email,
		Whatsapp:  product.Whatsapp,
		Password:  product.Password, // Return encrypted password if needed
		UpdatedBy: product.UpdatedBy,
	}

	return response, nil
}
