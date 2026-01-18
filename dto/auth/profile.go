package auth

type ProfileRequest struct {
	ID     string `json:"id"`
	UserID string `json:"user_id"`
	Email  string `json:"email"`
}

type ProfileResponse struct {
	ID       string `json:"id"`
	FullName string `json:"full_name"`
	Whatsapp string `json:"whatsapp"`
	Email    string `json:"email"`
}

type UpdateProfileRequest struct {
	ID        string `json:"id"`
	FullName  string `json:"full_name"`
	Whatsapp  string `json:"whatsapp"`
	Email     string `json:"email"`
	Password  string `json:"password"`
	UpdatedBy string `json:"updated_by"`
}

type UpdateProfileResponse struct {
	ID        string `json:"id"`
	FullName  string `json:"full_name"`
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
	Whatsapp  string `json:"whatsapp"`
	Email     string `json:"email"`
	Password  string `json:"password"`
	UpdatedBy string `json:"updated_by"`
}
