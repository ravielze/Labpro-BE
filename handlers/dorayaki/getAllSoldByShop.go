package dorayaki

import (
	"github.com/ravielze/Labpro-BE/constants"
	dto "github.com/ravielze/Labpro-BE/model/dto/dorayaki"
	"github.com/ravielze/oculi/common/functions"
	"github.com/ravielze/oculi/request"
)

func (h *handler) GetAllSoldByShop(req request.Context) (dto.DorayakisResponse, error) {
	data := *req.Data()
	shopId := functions.Atoi(data["parameter.shopId"], 0)
	if shopId == 0 {
		return dto.DorayakisResponse{}, constants.ErrFailedToParseID
	}

	if _, err := h.domain.Shop.Get(req, shopId); err != nil {
		return dto.DorayakisResponse{}, err
	}

	stocks, err := h.domain.Dorayaki.GetAllSoldByShop(req, shopId)
	if err != nil {
		return dto.DorayakisResponse{}, err
	}
	return dto.NewArrayResponse(stocks), nil
}
