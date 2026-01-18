package auth

import (
	"time"

	dto "github.com/srv-cashpay/auth/dto/auth"
	"github.com/srv-cashpay/auth/entity"
	entitymerchant "github.com/srv-cashpay/merchant/entity"
	util "github.com/srv-cashpay/util/s"
)

func (r *authRepository) Signup(req dto.SignupRequest) (dto.SignupResponse, error) {

	user := entity.AccessDoor{
		ID:           req.ID,
		FullName:     req.FullName,
		Whatsapp:     req.Whatsapp,
		Email:        req.Email,
		Password:     req.Password,
		AccessRoleID: req.AccessRoleID,
		MerchantID:   util.GenerateRandomString(),
	}

	if err := r.DB.Save(&user).Error; err != nil {
		return dto.SignupResponse{}, err
	}

	merchant := entitymerchant.MerchantDetail{
		ID:         user.MerchantID,
		UserID:     user.ID,
		CurrencyID: 1,
	}

	if err := r.DB.Save(&merchant).First(&merchant).Error; err != nil {
		return dto.SignupResponse{}, err
	}

	verified := entity.UserVerified{
		ID:        util.GenerateRandomString(),
		UserID:    user.ID,
		Token:     req.Token,
		Otp:       req.Otp,
		ExpiredAt: time.Now().Add(4 * time.Minute),
	}

	if err := r.DB.Save(&verified).First(&verified).Error; err != nil {
		return dto.SignupResponse{}, err
	}
	response := dto.SignupResponse{
		ID:           user.ID,
		FullName:     user.FullName,
		Whatsapp:     user.Whatsapp,
		Email:        user.Email,
		Password:     user.Password,
		Token:        verified.Token,
		AccessRoleID: user.AccessRoleID,
	}

	return response, nil
}
