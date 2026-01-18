package resi

import (
	dto "github.com/srv-api/web/dto"

	"gorm.io/gorm"
)

type DomainRepository interface {
	Track(courier, awb string) (*dto.TrackingResponse, error)
}

type resiRepository struct {
	DB *gorm.DB
}

func NewResiRepository(DB *gorm.DB) DomainRepository {
	return &resiRepository{
		DB: DB,
	}
}
