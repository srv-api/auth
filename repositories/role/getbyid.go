package tax

import (
	dto "github.com/srv-api/detail/dto"
	"github.com/srv-api/detail/entity"
)

func (b *taxRepository) GetById(req dto.GetByIdRequest) (*dto.TaxResponse, error) {
	tr := entity.Tax{
		ID: req.ID,
	}

	if err := b.DB.Where("id = ?", tr.ID).Take(&tr).Error; err != nil {
		return nil, err
	}

	response := &dto.TaxResponse{
		Tax:           tr.Tax,
		TaxPercentage: tr.TaxPercentage,
	}

	return response, nil
}
