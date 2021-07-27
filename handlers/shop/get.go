package shop

import (
	"github.com/ravielze/Labpro-BE/constants"
	dto "github.com/ravielze/Labpro-BE/model/dto/shop"
	"github.com/ravielze/oculi/common/functions"
	"github.com/ravielze/oculi/request"
)

func (h *handler) Get(req request.Context) (dto.ShopResponse, error) {
	data := *req.Data()
	shopId := functions.Atoi(data["parameter.id"], 0)
	if shopId == 0 {
		return dto.ShopResponse{}, constants.ErrFailedToParseID
	}
	s, err := h.domain.Shop.Get(req, shopId)
	if err != nil {
		return dto.ShopResponse{}, err
	}

	stocks, err := h.domain.Dorayaki.GetAllSoldByShop(req, shopId)
	if err != nil {
		return dto.ShopResponse{}, err
	}

	return dto.NewResponse(s, stocks), nil
}
