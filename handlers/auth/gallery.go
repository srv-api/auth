// handlers/auth_handler.go
package handlers

import (
	"fmt"
	"log"

	"github.com/labstack/echo/v4"
	dto "github.com/srv-api/auth/dto/auth"
	res "github.com/srv-api/util/s/response"
)

func (h *domainHandler) Gallery(c echo.Context) error {
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

	detailID, ok := c.Get("DetailId").(string)
	if !ok {
		log.Println("Error: DetailID not found in context")
		return res.ErrorBuilder(&res.ErrorConstant.InternalServerError, nil).Send(c)
	}
	log.Printf("DetailID: %s", detailID)

	// Parse multiple files from request
	form, err := c.MultipartForm()
	if err != nil {
		log.Printf("Error getting multipart form: %v", err)
		return res.ErrorBuilder(&res.ErrorConstant.BadRequest, err).Send(c)
	}

	files := form.File["images"] // Field name "images" untuk multiple files
	if len(files) == 0 {
		log.Println("No files uploaded")
		return res.ErrorBuilder(&res.ErrorConstant.BadRequest,
			fmt.Errorf("at least one image is required")).Send(c)
	}

	log.Printf("Number of files received: %d", len(files))

	// Validate max files (max 10 files per upload)
	const maxFiles = 10
	if len(files) > maxFiles {
		return res.ErrorBuilder(&res.ErrorConstant.BadRequest,
			fmt.Errorf("maximum %d files allowed per upload", maxFiles)).Send(c)
	}

	// Prepare request object
	req := dto.GalleryUploadRequest{
		UserID:    userID,
		DetailID:  detailID,
		CreatedBy: userID,
		UpdatedBy: updatedBy,
		Files:     files,
	}

	// Call service
	log.Println("Calling service.Gallery (multiple upload)...")
	resp, err := h.serviceAuth.Gallery(req)
	if err != nil {
		log.Printf("Service error: %v", err)
		return res.ErrorResponse(err).Send(c)
	}

	// Return appropriate response based on upload results
	if err != nil {
		log.Printf("Service error: %v", err)
		return res.ErrorResponse(err).Send(c)
	}
	// All success
	return res.SuccessResponse(resp).Send(c)
}

// GetGallery handler untuk mengambil semua gallery user
func (h *domainHandler) GetGallery(c echo.Context) error {
	userID, ok := c.Get("UserId").(string)
	if !ok {
		return res.ErrorBuilder(&res.ErrorConstant.InternalServerError, nil).Send(c)
	}

	gallery, err := h.serviceAuth.GetUserGallery(userID)
	if err != nil {
		log.Printf("Failed to get gallery: %v", err)
		return res.ErrorResponse(err).Send(c)
	}

	return res.SuccessResponse(gallery).Send(c)
}

// DeleteGallery handler untuk hapus file gallery
func (h *domainHandler) DeleteGallery(c echo.Context) error {
	fileID := c.Param("id")
	if fileID == "" {
		return res.ErrorBuilder(&res.ErrorConstant.BadRequest,
			fmt.Errorf("file id is required")).Send(c)
	}

	userID, ok := c.Get("UserId").(string)
	if !ok {
		return res.ErrorBuilder(&res.ErrorConstant.InternalServerError, nil).Send(c)
	}

	err := h.serviceAuth.DeleteGalleryFile(fileID, userID)
	if err != nil {
		log.Printf("Failed to delete gallery: %v", err)
		return res.ErrorResponse(err).Send(c)
	}

	return res.SuccessResponse(map[string]string{"message": "Gallery deleted successfully"}).Send(c)
}
