package auth

import (
	dto "github.com/srv-cashpay/auth/dto/auth"
)

func (u *authService) Profile(req dto.ProfileRequest) (dto.ProfileResponse, error) {
	// Validasi refresh token dan dapatkan user ID

	comments, err := u.Repo.Profile(req)
	if err != nil {
		return dto.ProfileResponse{}, err
	}

	return comments, nil
}
