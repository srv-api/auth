package handlers

import (
	"github.com/labstack/echo/v4"
	dto "github.com/srv-cashpay/auth/dto/auth"
	util "github.com/srv-cashpay/util/s"
	res "github.com/srv-cashpay/util/s/response"
)

func (h *domainHandler) Authenticator(c echo.Context) error {
	var req dto.AuthenticatorRequest
	var resp dto.AuthenticatorResponse

	err := c.Bind(&req)
	if err != nil {
		return res.ErrorBuilder(&res.ErrorConstant.BadRequest, err).Send(c)
	}
	resp, err = h.serviceAuth.Authenticator(req)
	if err != nil {
		if util.IsDuplicateEntryError(err) {
			return res.ErrorResponse(&res.ErrorConstant.Duplicate).Send(c)
		}
		return res.ErrorResponse(err).Send(c)
	}

	return res.SuccessResponse(resp).Send(c)
}
