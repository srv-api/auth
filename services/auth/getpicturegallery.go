package auth

import (
	"fmt"
	"os"

	dto "github.com/srv-api/auth/dto/auth"
)

func (b *authService) GetPictureGallery(req dto.GetPictureGalleryRequest) (*dto.GetGalleryResponse, error) {
	// Ambil data dari repository
	transaction, err := b.Repo.GetPictureGallery(req)
	if err != nil {
		return nil, err
	}

	// Pastikan path file benar
	filePath := "./" + transaction.FilePath // Tambahkan prefix untuk path lokal
	if _, err := os.Stat(filePath); os.IsNotExist(err) {
		return nil, fmt.Errorf("file not found")
	}

	transaction.FilePath = filePath
	return transaction, nil
}
