package tax

import (
	dto "github.com/srv-api/merchant/dto"

	"gorm.io/gorm"
)

type DomainRepository interface {
	Create(req dto.TaxRequest) (dto.TaxResponse, error)
	Get(req *dto.Pagination) (RepositoryResult, int)
	GetById(req dto.GetByIdRequest) (*dto.TaxResponse, error)
	Delete(req dto.DeleteRequest) (dto.DeleteResponse, error)
	BulkDelete(req dto.BulkDeleteRequest) (int, error)
	Update(req dto.TaxUpdateRequest) (dto.TaxUpdateResponse, error)
}

type taxRepository struct {
	DB *gorm.DB
}

func NewTaxRepository(DB *gorm.DB) DomainRepository {
	return &taxRepository{
		DB: DB,
	}
}
