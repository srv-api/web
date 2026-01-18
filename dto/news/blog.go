package news

type CreateNewsRequest struct {
	ID        string `json:"id"`
	UserID    string `json:"user_id"`
	Tag       string `json:"tag"`
	Title     string `json:"title"`
	File      string `json:"file"`
	Body      string `json:"body"`
	Comment   string `json:"comment"`
	Excerpt   string `json:"excerpt"`
	Status    string `json:"status"`
	CreatedBy string `json:"created_by"`
	CreatedAt string `json:"created_at"`
}

type CreateNewsResponse struct {
	ID        string `json:"id"`
	UserID    string `json:"user_id"`
	Tag       string `json:"tag"`
	Title     string `json:"title"`
	File      string `json:"file"`
	Body      string `json:"body"`
	Comment   string `json:"comment"`
	Excerpt   string `json:"excerpt"`
	Status    string `json:"status"`
	CreatedBy string `json:"created_by"`
	CreatedAt string `json:"created_at"`
}

type UpdateNewsRequest struct {
	ID        string `json:"id"`
	UserID    string `json:"user_id"`
	Tag       string `json:"tag"`
	Title     string `json:"title"`
	File      string `json:"file"`
	Body      string `json:"body"`
	Comment   string `json:"comment"`
	Excerpt   string `json:"excerpt"`
	Status    string `json:"status"`
	CreatedBy string `json:"created_by"`
	CreatedAt string `json:"created_at"`
	UpdatedBy string `json:"updated_by"`
	UpdatedAt string `json:"updated_at"`
}

type UpdateNewsResponse struct {
	ID        string `json:"id"`
	UserID    string `json:"user_id"`
	Tag       string `json:"tag"`
	Title     string `json:"title"`
	File      string `json:"file"`
	Body      string `json:"body"`
	Comment   string `json:"comment"`
	Excerpt   string `json:"excerpt"`
	Status    string `json:"status"`
	CreatedBy string `json:"created_by"`
	CreatedAt string `json:"created_at"`
	UpdatedBy string `json:"updated_by"`
	UpdatedAt string `json:"updated_at"`
}

type GetNewsByIdRequest struct {
	ID string `param:"id" validate:"required"`
}

type GetNewsByIdResponse struct {
	ID        string `json:"id"`
	UserID    string `json:"user_id"`
	Tag       string `json:"tag"`
	Title     string `json:"title"`
	File      string `json:"file"`
	Body      string `json:"body"`
	Comment   string `json:"comment"`
	Excerpt   string `json:"excerpt"`
	Status    string `json:"status"`
	CreatedBy string `json:"created_by"`
	CreatedAt string `json:"created_at"`
	UpdatedBy string `json:"updated_by"`
	UpdatedAt string `json:"updated_at"`
}
