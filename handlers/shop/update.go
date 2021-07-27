package shop

import (
	dto "github.com/ravielze/Labpro-BE/model/dto/shop"
	"github.com/ravielze/oculi/request"
)

func (h *handler) Update(req request.Context, item dto.ShopUpdateRequest) error {
	return h.domain.Shop.Update(req, item)
}
