package auth

import (
	dto "github.com/srv-cashpay/auth/dto/auth"
	util "github.com/srv-cashpay/util/s"
)

func (s *authService) Authenticator(req dto.AuthenticatorRequest) (dto.AuthenticatorResponse, error) {
	create := dto.AuthenticatorRequest{
		ID:       util.GenerateRandomString(),
		Status:   req.Status,
		TokenApp: util.GenerateRandomString(),
		Otp:      util.GenerateRandomNumeric(6),
	}

	created, err := s.Repo.Authenticator(create)
	if err != nil {
		return dto.AuthenticatorResponse{}, err
	}

	response := dto.AuthenticatorResponse{
		ID:       created.ID,
		Status:   created.Status,
		TokenApp: created.TokenApp,
		Otp:      created.Otp,
	}

	return response, nil
}
