package auth

import (
	dto "github.com/srv-cashpay/auth/dto/auth"
	"github.com/srv-cashpay/auth/entity"
)

func (u *authRepository) UpdateProfile(req dto.UpdateProfileRequest) (dto.UpdateProfileResponse, error) {
	tr := dto.GetByIdRequest{
		ID: req.ID,
	}

	request := entity.AccessDoor{
		Password: req.Password,
		FullName: req.FullName,
		Email:    req.Email,
		Whatsapp: req.Whatsapp,
		ID:       tr.ID,
	}

	mer, err := u.GetById(tr)
	if err != nil {
		return dto.UpdateProfileResponse{}, err
	}

	err = u.DB.Where("ID = ?", req.ID).Updates(entity.AccessDoor{
		FullName: request.FullName,
		Email:    request.Email,
		Whatsapp: request.Whatsapp,
		Password: request.Password,
	}).Error
	if err != nil {
		return dto.UpdateProfileResponse{}, err
	}

	response := dto.UpdateProfileResponse{
		FullName: request.FullName,
		Email:    request.Email,
		Whatsapp: request.Whatsapp,
		Password: request.Password,
		ID:       mer.ID,
	}

	return response, nil
}
