package news

import (
	"github.com/labstack/echo/v4"
	res "github.com/srv-api/util/s/response"
)

func (h *domainHandler) List(c echo.Context) error {
	news, err := h.serviceNews.List()
	if err != nil {
		return res.ErrorResponse(err).Send(c)
	}

	return res.SuccessResponse(news).Send(c)
}
