package auth

import (
	dto "github.com/srv-cashpay/auth/dto/auth"
	m "github.com/srv-cashpay/middlewares/middlewares"

	r "github.com/srv-cashpay/auth/repositories/auth"
)

type AuthService interface {
	Signup(req dto.SignupRequest) (dto.SignupResponse, error)
	Authenticator(req dto.AuthenticatorRequest) (dto.AuthenticatorResponse, error)
	Signin(req dto.SigninRequest) (*dto.SigninResponse, error)
	SigninByPhoneNumber(req dto.SigninRequest) (*dto.SigninResponse, error)
	Profile(req dto.ProfileRequest) (dto.ProfileResponse, error)
	UpdateProfile(req dto.UpdateProfileRequest) (dto.UpdateProfileResponse, error)
	RefreshAccessToken(req dto.RefreshTokenRequest) (string, error)
	SignInWithGoogle(req dto.GoogleSignInRequest) (*dto.AuthResponse, error)
	SignInWithGoogleWeb(req dto.GoogleSignInWebRequest) (*dto.AuthResponse, error)
}

type authService struct {
	Repo r.DomainRepository
	jwt  m.JWTService
}

func NewAuthService(Repo r.DomainRepository, jwtS m.JWTService) AuthService {
	return &authService{
		Repo: Repo,
		jwt:  jwtS,
	}
}
