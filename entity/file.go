package entity

type Galery struct {
	File []File
}

type File struct {
	ID          string `gorm:"primary_key,omitempty" json:"id"`
	UserID      string `gorm:"type:varchar(36);index,omitempty" json:"user_id"`
	FileName    string `gorm:"file_name,omitempty" json:"file_name"`
	FilePath    string `gorm:"file_path,omitempty" json:"file_path"`
	DataAccount string `gorm:"status,omitempty" json:"data_account"`
}
