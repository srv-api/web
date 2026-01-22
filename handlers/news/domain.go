package news

import (
	s "github.com/srv-api/web/services/news"

	"github.com/labstack/echo/v4"
)

type DomainHandler interface {
	Create(c echo.Context) error
	List(c echo.Context) error
	CreateComment(c echo.Context) error
	Detail(c echo.Context) error
	RedirectIDToSlug(c echo.Context) error
}

type domainHandler struct {
	serviceNews s.NewsService
}

func NewNewsHandler(service s.NewsService) DomainHandler {
	return &domainHandler{
		serviceNews: service,
	}
}
