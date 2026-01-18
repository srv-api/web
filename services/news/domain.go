package news

import (
	m "github.com/srv-api/middlewares/middlewares"
	dto "github.com/srv-api/web/dto"

	r "github.com/srv-api/web/repositories/news"
)

type NewsService interface {
	Create(req dto.CreateNewsRequest) (dto.CreateNewsResponse, error)
}

type newsService struct {
	Repo r.DomainRepository
	jwt  m.JWTService
}

func NewNewsService(Repo r.DomainRepository, jwtS m.JWTService) NewsService {
	return &newsService{
		Repo: Repo,
		jwt:  jwtS,
	}
}
