package auth

import (
	"sync"

	dto "github.com/srv-api/auth/dto/auth"
	detail "github.com/srv-api/detail/entity"

	"github.com/srv-api/auth/entity"

	"gorm.io/gorm"
)

type DomainRepository interface {
	CheckMerchantDetail(DetailID string, merchantDetail *detail.UserDetail) error
	Signup(req dto.SignupRequest) (dto.SignupResponse, error)
	Authenticator(req dto.AuthenticatorRequest) (dto.AuthenticatorResponse, error)
	Signin(req dto.SigninRequest) (*entity.AccessDoor, error)
	UpdateTokenVerified(userID string, otp string, token string) (dto.SigninResponse, error)
	UpdateUser(user *entity.AccessDoor) error
	SigninByPhoneNumber(req dto.SigninRequest) (*entity.AccessDoor, error)
	RefreshToken(req dto.RefreshTokenRequest) (*entity.AccessDoor, error)
	SaveUser(user *entity.AccessDoor) error
	Profile(req dto.ProfileRequest) (dto.ProfileResponse, error)
	UpdateProfile(req dto.UpdateProfileRequest) (dto.UpdateProfileResponse, error)
	FindByEncryptedEmail(encryptedEmail string) (*entity.AccessDoor, error)
	Create(user *entity.AccessDoor) error
	UpdateWhatsapp(userID string, phone string) error
	SaveFile(req dto.ProfilePictureRequest) (dto.ProfilePictureResponse, error)
	GetPicture(req dto.GetProfilePictureRequest) (*dto.GetProfilePictureResponse, error)
}

type authRepository struct {
	DB    *gorm.DB
	mu    sync.Mutex
	users map[string]*entity.AccessDoor
}

func NewAuthRepository(DB *gorm.DB) DomainRepository {
	return &authRepository{
		DB: DB,
	}
}
