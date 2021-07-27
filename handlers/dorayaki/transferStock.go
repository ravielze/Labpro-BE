package dorayaki

import (
	dto "github.com/ravielze/Labpro-BE/model/dto/dorayaki"
	"github.com/ravielze/oculi/request"
)

func (h *handler) TransferStock(req request.Context, item dto.TransferDorayakiRequest) error {
	req.NewTransaction()
	err := h.domain.Dorayaki.TransferStock(req, item)
	if err != nil {
		req.RollbackTransaction()
		return err
	}
	req.CommitTransaction()
	return nil
}
