package reset_password

import (
	res "github.com/srv-cashpay/util/s/response"

	dto "github.com/srv-cashpay/auth/dto/auth"

	"github.com/labstack/echo/v4"
)

func (u *resetHandler) ResendVerification(c echo.Context) error {
	var req dto.ResendResetRequest

	err := c.Bind(&req)
	if err != nil {
		return res.ErrorBuilder(&res.ErrorConstant.BadRequest, err).Send(c)
	}

	token := c.QueryParam("token")
	req.Token = token

	data, err := u.serviceReset.ResendReset(req)
	if err != nil {
		return c.HTML(400, "Verification failed: "+err.Error())
	}

	return res.SuccessResponse(data).Send(c)
}
