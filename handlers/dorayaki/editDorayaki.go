package dorayaki

import (
	dto "github.com/ravielze/Labpro-BE/model/dto/dorayaki"
	"github.com/ravielze/oculi/request"
)

func (h *handler) EditDorayaki(req request.Context, item dto.DorayakiUpdateRequest) error {
	if _, err := h.domain.Dorayaki.GetDorayaki(req, item.ID); err != nil {
		return err
	}
	return h.domain.Dorayaki.EditDorayaki(req, item)
}
