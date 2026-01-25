package product

import (
	"sync"

	dto "github.com/srv-api/web/dto"

	"gorm.io/gorm"
)

type DomainRepository interface {
	Web(req *dto.Pagination) (RepositoryResult, int)
}

type productRepository struct {
	DB *gorm.DB
	mu sync.Mutex
}

func NewProductRepository(DB *gorm.DB) DomainRepository {
	return &productRepository{
		DB: DB,
	}
}
