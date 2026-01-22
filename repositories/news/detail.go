package news

import "github.com/srv-api/web/entity"

func (r *newsRepository) Detail(slug string) (entity.NewsBlog, error) {
	var blog entity.NewsBlog

	err := r.DB.
		Preload("Comments").
		Where("slug = ? AND status = ?", slug, "published").
		First(&blog).Error

	return blog, err
}
