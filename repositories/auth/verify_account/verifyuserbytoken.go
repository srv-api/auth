package repositories

import (
	"errors"
	"time"

	dto "github.com/srv-api/auth/dto/auth"
	"github.com/srv-api/auth/entity"

	"gorm.io/gorm"
)

func (u *verifyRepository) VerifyUserByToken(req dto.VerificationRequest) (*dto.VerificationResponse, error) {
	var userVerified entity.UserVerified
	now := time.Now()

	// Cari UserVerified berdasarkan token dan otp
	if err := u.DB.Where("token = ? AND otp = ?", req.Token, req.Otp).First(&userVerified).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, errors.New("Invalid verification token or OTP")
		}
		return nil, err
	}

	// Cek expired
	if now.After(userVerified.ExpiredAt) {
		return nil, errors.New("OTP has expired")
	}

	// ✅ Gunakan tabel AccessDoor (bukan UserMerchant)
	var user entity.AccessDoor
	if err := u.DB.Where("id = ?", userVerified.UserID).First(&user).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, errors.New("User not found")
		}
		return nil, err
	}

	// ✅ Return response dengan data lengkap
	return &dto.VerificationResponse{
		ID:             userVerified.ID,
		UserID:         user.ID,
		DetailID:       user.DetailID,
		FullName:       user.FullName,
		Email:          user.Email,
		TokenVerified:  userVerified.Token,
		Otp:            userVerified.Otp,
		ExpiredAt:      userVerified.ExpiredAt,
		Verified:       false,
		StatusAccount:  user.Suspended,
		AccountExpired: userVerified.AccountExpired,
	}, nil
}
