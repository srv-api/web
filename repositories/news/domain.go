package news

import (
	dto "github.com/srv-api/web/dto"

	"gorm.io/gorm"
)

type DomainRepository interface {
	Create(req dto.CreateNewsRequest) (dto.CreateNewsResponse, error)
}

type newsRepository struct {
	DB *gorm.DB
}

func NewNewsRepository(DB *gorm.DB) DomainRepository {
	return &newsRepository{
		DB: DB,
	}
}
