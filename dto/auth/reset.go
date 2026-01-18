package auth

import "time"

type Reset struct {
	Token       string `query:"token" validate:"required,token"`
	NewPassword string `json:"new_password,omitempty" validate:"omitempty"`
}

type Request struct {
	Email string `json:"email" validate:"required,email"`
}

type ResetPasswordRequest struct {
	Email     string    `json:"email" validate:"required,email"`
	UserID    string    `json:"user_id" validate:"required,user_id"`
	Token     string    `json:"token,omitempty" validate:"omitempty"`
	Otp       string    `json:"otp"`
	ExpiredAt time.Time `json:"expired_at"`
}

type ResendResetRequest struct {
	Email string `json:"email" validate:"required,email"`
	Otp   string `json:"otp"`
	Token string `query:"token" validate:"required,token"`
}

type ResetPasswordResponse struct {
	Email     string    `json:"email" validate:"required,email"`
	UserID    string    `json:"user_id" validate:"required,user_id"`
	Token     string    `json:"token,omitempty" validate:"omitempty"`
	Otp       string    `json:"otp"`
	ExpiredAt time.Time `json:"expired_at"`
}

type VerifyResetRequest struct {
	Email     string    `json:"email" validate:"required,email"`
	UserID    string    `json:"user_id" validate:"required,user_id"`
	Token     string    `json:"token,omitempty" validate:"omitempty"`
	Otp       string    `json:"otp"`
	ExpiredAt time.Time `json:"expired_at"`
}

type VerifyResetResponse struct {
	Email     string    `json:"email" validate:"required,email"`
	UserID    string    `json:"user_id" validate:"required,user_id"`
	Token     string    `json:"token,omitempty" validate:"omitempty"`
	Otp       string    `json:"otp"`
	ExpiredAt time.Time `json:"expired_at"`
}
