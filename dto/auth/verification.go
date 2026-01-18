package auth

// Register
type VerificationRequest struct {
	Otp   string `json:"otp" form:"otp" validate:"required,otp"`
	Token string `query:"token" validate:"required,token"`
}

type ResendVerificationRequest struct {
	Email string `json:"email" validate:"required,email"`
	Otp   string `json:"otp"`
	Token string `query:"token" validate:"required,token"`
}

type ResendVerificationResponse struct {
	Email string `json:"email"`
	Otp   string `json:"otp"`
	Token string `query:"token"`
}

type VerificationResponse struct {
	Otp string `json:"otp" form:"otp" validate:"required,otp"`
}
