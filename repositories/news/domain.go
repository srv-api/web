package news

import (
	dto "github.com/srv-api/web/dto"
	"github.com/srv-api/web/entity"

	"gorm.io/gorm"
)

type DomainRepository interface {
	Create(req dto.CreateNewsRequest) (dto.CreateNewsResponse, error)
	List() ([]entity.NewsBlog, error)
}

type newsRepository struct {
	DB *gorm.DB
}

func NewNewsRepository(DB *gorm.DB) DomainRepository {
	return &newsRepository{
		DB: DB,
	}
}
