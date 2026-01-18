package handlers

import (
	"github.com/labstack/echo/v4"
	dto "github.com/srv-cashpay/auth/dto/auth"
	util "github.com/srv-cashpay/util/s"
	res "github.com/srv-cashpay/util/s/response"
)

func (h *domainHandler) Signup(c echo.Context) error {
	var req dto.SignupRequest
	var resp dto.SignupResponse

	err := c.Bind(&req)
	if err != nil {
		return res.ErrorBuilder(&res.ErrorConstant.BadRequest, err).Send(c)
	}
	resp, err = h.serviceAuth.Signup(req)
	if err != nil {
		if util.IsDuplicateEntryError(err) {
			return res.ErrorResponse(&res.ErrorConstant.Duplicate).Send(c)
		}
		return res.ErrorResponse(err).Send(c)
	}

	return res.SuccessResponse(resp).Send(c)
}
