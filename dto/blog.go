package dto

type CreateNewsRequest struct {
	ID              string `json:"id"`
	UserID          string `json:"user_id"`
	MerchantID      string `json:"merchant_id"`
	Tag             string `json:"tag"`
	Title           string `json:"title"`
	FileName        string `json:"file_name"`
	FilePath        string `json:"file_path"`
	Body            string `json:"body"`
	Comment         string `json:"comment"`
	Excerpt         string `json:"excerpt"`
	Slug            string `json:"slug"`
	MetaTitle       string `json:"meta_title"`
	MetaDescription string `json:"meta_description"`
	Status          string `json:"status"`
	CreatedBy       string `json:"created_by"`
	CreatedAt       string `json:"created_at"`
}

type CreateNewsResponse struct {
	ID              string `json:"id"`
	UserID          string `json:"user_id"`
	MerchantID      string `json:"merchant_id"`
	Tag             string `json:"tag"`
	Title           string `json:"title"`
	FileName        string `json:"file_name"`
	FilePath        string `json:"file_path"`
	Body            string `json:"body"`
	Comment         string `json:"comment"`
	Excerpt         string `json:"excerpt"`
	Status          string `json:"status"`
	Slug            string `json:"slug"`
	MetaTitle       string `json:"meta_title"`
	MetaDescription string `json:"meta_description"`
	CreatedBy       string `json:"created_by"`
	CreatedAt       string `json:"created_at"`
}

type UpdateNewsRequest struct {
	ID         string `json:"id"`
	UserID     string `json:"user_id"`
	MerchantID string `json:"merchant_id"`
	Tag        string `json:"tag"`
	Title      string `json:"title"`
	File       string `json:"file"`
	Body       string `json:"body"`
	Comment    string `json:"comment"`
	Excerpt    string `json:"excerpt"`
	Status     string `json:"status"`
	CreatedBy  string `json:"created_by"`
	CreatedAt  string `json:"created_at"`
	UpdatedBy  string `json:"updated_by"`
	UpdatedAt  string `json:"updated_at"`
}

type UpdateNewsResponse struct {
	ID         string `json:"id"`
	UserID     string `json:"user_id"`
	MerchantID string `json:"merchant_id"`
	Tag        string `json:"tag"`
	Title      string `json:"title"`
	File       string `json:"file"`
	Body       string `json:"body"`
	Comment    string `json:"comment"`
	Excerpt    string `json:"excerpt"`
	Status     string `json:"status"`
	CreatedBy  string `json:"created_by"`
	CreatedAt  string `json:"created_at"`
	UpdatedBy  string `json:"updated_by"`
	UpdatedAt  string `json:"updated_at"`
}

type GetNewsByIdRequest struct {
	ID string `param:"id" validate:"required"`
}

type GetNewsByIdResponse struct {
	ID         string `json:"id"`
	UserID     string `json:"user_id"`
	MerchantID string `json:"merchant_id"`
	Tag        string `json:"tag"`
	Title      string `json:"title"`
	File       string `json:"file"`
	Body       string `json:"body"`
	Comment    string `json:"comment"`
	Excerpt    string `json:"excerpt"`
	Status     string `json:"status"`
	CreatedBy  string `json:"created_by"`
	CreatedAt  string `json:"created_at"`
	UpdatedBy  string `json:"updated_by"`
	UpdatedAt  string `json:"updated_at"`
}

type GetNewsListResponse struct {
	ID        string `json:"id"`
	Title     string `json:"title"`
	Tag       string `json:"tag"`
	Excerpt   string `json:"excerpt"`
	Status    string `json:"status"`
	FileName  string `json:"file_name"`
	FilePath  string `json:"file_path"`
	ImageURL  string `json:"image_url"`
	Slug      string `json:"slug"`
	CreatedAt string `json:"created_at"`
}
