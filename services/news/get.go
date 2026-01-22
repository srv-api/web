package news

import (
	"fmt"

	dto "github.com/srv-api/web/dto"
)

func (s *newsService) List() ([]dto.GetNewsListResponse, error) {
	items, err := s.Repo.List()
	if err != nil {
		return nil, err
	}

	var result []dto.GetNewsListResponse
	for _, n := range items {
		result = append(result, dto.GetNewsListResponse{
			ID:        n.ID,
			Title:     n.Title,
			Tag:       n.Tag,
			Excerpt:   n.Excerpt,
			Status:    n.Status,
			FileName:  n.FileName,
			FilePath:  n.FilePath,
			ImageURL:  s.buildImageURL(n.FilePath),
			Slug:      n.Slug,
			CreatedAt: n.CreatedAt,
		})
	}

	return result, nil
}

// helper
func (s *newsService) buildImageURL(path string) string {
	if path == "" {
		return ""
	}
	return fmt.Sprintf("https://api.cashpay.co.id/web/uploads/%s", path)
}
