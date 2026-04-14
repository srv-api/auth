package handlers

import (
	"github.com/labstack/echo/v4"
	dto "github.com/srv-api/auth/dto/auth"
	res "github.com/srv-api/util/s/response"
)

func (b *domainHandler) GetPicture(c echo.Context) error {
	var req dto.GetProfilePictureRequest

	idUint, err := res.IsNumber(c, "file_name")
	if err != nil {
		return res.ErrorBuilder(&res.ErrorConstant.BadRequest, err).Send(c)
	}

	req.FileName = idUint

	transaction, err := b.serviceAuth.GetPicture(req)
	if err != nil {
		return res.ErrorBuilder(&res.ErrorConstant.NotFound, err).Send(c)

	}

	return c.File(transaction.FilePath)

}
