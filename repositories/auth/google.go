package auth

import (
	"errors"
	"time"

	"github.com/srv-cashpay/auth/entity"
	entitymerchant "github.com/srv-cashpay/merchant/entity"
	util "github.com/srv-cashpay/util/s"
	"gorm.io/gorm"
)

func (r *authRepository) FindByEncryptedEmail(encryptedEmail string) (*entity.AccessDoor, error) {
	var user entity.AccessDoor
	if err := r.DB.
		Preload("Merchant").
		Preload("Verified").
		Where("email = ?", encryptedEmail).
		First(&user).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, errors.New("user not found")
		}
		return nil, err
	}
	return &user, nil
}

func (r *authRepository) Create(user *entity.AccessDoor) error {
	if err := r.DB.Create(user).Error; err != nil {
		return err
	}

	// Buat merchant
	merchant := entitymerchant.MerchantDetail{
		ID:         util.GenerateRandomString(),
		UserID:     user.ID,
		CurrencyID: 1,
	}
	if err := r.DB.Save(&merchant).Error; err != nil {
		return err
	}

	// Buat user_verified
	verified := entity.UserVerified{
		ID:        util.GenerateRandomString(),
		UserID:    user.ID,
		Token:     util.GenerateRandomString(),
		ExpiredAt: time.Now().Add(4 * time.Minute),
	}
	if err := r.DB.Create(&verified).Error; err != nil {
		return err
	}

	return nil
}

func (r *authRepository) UpdateWhatsapp(userID string, phone string) error {
	return r.DB.Model(&entity.AccessDoor{}).
		Where("id = ?", userID).
		Update("whatsapp", phone).Error
}
