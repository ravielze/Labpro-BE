package dorayaki

import (
	"github.com/labstack/echo/v4"
	"github.com/ravielze/Labpro-BE/handlers"
	"github.com/ravielze/Labpro-BE/resources"
	"go.uber.org/dig"
)

type (
	Controller struct {
		dig.In

		Handler  handlers.Handler
		Resource resources.Resource
	}
)

func (c Controller) Register(ec *echo.Group) error {
	g := ec.Group("/dorayaki")
	g.POST("", c.CreateDorayaki)
	g.PATCH("", c.EditDorayaki)
	g.DELETE("/:dorayakiId", c.DeleteDorayaki)

	g2 := ec.Group("/stock")
	g2.GET("/:shopId/:dorayakiId", c.GetStock)
	g2.GET("/:shopId", c.GetAllSoldByShop)
	g2.POST("/transfer", c.TransferStock)
	g2.PATCH("", c.UpdateStock)
	return nil
}
