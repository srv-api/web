package resi

import (
	s "github.com/srv-api/web/services/resi"

	"github.com/labstack/echo/v4"
)

type TrackingHandler interface {
	Track(c echo.Context) error
}

type trackingHandler struct {
	serviceResi s.ResiService
}

func NewResiHandler(service s.ResiService) TrackingHandler {
	return &trackingHandler{
		serviceResi: service,
	}
}
