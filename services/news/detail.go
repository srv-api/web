package news

import "github.com/srv-api/web/entity"

func (s *newsService) Detail(id string) (entity.NewsBlog, error) {
	return s.Repo.Detail(id)
}
