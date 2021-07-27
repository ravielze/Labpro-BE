package shop

import (
	dto "github.com/ravielze/Labpro-BE/model/dto/shop"
	"github.com/ravielze/oculi/common/functions"
	"github.com/ravielze/oculi/request"
)

func (h *handler) GetAll(req request.Context) (dto.ShopsResponse, error) {
	data := *req.Data()
	page := int(functions.Atoi(data["query.page"], 1))

	shops, pageTotal, err := h.domain.Shop.GetAll(req, page)
	if err != nil {
		return dto.ShopsResponse{}, err
	}
	return dto.NewArrayResponse(shops, pageTotal), nil
}
