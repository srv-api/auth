package tax

import (
	dto "github.com/srv-api/merchant/dto"
	"github.com/srv-api/merchant/entity"
)

func (b *taxRepository) Delete(req dto.DeleteRequest) (dto.DeleteResponse, error) {
	tr := dto.GetByIdRequest{
		ID: req.ID,
	}

	_, err := b.GetById(tr)
	if err != nil {
		return dto.DeleteResponse{}, err
	}

	// Use GORM BeforeDelete hook to set DeletedBy
	if err := b.DB.Where("id = ?", req.ID).Delete(&entity.Tax{}).Error; err != nil {
		return dto.DeleteResponse{}, err
	}

	response := dto.DeleteResponse{
		DeletedBy: req.DeletedBy,
	}

	return response, nil
}
