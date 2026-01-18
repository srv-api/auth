package tax

import (
	dto "github.com/srv-api/merchant/dto"
)

func (b *taxService) GetById(req dto.GetByIdRequest) (*dto.TaxResponse, error) {
	transaction, err := b.Repo.GetById(req)
	if err != nil {
		return nil, err
	}

	return transaction, nil
}
