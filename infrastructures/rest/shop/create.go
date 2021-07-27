package shop

import (
	"github.com/labstack/echo/v4"
	dto "github.com/ravielze/Labpro-BE/model/dto/shop"
	oculiContext "github.com/ravielze/oculi/context"
	request "github.com/ravielze/oculi/request/echo"
)

func (c *Controller) Create(ec echo.Context) error {
	ctx := ec.(*oculiContext.Context)
	req := request.New(ctx, c.Resource.Database)

	var item dto.ShopCreateRequest
	ctx.BindValidate(&item)

	result := ctx.Process(
		oculiContext.NewFunction(c.Handler.Shop.Create, req, item),
		nil,
		nil,
	)

	return c.Resource.Responder.NewJSONResponse(ctx, req, result)
}
