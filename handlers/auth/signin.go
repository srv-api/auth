package handlers

import (
	"strings"

	"github.com/labstack/echo/v4"
	dto "github.com/srv-cashpay/auth/dto/auth"
	res "github.com/srv-cashpay/util/s/response"
)

func (u *domainHandler) Signin(c echo.Context) error {
	var req dto.SigninRequest
	var resp *dto.SigninResponse
	var errResponse error

	err := c.Bind(&req)
	if err != nil {
		return res.ErrorBuilder(&res.ErrorConstant.BadRequest, err).Send(c)
	}

	switch {
	case req.Email != "":
		req.Email = strings.ToLower(req.Email)
		resp, errResponse = u.serviceAuth.Signin(req)
	case req.Whatsapp != "":
		resp, errResponse = u.serviceAuth.SigninByPhoneNumber(req)
	default:
		return res.ErrorBuilder(&res.ErrorConstant.BadRequest, err).Send(c)
	}

	if errResponse != nil {
		return res.ErrorResponse(errResponse).Send(c)
	}

	return res.SuccessResponse(resp).Send(c)
}
