package handlers

import (
	"github.com/labstack/echo/v4"
	dto "github.com/srv-cashpay/auth/dto/auth"
	res "github.com/srv-cashpay/util/s/response"
)

func (h *domainHandler) Profile(c echo.Context) error {
	var req dto.ProfileRequest
	var resp dto.ProfileResponse

	userId, ok := c.Get("UserId").(string)
	if !ok {
		return res.ErrorBuilder(&res.ErrorConstant.InternalServerError, nil).Send(c)
	}

	req.UserID = userId

	resp, err := h.serviceAuth.Profile(req)
	if err != nil {
		return res.ErrorBuilder(&res.ErrorConstant.BadRequest, err).Send(c)
	}

	return res.SuccessResponse(resp).Send(c)
}
