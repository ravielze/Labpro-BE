package dorayaki

import (
	dto "github.com/ravielze/Labpro-BE/model/dto/dorayaki"
	"github.com/ravielze/oculi/request"
)

func (h *handler) UpdateStock(req request.Context, item dto.StockUpdateRequest) (dto.DorayakiResponse, error) {
	req.NewTransaction()
	stock, err := h.domain.Dorayaki.UpdateStock(req, item)
	if err != nil {
		req.RollbackTransaction()
		return dto.DorayakiResponse{}, err
	}
	req.CommitTransaction()
	return dto.NewResponse(stock), nil
}
