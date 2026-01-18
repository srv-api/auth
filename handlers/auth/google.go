package handlers

import (
	"net/http"

	dto "github.com/srv-cashpay/auth/dto/auth"
	util "github.com/srv-cashpay/util/s"
	res "github.com/srv-cashpay/util/s/response"

	"github.com/labstack/echo/v4"
)

func (h *domainHandler) GoogleSignIn(c echo.Context) error {
	var req dto.GoogleSignInRequest
	if err := c.Bind(&req); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "Invalid request"})
	}

	resp, err := h.serviceAuth.SignInWithGoogle(req)
	if err != nil {
		if util.IsDuplicateEntryError(err) {
			return res.ErrorResponse(&res.ErrorConstant.Duplicate).Send(c)
		}
		return res.ErrorResponse(err).Send(c)
	}
	return res.SuccessResponse(resp).Send(c)

}

func (h *domainHandler) GoogleSignInWeb(c echo.Context) error {
	var req dto.GoogleSignInWebRequest
	if err := c.Bind(&req); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "Invalid request"})
	}

	resp, err := h.serviceAuth.SignInWithGoogleWeb(req)
	if err != nil {
		if util.IsDuplicateEntryError(err) {
			return res.ErrorResponse(&res.ErrorConstant.Duplicate).Send(c)
		}
		return res.ErrorResponse(err).Send(c)
	}
	return res.SuccessResponse(resp).Send(c)
}
