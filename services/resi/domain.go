package resi

import (
	m "github.com/srv-api/middlewares/middlewares"
	dto "github.com/srv-api/web/dto"

	r "github.com/srv-api/web/repositories/resi"
)

type ResiService interface {
	TrackPackage(courier, awb string) (*dto.TrackingResponse, error)
}

type resiService struct {
	Repo r.DomainRepository
	jwt  m.JWTService
}

func NewResiService(Repo r.DomainRepository, jwtS m.JWTService) ResiService {
	return &resiService{
		Repo: Repo,
		jwt:  jwtS,
	}
}
