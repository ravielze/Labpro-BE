package dorayaki

import (
	"github.com/labstack/echo/v4"
	oculiContext "github.com/ravielze/oculi/context"
	request "github.com/ravielze/oculi/request/echo"
)

func (c *Controller) GetAllDorayaki(ec echo.Context) error {
	ctx := ec.(*oculiContext.Context)
	req := request.New(ctx, c.Resource.Database).Query("page", "")

	result := ctx.Process(
		oculiContext.NewFunction(c.Handler.Dorayaki.GetAllDorayaki, req),
		nil,
		nil,
	)

	return c.Resource.Responder.NewJSONResponse(ctx, req, result)
}
