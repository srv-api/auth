package auth

import (
	dto "github.com/srv-api/auth/dto/auth"
)

func (u *authService) ProfileVisit(req dto.ProfileVisitRequest) (dto.ProfileResponse, error) {
	// Validasi refresh token dan dapatkan user ID

	comments, err := u.Repo.ProfileVisit(req)
	if err != nil {
		return dto.ProfileResponse{}, err
	}

	return comments, nil
}
