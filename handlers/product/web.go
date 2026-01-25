package product

import (
	"github.com/labstack/echo/v4"
	res "github.com/srv-api/util/s/response"
	"github.com/srv-api/web/helpers"
)

func (b *domainHandler) Web(c echo.Context) error {
	paginationDTO := helpers.GeneratePaginationRequest(c)

	merchant_slug := c.Param("merchant_slug")
	if merchant_slug == "" {
		return res.ErrorBuilder(&res.ErrorConstant.InternalServerError, nil).Send(c)
	}
	paginationDTO.MerchantSlug = merchant_slug

	if err := c.Bind(&paginationDTO); err != nil {
		return c.JSON(400, "Invalid request")
	}

	users := b.serviceProduct.Web(c, paginationDTO)

	return c.JSON(200, users)
}
