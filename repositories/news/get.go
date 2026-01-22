package news

import "github.com/srv-api/web/entity"

func (r *newsRepository) List() ([]entity.NewsBlog, error) {
	var news []entity.NewsBlog

	err := r.DB.
		Where("status = ?", "published").
		Order("created_at DESC").
		Find(&news).Error

	if err != nil {
		return nil, err
	}

	return news, nil
}
