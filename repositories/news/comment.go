package news

import "github.com/srv-api/web/entity"

func (r *newsRepository) CreateComment(comment entity.NewsComment) error {
	return r.DB.Create(&comment).Error
}
