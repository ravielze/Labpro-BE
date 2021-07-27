package dorayaki

import (
	dto "github.com/ravielze/Labpro-BE/model/dto/dorayaki"
	"github.com/ravielze/oculi/request"
)

func (h *handler) TransferStock(req request.Context, item dto.TransferDorayakiRequest) error {
	return h.domain.Dorayaki.TransferStock(req, item)
}
