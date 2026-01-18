package auth

import (
	"errors"
	"time"

	dto "github.com/srv-cashpay/auth/dto/auth"
	"github.com/srv-cashpay/auth/entity"
	res "github.com/srv-cashpay/util/s/response"
)

func (u *verifyService) VerifyUserByToken(req dto.VerificationRequest) (*entity.UserVerified, error) {
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

	return user, nil
}
