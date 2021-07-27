package dorayaki

import (
	dto "github.com/ravielze/Labpro-BE/model/dto/dorayaki"
	"github.com/ravielze/oculi/request"
)

func (h *handler) UpdateStock(req request.Context, item dto.StockUpdateRequest) (dto.DorayakiResponse, error) {
	stock, err := h.domain.Dorayaki.UpdateStock(req, item)
	if err != nil {
		return dto.DorayakiResponse{}, err
	}
	return dto.NewResponse(stock), nil
}
