package handlers

import (
	"net/http"

	dto "github.com/srv-api/auth/dto/auth"
	util "github.com/srv-api/util/s"
	res "github.com/srv-api/util/s/response"

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
	req := dto.GoogleSignInWebRequest{
		Code: c.QueryParam("code"),
	}

	if req.Code == "" {
		return c.JSON(http.StatusBadRequest, map[string]string{
			"error": "code is missing",
		})
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
