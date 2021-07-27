package dorayaki

import (
	"github.com/ravielze/Labpro-BE/constants"
	dto "github.com/ravielze/Labpro-BE/model/dto/dorayaki"
	"github.com/ravielze/oculi/common/functions"
	"github.com/ravielze/oculi/request"
)

func (h *handler) GetStock(req request.Context) (dto.DorayakiResponse, error) {
	data := *req.Data()
	shopId := functions.Atoi(data["parameter.shopId"], 0)
	if shopId == 0 {
		return dto.DorayakiResponse{}, constants.ErrFailedToParseID
	}
	dorayakiId := functions.Atoi(data["parameter.dorayakiId"], 0)
	if dorayakiId == 0 {
		return dto.DorayakiResponse{}, constants.ErrFailedToParseID
	}

	stock, err := h.domain.Dorayaki.GetStock(req, dorayakiId, shopId)
	if err != nil {
		return dto.DorayakiResponse{}, err
	}
	return dto.NewResponse(stock), nil
}
