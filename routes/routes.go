package routes

import (
	"github.com/labstack/echo/v4"
	"github.com/srv-api/middlewares/middlewares"
	"github.com/srv-api/web/configs"
	h_news "github.com/srv-api/web/handlers/news"
	r_news "github.com/srv-api/web/repositories/news"
	s_news "github.com/srv-api/web/services/news"

	h_track "github.com/srv-api/web/handlers/resi"
	r_track "github.com/srv-api/web/repositories/resi"
	s_track "github.com/srv-api/web/services/resi"
)

var (
	DB    = configs.InitDB()
	JWT   = middlewares.NewJWTService()
	newsR = r_news.NewNewsRepository(DB)
	newsS = s_news.NewNewsService(newsR, JWT)
	newsH = h_news.NewNewsHandler(newsS)

	trackR = r_track.NewResiRepository(DB)
	trackS = s_track.NewResiService(trackR, JWT)
	trackH = h_track.NewResiHandler(trackS)
)

func New() *echo.Echo {
	e := echo.New()

	news := e.Group("/news", middlewares.AuthorizeJWT(JWT))
	{
		news.POST("/create/news", newsH.Create)
	}
	track := e.Group("/track")
	{
		track.GET("/jne", trackH.Track)
	}

	return e
}
