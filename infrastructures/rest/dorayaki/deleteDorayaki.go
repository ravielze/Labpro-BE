package dorayaki

import (
	"github.com/labstack/echo/v4"
	"github.com/ravielze/Labpro-BE/constants"
	oculiContext "github.com/ravielze/oculi/context"
	request "github.com/ravielze/oculi/request/echo"
)

func (c *Controller) DeleteDorayaki(ec echo.Context) error {
	ctx := ec.(*oculiContext.Context)
	req := request.New(ctx, c.Resource.Database).Param("dorayakiId")

	result := ctx.Process(
		oculiContext.NewFunction(c.Handler.Dorayaki.DeleteDorayaki, req),
		nil,
		constants.StandardDorayakiMappers,
	)

	return c.Resource.Responder.NewJSONResponse(ctx, req, result)
}
