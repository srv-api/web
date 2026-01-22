package news

import "github.com/srv-api/web/entity"

func (r *newsRepository) Detail(id string) (entity.NewsBlog, error) {
	var blog entity.NewsBlog

	err := r.DB.
		Preload("Comments").
		First(&blog, "id = ?", id).Error

	return blog, err
}
