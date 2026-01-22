package news

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

func (h *domainHandler) RedirectIDToSlug(c echo.Context) error {
	id := c.Param("id")

	news, err := h.serviceNews.Detail(id)
	if err != nil {
		return c.JSON(http.StatusNotFound, "not found")
	}

	return c.Redirect(
		http.StatusMovedPermanently,
		"/web/news/"+news.Slug,
	)
}
