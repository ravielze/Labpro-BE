package dorayaki

import (
	dto "github.com/ravielze/Labpro-BE/model/dto/dorayaki"
	"github.com/ravielze/oculi/request"
)

func (h *handler) TransferStock(req request.Context, item dto.TransferDorayakiRequest) error {
	if _, err := h.domain.Dorayaki.GetDorayaki(req, item.DorayakiID); err != nil {
		return err
	}

	if _, err := h.domain.Shop.Get(req, item.FromShopID); err != nil {
		return err
	}

	if _, err := h.domain.Shop.Get(req, item.ToShopID); err != nil {
		return err
	}

	req.NewTransaction()
	err := h.domain.Dorayaki.TransferStock(req, item)
	if err != nil {
		req.RollbackTransaction()
		return err
	}
	req.CommitTransaction()
	return nil
}
