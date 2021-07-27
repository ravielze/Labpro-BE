package dorayaki

import (
	dto "github.com/ravielze/Labpro-BE/model/dto/dorayaki"
	"github.com/ravielze/oculi/request"
)

func (h *handler) CreateDorayaki(req request.Context, item dto.DorayakiCreateRequest) (dto.DorayakiBaseResponse, error) {
	d, err := h.domain.Dorayaki.CreateDorayaki(req, item)
	if err != nil {
		return dto.DorayakiBaseResponse{}, err
	}
	return dto.NewBaseResponse(d), nil
}
