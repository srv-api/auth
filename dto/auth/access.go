package auth

type AccessRequest struct {
	Access string `param:"access" validate:"required"`
}

type AccessResponse struct {
	Access string `json:"access"`
}

type GetProfilePictureRequest struct {
	FileName string `param:"file_name" validate:"required"`
}

type GetProfilePictureResponse struct {
	FileName string `json:"file_name"`
	FilePath string `json:"file_path"`
}
