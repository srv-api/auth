package tax

import (
	"crypto/rand"
	"fmt"
	"strconv"

	dto "github.com/srv-api/merchant/dto"
	"github.com/srv-api/merchant/entity"
)

func (r *taxRepository) Create(req dto.TaxRequest) (dto.TaxResponse, error) {
	// Insert or update the auto_increment value based on merchant_id
	var autoIncrement int
	err := r.DB.Raw(`
		INSERT INTO merchant_auto_increments (merchant_id, next_increment)
		VALUES (?, 1)
		ON CONFLICT (merchant_id) DO UPDATE
		SET next_increment = merchant_auto_increments.next_increment + 1
		RETURNING next_increment - 1;
	`, req.MerchantID).Scan(&autoIncrement).Error

	if err != nil {
		return dto.TaxResponse{}, err
	}

	// Generate Tax ID with prefix and auto increment value
	prefix := "p="
	secureID, err := generateTaxID(prefix, autoIncrement)
	if err != nil {
		return dto.TaxResponse{}, err
	}

	// Create the new tax entry
	create := entity.Tax{
		ID:            secureID,
		Tax:           req.Tax,
		TaxPercentage: req.TaxPercentage,
		Status:        req.Status,
		UserID:        req.UserID,
		MerchantID:    req.MerchantID,
		CreatedBy:     req.CreatedBy,
		Description:   req.Description,
	}

	// Save the new tax to the database
	if err := r.DB.Save(&create).Error; err != nil {
		return dto.TaxResponse{}, err
	}

	// Map the status from integer to string
	statusMap := map[int]string{
		1: "active",
		2: "inactive",
	}

	createdStatus, err := strconv.Atoi(fmt.Sprintf("%v", create.Status))
	if err != nil {
		return dto.TaxResponse{}, fmt.Errorf("invalid status value: %v", create.Status)
	}

	statusString, ok := statusMap[createdStatus]
	if !ok {
		return dto.TaxResponse{}, fmt.Errorf("invalid status value in database")
	}

	// Build the response for the created tax
	response := dto.TaxResponse{
		ID:            create.ID,
		Tax:           create.Tax,
		TaxPercentage: create.TaxPercentage,
		UserID:        create.UserID,
		MerchantID:    create.MerchantID,
		Description:   create.Description,
		Status:        statusString,
		CreatedBy:     create.CreatedBy,
	}

	return response, nil
}

// Function to generate the tax ID
func generateTaxID(prefix string, autoIncrement int) (string, error) {
	// Format auto-increment value as a 5-digit string
	autoIncStr := fmt.Sprintf("%05d", autoIncrement)

	// Generate a secure random part of the tax ID
	securePart, err := generateSecurePart()
	if err != nil {
		return "", err
	}

	// Combine the prefix, secure random part, and auto-increment value to form the final tax ID
	return fmt.Sprintf("%s%s%s", prefix, securePart, autoIncStr), nil
}

// Function to generate a secure random part for the tax ID
func generateSecurePart() (string, error) {
	const chars = "ABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789abcdefghijklmnopqrstuvwxyz"

	// Generate a random string of length 12
	securePart := make([]byte, 12)
	_, err := rand.Read(securePart)
	if err != nil {
		return "", err
	}

	// Map each byte to a character from the chars string
	for i := range securePart {
		securePart[i] = chars[securePart[i]%byte(len(chars))]
	}

	return string(securePart), nil
}
