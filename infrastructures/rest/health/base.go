package health

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
	g := c.Resource.Echo()
	g.DELETE("/reset", c.Reset)
	return nil
}
