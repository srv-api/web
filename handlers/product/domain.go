package product

import (
	s "github.com/srv-api/web/services/product"

	"github.com/labstack/echo/v4"
)

type DomainHandler interface {
	Web(c echo.Context) error
}

type domainHandler struct {
	serviceProduct s.ProductService
}

func NewProductHandler(service s.ProductService) DomainHandler {
	return &domainHandler{
		serviceProduct: service,
	}
}
