package auth

type AuthenticatorRequest struct {
	ID       string `json:"id"`
	Status   string `json:"status" form:"status" validate:"required,status"`
	TokenApp string `json:"token_app"`
	Otp      string `json:"otp" form:"otp" validate:"required,otp"`
}

type AuthenticatorResponse struct {
	ID       string `json:"id"`
	Status   string `json:"status" form:"status" validate:"required,status"`
	TokenApp string `json:"token_app"`
	Otp      string `json:"otp" form:"otp" validate:"required,otp"`
}
