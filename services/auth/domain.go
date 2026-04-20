package auth

import (
	dto "github.com/srv-api/auth/dto/auth"
	m "github.com/srv-api/middlewares/middlewares"

	r "github.com/srv-api/auth/repositories/auth"
)

type AuthService interface {
	Upload(req dto.ProfilePictureRequest) (dto.ProfilePictureResponse, error)
	GetPicture(req dto.GetProfilePictureRequest) (*dto.GetProfilePictureResponse, error)
	Signup(req dto.SignupRequest) (dto.SignupResponse, error)
	Authenticator(req dto.AuthenticatorRequest) (dto.AuthenticatorResponse, error)
	Signin(req dto.SigninRequest) (*dto.SigninResponse, error)
	SigninByPhoneNumber(req dto.SigninRequest) (*dto.SigninResponse, error)
	Profile(req dto.ProfileRequest) (dto.ProfileResponse, error)
	ProfileVisit(req dto.ProfileVisitRequest) (dto.ProfileResponse, error)
	UpdateProfile(req dto.UpdateProfileRequest) (dto.UpdateProfileResponse, error)
	RefreshAccessToken(req dto.RefreshTokenRequest) (string, error)
	SignInWithGoogle(req dto.GoogleSignInRequest) (*dto.AuthResponse, error)
	SignInWithGoogleWeb(req dto.GoogleSignInWebRequest) (*dto.AuthResponse, error)
	Gallery(req dto.GalleryUploadRequest) (dto.MultipleGalleryResponse, error)
	processSingleGalleryFile(req dto.SingleGalleryRequest) (dto.GalleryResponse, error)
	GetUserGallery(userID string) ([]dto.GetGalleryResponse, error)
	GetPictureGallery(req dto.GetPictureGalleryRequest) (*dto.GetGalleryResponse, error)
	DeleteGalleryFile(fileID, userID string) error
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
