package health

import (
	"github.com/labstack/echo/v4"
	"github.com/ravielze/Labpro-BE/constants"
	oculiContext "github.com/ravielze/oculi/context"
	request "github.com/ravielze/oculi/request/echo"
)

func (c *Controller) Check(ec echo.Context) error {
	ctx := ec.(*oculiContext.Context)
	req := request.New(ctx, c.Resource.Database)

	result := ctx.Process(
		oculiContext.NewFunction(c.Handler.Health.Check, req),
		nil,
		constants.HealthMappers,
	)
	return c.Resource.Responder.NewJSONResponse(ctx, req, result)
}
