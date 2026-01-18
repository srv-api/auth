package entity

type Authenticator struct {
	ID       string `gorm:"primary_key" json:"id"`
	Status   string `gorm:"status" json:"status"`
	TokenApp string `gorm:"token_app" json:"token_app"`
	Otp      string `gorm:"otp" json:"otp"`
}
