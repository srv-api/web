package news

import "github.com/srv-api/web/entity"

func (r *newsRepository) CreateComment(comment entity.NewsComment) error {
	return r.DB.Create(&comment).Error
}

func (r *newsRepository) FindBlogBySlug(slug string) (entity.NewsBlog, error) {
	var blog entity.NewsBlog
	err := r.DB.Where("slug = ?", slug).First(&blog).Error
	return blog, err
}
