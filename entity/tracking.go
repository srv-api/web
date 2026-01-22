package entity

type Tracking struct {
	ID         string `gorm:"primary_key;type:varchar(39)" json:"id"`
	IP         string `gorm:"type:varchar(36);index,omitempty" json:"ip"`
	MerchantID string `gorm:"type:varchar(36);index" json:"merchant_id"`
	Tag        string `gorm:"tag,omitempty" json:"tag"`
	Title      string `gorm:"title,omitempty" json:"title"`
	FileName   string `gorm:"file_name,omitempty" json:"file_name"`
	FilePath   string `gorm:"file_path,omitempty" json:"file_path"`
	Body       string `gorm:"body,omitempty" json:"body"`
	Comment    string `gorm:"comment,omitempty" json:"comment"`
	Excerpt    string `gorm:"excerpt,omitempty" json:"excerpt"`
	Status     string `gorm:"status,omitempty" json:"status"`
	CreatedBy  string `json:"created_by"`
	CreatedAt  string `json:"created_at"`
}
