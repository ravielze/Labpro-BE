package shop

import (
	"github.com/ravielze/Labpro-BE/model/dao"
	dto "github.com/ravielze/Labpro-BE/model/dto/shop"
	"github.com/ravielze/oculi/request"
)

func (h *handler) Create(req request.Context, item dto.ShopCreateRequest) (dto.ShopResponse, error) {
	s, err := h.domain.Shop.Create(req, item)
	if err != nil {
		return dto.ShopResponse{}, err
	}
	return dto.NewResponse(s, []dao.Stock{}), nil
}
