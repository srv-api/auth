package auth

import (
	"errors"
	"time"

	res "github.com/srv-api/util/s/response"

	dto "github.com/srv-api/auth/dto/auth"
)

func (u *verifyService) VerifyUserByToken(req dto.VerificationRequest) (*dto.VerificationResponse, error) {
	// Use your repository or service to fetch the user by token from the database
	user, err := u.Repo.VerifyUserByToken(req)
	if err != nil {
		// Handle the error (e.g., database query error)
		return nil, err
	}

	// Pemeriksaan waktu kadaluwarsa OTP
	if time.Now().After(user.ExpiredAt) {
		return nil, res.ErrorBuilder(&res.ErrorConstant.ExpiredToken, err)
	}

	// Simulate updating user verification status (replace with your actual logic)
	user.Verified = true
	if err := u.Repo.UpdateUserVerificationStatus(user); err != nil {
		// Handle the error (e.g., database update error)
		return nil, errors.New("Internal Server Error")
	}

	// ✅ Generate tokens - PASTIKAN user.FullName dan user.Merchant.ID tersedia
	// Jika user dari repository tidak memiliki FullName, Anda perlu mengambil data dari tabel UserMerchant/AccessDoor

	// Opsi 1: Jika user sudah memiliki FullName dan MerchantID
	accesstoken, err := u.jwt.GenerateToken(user.UserID, user.FullName, user.MerchantID)
	if err != nil {
		return nil, errors.New("Failed to generate access token")
	}

	refreshtoken, err := u.jwt.GenerateRefreshToken(user.UserID, user.FullName, user.MerchantID)
	if err != nil {
		return nil, errors.New("Failed to generate refresh token")
	}

	// ✅ Return response dengan tokens
	return &dto.VerificationResponse{
		ID:            user.ID,
		UserID:        user.UserID,
		MerchantID:    user.MerchantID,
		FullName:      user.FullName,
		Email:         user.Email,
		Otp:           user.Otp,
		ExpiredAt:     user.ExpiredAt,
		Verified:      true,
		StatusAccount: user.StatusAccount,
		Token:         accesstoken,
		RefreshToken:  refreshtoken,
		TokenVerified: user.Token,
	}, nil
}
