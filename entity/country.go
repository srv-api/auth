package entity

type Country struct {
	ID          uint   `gorm:"primary_key:auto_increment" json:"id"`
	Country     string `gorm:"country" json:"country"`
	CountryCode string `gorm:"country_code" json:"country_code"`
}
