package auth

import (
	"time"

	dto "github.com/srv-cashpay/auth/dto/auth"
	"github.com/srv-cashpay/auth/entity"
	res "github.com/srv-cashpay/util/s/response"
)

func (u *verifyService) VerifyOtpReset(req dto.VerifyResetRequest) (*entity.PasswordResetToken, error) {
	// Use your repository or service to fetch the user by token from the database
	user, err := u.Repo.VerifyOtpReset(req)
	if err != nil {
		// Handle the error (e.g., database query error)
		return nil, err
	}

	// Pemeriksaan waktu kadaluwarsa OTP
	if time.Now().After(user.ExpiredAt) {
		return nil, res.ErrorBuilder(&res.ErrorConstant.ExpiredToken, err)
	}

	return user, nil
}
