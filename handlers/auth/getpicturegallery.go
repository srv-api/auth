package handlers

import (
	"github.com/labstack/echo/v4"
	dto "github.com/srv-api/auth/dto/auth"
	res "github.com/srv-api/util/s/response"
)

func (b *domainHandler) GetPictureGallery(c echo.Context) error {
	var req dto.GetPictureGalleryRequest

	idUint, err := res.IsNumber(c, "file_name")
	if err != nil {
		return res.ErrorBuilder(&res.ErrorConstant.BadRequest, err).Send(c)
	}

	req.FilePath = idUint

	transaction, err := b.serviceAuth.GetPictureGallery(req)
	if err != nil {
		return res.ErrorBuilder(&res.ErrorConstant.NotFound, err).Send(c)

	}

	return c.File(transaction.FilePath)

}
