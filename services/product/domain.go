package product

import (
	"github.com/labstack/echo/v4"
	m "github.com/srv-api/middlewares/middlewares"
	dto "github.com/srv-api/web/dto"

	r "github.com/srv-api/web/repositories/product"
)

type ProductService interface {
	Web(context echo.Context, req *dto.Pagination) dto.Response
}

type productService struct {
	Repo r.DomainRepository
	jwt  m.JWTService
}

func NewProductService(Repo r.DomainRepository, jwtS m.JWTService) ProductService {
	return &productService{
		Repo: Repo,
		jwt:  jwtS,
	}
}
