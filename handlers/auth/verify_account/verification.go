package verify_account

import (
	dto "github.com/srv-cashpay/auth/dto/auth"
	res "github.com/srv-cashpay/util/s/response"

	"github.com/labstack/echo/v4"
)

func (u *verifyHandler) HandleVerification(c echo.Context) error {
	var req dto.VerificationRequest

	err := c.Bind(&req)
	if err != nil {
		return res.ErrorBuilder(&res.ErrorConstant.BadRequest, err).Send(c)
	}

	token := c.QueryParam("token")
	req.Token = token

	data, err := u.serviceVerify.VerifyUserByToken(req)
	if err != nil {
		return res.ErrorResponse(err).Send(c)
	}

	return res.SuccessResponse(data).Send(c)
}
