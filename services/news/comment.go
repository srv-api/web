package news

import (
	"time"

	"github.com/google/uuid"
	dto "github.com/srv-api/web/dto"
	"github.com/srv-api/web/entity"
)

func (s *newsService) CreateComment(slug string, req dto.CreateCommentRequest) error {
	blog, err := s.Repo.FindBlogBySlug(slug)
	if err != nil {
		return err
	}

	comment := entity.NewsComment{
		ID:        uuid.NewString(),
		BlogID:    blog.ID,
		Name:      req.Name,
		Email:     req.Email,
		Comment:   req.Comment,
		CreatedAt: time.Now(),
	}

	return s.Repo.CreateComment(comment)
}
