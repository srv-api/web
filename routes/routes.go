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

	h_product "github.com/srv-api/web/handlers/product"
	r_product "github.com/srv-api/web/repositories/product"
	s_product "github.com/srv-api/web/services/product"
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

	productR = r_product.NewProductRepository(DB)
	productS = s_product.NewProductService(productR, JWT)
	productH = h_product.NewProductHandler(productS)
)

func New() *echo.Echo {
	e := echo.New()
	e.Static("/web/uploads", "./uploads")
	e.GET("/web/:merchant_slug", productH.Web)

	news := e.Group("/web", middlewares.AuthorizeJWT(JWT))
	{
		news.POST("/create/news", newsH.Create)
	}

	web := e.Group("/web")
	{
		web.GET("/track/line", trackH.Track)
		web.GET("/list/news", newsH.List)
		web.POST("/news/:slug/comment", newsH.CreateComment)
		web.GET("/news/:slug", newsH.Detail)
		web.GET("/news/id/:id", newsH.RedirectIDToSlug)
	}

	return e
}
