package news

import (
	dto "github.com/srv-api/web/dto"
	"github.com/srv-api/web/entity"
)

func (r *newsRepository) Create(req dto.CreateNewsRequest) (dto.CreateNewsResponse, error) {
	create := entity.NewsBlog{
		ID:              req.ID,
		UserID:          req.UserID,
		MerchantID:      req.MerchantID,
		Title:           req.Title,
		Tag:             req.Tag,
		FileName:        req.FileName,
		FilePath:        req.FilePath,
		Body:            req.Body,
		Excerpt:         req.Excerpt,
		Status:          req.Status,
		CreatedBy:       req.CreatedBy,
		Slug:            req.Slug,
		MetaTitle:       req.MetaTitle,
		MetaDescription: req.MetaDescription,
	}

	if err := r.DB.Create(&create).Error; err != nil {
		return dto.CreateNewsResponse{}, err
	}

	return dto.CreateNewsResponse{
		ID:              create.ID,
		UserID:          create.UserID,
		MerchantID:      create.MerchantID,
		Title:           create.Title,
		Tag:             create.Tag,
		FileName:        create.FileName,
		FilePath:        create.FilePath,
		Body:            create.Body,
		Excerpt:         create.Excerpt,
		Status:          create.Status,
		CreatedBy:       create.CreatedBy,
		Slug:            create.Slug,
		MetaTitle:       create.MetaTitle,
		MetaDescription: create.MetaDescription,
	}, nil
}
