package news

import (
	util "github.com/srv-api/util/s"
	dto "github.com/srv-api/web/dto"
)

func (s *newsService) Create(req dto.CreateNewsRequest) (dto.CreateNewsResponse, error) {
	create := dto.CreateNewsRequest{
		ID:         util.GenerateRandomString(),
		UserID:     req.UserID,
		MerchantID: req.MerchantID,
		Title:      req.Title,
		Tag:        req.Tag,
		File:       req.File,
		Body:       req.Body,
		Excerpt:    req.Excerpt,
		Status:     req.Status,
		CreatedBy:  req.CreatedBy,
	}

	created, err := s.Repo.Create(create)
	if err != nil {
		return dto.CreateNewsResponse{}, err
	}

	response := dto.CreateNewsResponse{
		ID:         created.ID,
		UserID:     created.UserID,
		MerchantID: created.MerchantID,
		Title:      created.Title,
		Tag:        created.Tag,
		File:       created.File,
		Body:       created.Body,
		Excerpt:    created.Excerpt,
		Status:     created.Status,
		CreatedBy:  created.CreatedBy,
	}

	return response, nil
}
