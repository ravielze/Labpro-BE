package shop

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
	g := ec.Group("/shop")
	g.POST("", c.Create)
	g.DELETE("/:id", c.Delete)
	g.GET("/:id", c.Get)
	g.GET("", c.GetAll)
	g.PATCH("", c.Update)
	return nil
}
