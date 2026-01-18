package reset_password

import (
	s "github.com/srv-cashpay/auth/services/auth/reset_password"

	"github.com/labstack/echo/v4"
)

type ResetHandler interface {
	VerifyResetPassword(c echo.Context) error  //reset password
	RequestResetPassword(c echo.Context) error //request reset password
	ResetPassword(c echo.Context) error        //reset password
	ResendVerification(c echo.Context) error   //resend code
}

type resetHandler struct {
	serviceReset s.ResetService
}

func NewResetHandler(service s.ResetService) ResetHandler {
	return &resetHandler{
		serviceReset: service,
	}
}
