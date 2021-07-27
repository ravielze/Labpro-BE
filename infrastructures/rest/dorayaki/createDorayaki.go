package dorayaki

import (
	"github.com/labstack/echo/v4"
	dto "github.com/ravielze/Labpro-BE/model/dto/dorayaki"
	oculiContext "github.com/ravielze/oculi/context"
	request "github.com/ravielze/oculi/request/echo"
)

func (c *Controller) CreateDorayaki(ec echo.Context) error {
	ctx := ec.(*oculiContext.Context)
	req := request.New(ctx, c.Resource.Database)

	var item dto.DorayakiCreateRequest
	ctx.BindValidate(&item)

	result := ctx.Process(
		oculiContext.NewFunction(c.Handler.Dorayaki.CreateDorayaki, req, item),
		nil,
		nil,
	)

	return c.Resource.Responder.NewJSONResponse(ctx, req, result)
}
