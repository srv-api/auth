package auth

import (
	dto "github.com/srv-cashpay/auth/dto/auth"
	"github.com/srv-cashpay/auth/entity"
)

func (b *authRepository) GetById(req dto.GetByIdRequest) (*dto.GetProfileResponse, error) {
	tr := entity.AccessDoor{
		ID: req.ID,
	}

	if err := b.DB.Where("id = ?", tr.ID).Take(&tr).Error; err != nil {
		return nil, err
	}

	response := &dto.GetProfileResponse{
		FullName: tr.FullName,
		Email:    tr.Email,
		Whatsapp: tr.Whatsapp,
		Password: tr.Password,
	}

	return response, nil
}
