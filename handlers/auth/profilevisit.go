package handlers

import (
	"github.com/labstack/echo/v4"
	dto "github.com/srv-api/auth/dto/auth"
	res "github.com/srv-api/util/s/response"
)

func (h *domainHandler) ProfileVisit(c echo.Context) error {
	var req dto.ProfileVisitRequest
	var resp dto.ProfileResponse

	idUint, err := res.IsNumber(c, "id")
	if err != nil {
		return res.ErrorBuilder(&res.ErrorConstant.BadRequest, err).Send(c)
	}
	req.ID = idUint

	userId, ok := c.Get("UserId").(string)
	if !ok {
		return res.ErrorBuilder(&res.ErrorConstant.InternalServerError, nil).Send(c)
	}

	req.UserID = userId

	resp, error := h.serviceAuth.ProfileVisit(req)
	if error != nil {
		return res.ErrorBuilder(&res.ErrorConstant.BadRequest, error).Send(c)
	}

	return res.SuccessResponse(resp).Send(c)
}
