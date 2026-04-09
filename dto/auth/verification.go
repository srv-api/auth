package auth

import "time"

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
	ID             string    `json:"id"`
	UserID         string    `json:"user_id"`
	MerchantID     string    `json:"merchant_id,omitempty"`
	FullName       string    `json:"full_name,omitempty"`
	Email          string    `json:"email,omitempty"`
	Otp            string    `json:"otp,omitempty"`
	ExpiredAt      time.Time `json:"expired_at,omitempty"`
	Verified       bool      `json:"verified"`
	StatusAccount  bool      `json:"status_account"`
	AccountExpired time.Time `json:"account_expired,omitempty"`
	AccessToken    string    `json:"access_token"`
	RefreshToken   string    `json:"refresh_token"`
	TokenVerified  string    `json:"token_verified,omitempty"`
}
