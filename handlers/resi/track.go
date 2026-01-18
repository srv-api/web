package resi

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

func (h *trackingHandler) Track(c echo.Context) error {
	courier := c.QueryParam("courier")
	awb := c.QueryParam("awb")

	if courier == "" || awb == "" {
		return c.JSON(http.StatusBadRequest, echo.Map{
			"message": "courier and awb required",
		})
	}

	data, err := h.serviceResi.TrackPackage(courier, awb)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{
			"message": err.Error(),
		})
	}

	return c.JSON(http.StatusOK, data)
}
