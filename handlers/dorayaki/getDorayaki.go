package dorayaki

import (
	"github.com/ravielze/Labpro-BE/constants"
	dto "github.com/ravielze/Labpro-BE/model/dto/dorayaki"
	"github.com/ravielze/oculi/common/functions"
	"github.com/ravielze/oculi/request"
)

func (h *handler) GetDorayaki(req request.Context) (dto.DorayakiBaseResponse, error) {
	data := *req.Data()
	dorayakiId := functions.Atoi(data["parameter.dorayakiId"], 0)
	if dorayakiId == 0 {
		return dto.DorayakiBaseResponse{}, constants.ErrFailedToParseID
	}
	d, err := h.domain.Dorayaki.GetDorayaki(req, dorayakiId)
	if err != nil {
		return dto.DorayakiBaseResponse{}, err
	}
	return dto.NewBaseResponse(d), nil
}
