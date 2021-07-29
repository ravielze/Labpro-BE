package dorayaki

import (
	dto "github.com/ravielze/Labpro-BE/model/dto/dorayaki"
	"github.com/ravielze/oculi/common/functions"
	"github.com/ravielze/oculi/request"
)

func (h *handler) GetAllDorayaki(req request.Context) (dto.DorayakisBaseResponse, error) {
	data := *req.Data()
	page := int(functions.Atoi(data["query.page"], 1))

	dorayakis, pageTotal, err := h.domain.Dorayaki.GetAllDorayaki(req, page)
	if err != nil {
		return dto.DorayakisBaseResponse{}, err
	}
	return dto.NewArrayBaseResponse(dorayakis, pageTotal), nil
}
