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
	g.GET("/stock/:shopId/:dorayakiId", c.GetStock)
	g.GET("/stock/:shopId", c.GetAllSoldByShop)
	g.POST("/stock/transfer", c.TransferStock)
	g.PATCH("/stock", c.UpdateStock)
	return nil
}
