package entity

type NewsBlog struct {
	ID              string        `gorm:"primary_key;type:varchar(39)" json:"id"`
	UserID          string        `gorm:"type:varchar(36);index,omitempty" json:"user_id"`
	MerchantID      string        `gorm:"type:varchar(36);index" json:"merchant_id"`
	Tag             string        `gorm:"tag,omitempty" json:"tag"`
	Title           string        `gorm:"title,omitempty" json:"title"`
	FileName        string        `gorm:"file_name,omitempty" json:"file_name"`
	FilePath        string        `gorm:"file_path,omitempty" json:"file_path"`
	Body            string        `gorm:"body,omitempty" json:"body"`
	Comments        []NewsComment `gorm:"foreignKey:BlogID"`
	Excerpt         string        `gorm:"excerpt,omitempty" json:"excerpt"`
	Slug            string        `gorm:"slug,omitempty" json:"slug"`
	MetaTitle       string        `gorm:"meta_title,omitempty" json:"meta_title"`
	MetaDescription string        `gorm:"meta_description,omitempty" json:"meta_description"`
	Status          string        `gorm:"status,omitempty" json:"status"`
	CreatedBy       string        `json:"created_by"`
	CreatedAt       string        `json:"created_at"`
}
