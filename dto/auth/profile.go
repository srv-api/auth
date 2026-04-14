package auth

import "mime/multipart"

type ProfileRequest struct {
	ID     string `json:"id"`
	UserID string `json:"user_id"`
	Email  string `json:"email"`
	Gender string `json:"gender"`
}

type ProfileResponse struct {
	ID       string `json:"id"`
	FullName string `json:"full_name"`
	Whatsapp string `json:"whatsapp"`
	Email    string `json:"email"`
	Gender   string `json:"gender"`
}

type ProfilePictureResponse struct {
	FilePath string `json:"file_path"`
}

type ProfilePictureRequest struct {
	ID          string `json:"id"`
	File        *multipart.FileHeader
	CreatedBy   string `json:"created_by"`
	UpdatedBy   string `json:"updated_by"`
	ProductID   string `json:"product_id"`
	UserID      string `json:"user_id"`
	DetailID    string `json:"detail_id"`
	Destination string `json:"destination"`
}

type UpdateProfileRequest struct {
	ID        string `json:"id"`
	FullName  string `json:"full_name"`
	Gender    string `json:"gender"`
	Whatsapp  string `json:"whatsapp"`
	Email     string `json:"email"`
	Password  string `json:"password"`
	UpdatedBy string `json:"updated_by"`
}

type UpdateProfileResponse struct {
	ID        string `json:"id"`
	FullName  string `json:"full_name"`
	Gender    string `json:"gender"`
	Whatsapp  string `json:"whatsapp"`
	Email     string `json:"email"`
	Password  string `json:"password"`
	UpdatedBy string `json:"updated_by"`
}

type GetByIdRequest struct {
	ID string `param:"id" validate:"required"`
}

type GetProfileResponse struct {
	ID        string `json:"id"`
	FullName  string `json:"full_name"`
	Gender    string `json:"gender"`
	Whatsapp  string `json:"whatsapp"`
	Email     string `json:"email"`
	Password  string `json:"password"`
	UpdatedBy string `json:"updated_by"`
}
