package tax

import "github.com/srv-api/merchant/dto"

func (b *taxService) Update(req dto.TaxUpdateRequest) (dto.TaxUpdateResponse, error) {
	request := dto.TaxUpdateRequest{
		Tax:           req.Tax,
		TaxPercentage: req.TaxPercentage,
		Status:        req.Status,
		UpdatedBy:     req.UpdatedBy,
		UserID:        req.UserID,
		Description:   req.Description,
	}

	tax, err := b.Repo.Update(req)
	if err != nil {
		return tax, err
	}

	response := dto.TaxUpdateResponse{
		Tax:           request.Tax,
		TaxPercentage: request.TaxPercentage,
		Status:        request.Status,
		UpdatedBy:     request.UpdatedBy,
		UserID:        request.UserID,
		Description:   request.Description,
	}

	return response, nil
}
