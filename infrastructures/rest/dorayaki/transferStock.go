package dorayaki

import (
	"github.com/labstack/echo/v4"
	"github.com/ravielze/Labpro-BE/constants"
	dto "github.com/ravielze/Labpro-BE/model/dto/dorayaki"
	oculiContext "github.com/ravielze/oculi/context"
	request "github.com/ravielze/oculi/request/echo"
)

func (c *Controller) TransferStock(ec echo.Context) error {
	ctx := ec.(*oculiContext.Context)
	req := request.New(ctx, c.Resource.Database)

	var item dto.TransferDorayakiRequest
	ctx.BindValidate(&item)

	result := ctx.Process(
		oculiContext.NewFunction(c.Handler.Dorayaki.TransferStock, req, item),
		nil,
		constants.StandardDorayakiMappers,
	)

	return c.Resource.Responder.NewJSONResponse(ctx, req, result)
}
