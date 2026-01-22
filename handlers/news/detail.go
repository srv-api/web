package news

import (
	"github.com/labstack/echo/v4"
	res "github.com/srv-api/util/s/response"
)

func (h *domainHandler) Detail(c echo.Context) error {
	slug := c.Param("slug")

	data, err := h.serviceNews.Detail(slug)
	if err != nil {
		return res.ErrorResponse(err).Send(c)
	}

	return res.SuccessResponse(data).Send(c)
}
