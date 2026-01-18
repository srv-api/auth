package tax

import (
	"github.com/labstack/echo/v4"
	dto "github.com/srv-api/merchant/dto"
	m "github.com/srv-api/middlewares/middlewares"

	r "github.com/srv-api/merchant/repositories/tax"
)

type TaxService interface {
	Create(req dto.TaxRequest) (dto.TaxResponse, error)
	Get(context echo.Context, req *dto.Pagination) dto.Response
	GetById(req dto.GetByIdRequest) (*dto.TaxResponse, error)
	Delete(req dto.DeleteRequest) (dto.DeleteResponse, error)
	BulkDelete(req dto.BulkDeleteRequest) (dto.BulkDeleteResponse, error)
	Update(req dto.TaxUpdateRequest) (dto.TaxUpdateResponse, error)
}

type taxService struct {
	Repo r.DomainRepository
	jwt  m.JWTService
}

func NewTaxService(Repo r.DomainRepository, jwtS m.JWTService) TaxService {
	return &taxService{
		Repo: Repo,
		jwt:  jwtS,
	}
}
