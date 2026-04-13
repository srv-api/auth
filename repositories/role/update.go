package tax

import (
	"github.com/srv-api/detail/dto"
	"github.com/srv-api/detail/entity"
)

func (b *taxRepository) Update(req dto.TaxUpdateRequest) (dto.TaxUpdateResponse, error) {
	// Menyiapkan struktur update untuk produk
	updateTax := entity.Tax{
		Tax:           req.Tax,
		TaxPercentage: req.TaxPercentage,
		Status:        req.Status, // Pastikan status boolean diterima dengan benar
		UpdatedBy:     req.UpdatedBy,
		UserID:        req.UserID,
		Description:   req.Description,
		DetailID:      req.DetailID,
	}

	// Cek apakah produk ada terlebih dahulu
	var existingTax entity.Tax
	err := b.DB.Where("id = ?", req.ID).First(&existingTax).Error
	if err != nil {
		return dto.TaxUpdateResponse{}, err
	}

	// Update produk dengan nilai yang baru
	err = b.DB.Model(&existingTax).Updates(updateTax).Error
	if err != nil {
		return dto.TaxUpdateResponse{}, err
	}

	// Menyiapkan response setelah pembaruan berhasil
	response := dto.TaxUpdateResponse{
		Tax:           updateTax.Tax,
		TaxPercentage: updateTax.TaxPercentage,
		Status:        updateTax.Status,
		UpdatedBy:     updateTax.UpdatedBy,
		UserID:        updateTax.UserID,
		DetailID:      updateTax.DetailID,
		Description:   updateTax.Description,
	}

	return response, nil
}
