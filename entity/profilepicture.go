package entity

type ProfilePicture struct {
	ID          string `gorm:"primary_key" json:"id"`
	UserID      string `gorm:"type:varchar(36);index" json:"user_id"`
	DetailID    string `gorm:"type:varchar(36);index" json:"detail_id"`
	FileName    string `gorm:"file_name" json:"file_name"`
	FilePath    string `gorm:"file_path" json:"file_path"`
	DataAccount string `gorm:"status,omitempty" json:"data_account"`
	CreatedBy   string `gorm:"created_by" json:"created_by"`
}
