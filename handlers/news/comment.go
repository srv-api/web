package news

import (
	"github.com/labstack/echo/v4"
	res "github.com/srv-api/util/s/response"
	dto "github.com/srv-api/web/dto"
)

func (h *domainHandler) CreateComment(c echo.Context) error {
	slug := c.Param("slug")

	var req dto.CreateCommentRequest
	if err := c.Bind(&req); err != nil {
		return res.ErrorBuilder(&res.ErrorConstant.BadRequest, err).Send(c)
	}

	if err := h.serviceNews.CreateComment(slug, req); err != nil {
		return res.ErrorResponse(err).Send(c)
	}

	return res.SuccessResponse("comment added").Send(c)
}
