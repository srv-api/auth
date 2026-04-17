package handlers

import (
	"fmt"
	"log"
	"net/http"

	"github.com/labstack/echo/v4"
	dto "github.com/srv-api/auth/dto/auth"
	res "github.com/srv-api/util/s/response"
)

func (h *domainHandler) UploadImage(c echo.Context) error {
	// Extract ID from path parameter
	id := c.Param("id")
	log.Printf("ID from path: %s", id)

	if id == "" {
		log.Println("Error: ID is empty")
		return res.ErrorBuilder(&res.ErrorConstant.BadRequest,
			fmt.Errorf("id cannot be empty")).Send(c)
	}

	// Extract UserId from the context
	userID, ok := c.Get("UserId").(string)
	if !ok {
		log.Println("Error: UserId not found in context")
		return res.ErrorBuilder(&res.ErrorConstant.InternalServerError, nil).Send(c)
	}
	log.Printf("UserID: %s", userID)

	updatedBy, ok := c.Get("UpdatedBy").(string)
	if !ok {
		log.Println("Error: UpdatedBy not found in context")
		return res.ErrorBuilder(&res.ErrorConstant.InternalServerError, nil).Send(c)
	}
	log.Printf("UpdatedBy: %s", updatedBy)

	DetailID, ok := c.Get("DetailId").(string)
	if !ok {
		log.Println("Error: DetailID not found in context")
		return res.ErrorBuilder(&res.ErrorConstant.InternalServerError, nil).Send(c)
	}
	log.Printf("DetailID: %s", DetailID)

	// Parse file from request
	file, err := c.FormFile("image")
	if err != nil {
		log.Printf("Error getting form file: %v", err)
		return res.ErrorBuilder(&res.ErrorConstant.BadRequest, err).Send(c)
	}
	log.Printf("File received: %s, size: %d", file.Filename, file.Size)

	// Validate file size
	const maxFileSize = 2 * 1024 * 1024
	if file.Size > maxFileSize {
		log.Printf("File too large: %d > %d", file.Size, maxFileSize)
		return c.JSON(http.StatusBadRequest, echo.Map{"message": "File size exceeds the 2MB limit"})
	}

	// Prepare request object
	req := dto.ProfilePictureRequest{
		ID:        id,
		UserID:    userID,
		UpdatedBy: updatedBy,
		CreatedBy: userID, // Add this
		File:      file,
		DetailID:  DetailID,
	}
	log.Printf("Request object prepared: %+v", req)

	// Call service
	log.Println("Calling service.Upload...")
	resp, err := h.serviceAuth.Upload(req)
	if err != nil {
		log.Printf("Service error: %v", err)
		return res.ErrorResponse(err).Send(c)
	}
	log.Printf("Service success, response: %+v", resp)

	return res.SuccessResponse(resp).Send(c)
}
