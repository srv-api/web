package news

import (
	util "github.com/srv-api/util/s"
	dto "github.com/srv-api/web/dto"
)

func (s *newsService) Create(req dto.CreateNewsRequest) (dto.CreateNewsResponse, error) {
	req.ID = util.GenerateRandomString()

	created, err := s.Repo.Create(req)
	if err != nil {
		return dto.CreateNewsResponse{}, err
	}

	return dto.CreateNewsResponse{
		ID:              created.ID,
		UserID:          created.UserID,
		MerchantID:      created.MerchantID,
		Title:           created.Title,
		Tag:             created.Tag,
		FileName:        created.FileName,
		FilePath:        created.FilePath,
		Body:            created.Body,
		Excerpt:         created.Excerpt,
		Status:          created.Status,
		CreatedBy:       created.CreatedBy,
		Slug:            created.Slug,
		MetaTitle:       created.MetaTitle,
		MetaDescription: created.MetaDescription,
	}, nil
}
