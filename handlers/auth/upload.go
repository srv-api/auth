package handlers

import (
	"fmt"
	"net/http"

	"github.com/labstack/echo/v4"
	dto "github.com/srv-api/auth/dto/auth"
	res "github.com/srv-api/util/s/response"
)

func (h *domainHandler) UploadImage(c echo.Context) error {
	// Prepare the response
	var resp dto.ProfilePictureResponse

	// Extract UserId from the context
	userID, ok := c.Get("UserId").(string)
	if !ok {
		return res.ErrorBuilder(&res.ErrorConstant.InternalServerError, nil).Send(c)
	}

	updatedBy, ok := c.Get("UpdatedBy").(string)
	if !ok {
		return res.ErrorBuilder(&res.ErrorConstant.InternalServerError, nil).Send(c)
	}

	DetailID, ok := c.Get("DetailID").(string)
	if !ok {
		return res.ErrorBuilder(&res.ErrorConstant.InternalServerError, nil).Send(c)
	}

	id := c.Param("id")
	if id == "" {
		return res.ErrorBuilder(&res.ErrorConstant.BadRequest,
			fmt.Errorf("id cannot be empty")).Send(c)
	}

	// Parse file from request
	file, err := c.FormFile("image") // Ensure "image" matches the form-data field name
	if err != nil {
		return res.ErrorBuilder(&res.ErrorConstant.BadRequest, err).Send(c)
	}

	// Validate file size (2 MB limit)
	const maxFileSize = 2 * 1024 * 1024 // 2 MB
	if file.Size > maxFileSize {
		return c.JSON(http.StatusBadRequest, echo.Map{"message": "File size exceeds the 2MB limit"})
	}

	// Prepare request object
	req := dto.ProfilePictureRequest{
		ID:        id,
		UserID:    userID,
		UpdatedBy: updatedBy,
		File:      file,
		DetailID:  DetailID,
	}

	// Call service to process the upload
	resp, err = h.serviceAuth.Upload(req)
	if err != nil {
		return res.ErrorResponse(err).Send(c)
	}

	// Send success response
	return res.SuccessResponse(resp).Send(c)
}
