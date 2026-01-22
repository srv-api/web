package news

import (
	"github.com/google/uuid"
	dto "github.com/srv-api/web/dto"
	"github.com/srv-api/web/entity"
)

func (s *newsService) CreateComment(blogID string, req dto.CreateCommentRequest) error {
	comment := entity.NewsComment{
		ID:      uuid.NewString(),
		BlogID:  blogID,
		Name:    req.Name,
		Email:   req.Email,
		Comment: req.Comment,
	}

	return s.Repo.CreateComment(comment)
}
