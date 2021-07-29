package dorayaki

import (
	"github.com/ravielze/Labpro-BE/constants"
	"github.com/ravielze/oculi/common/functions"
	"github.com/ravielze/oculi/request"
)

func (h *handler) DeleteDorayaki(req request.Context) error {
	data := *req.Data()
	dorayakiId := functions.Atoi(data["parameter.dorayakiId"], 0)
	if dorayakiId == 0 {
		return constants.ErrFailedToParseID
	}
	if _, err := h.domain.Dorayaki.GetDorayaki(req, dorayakiId); err != nil {
		return err
	}
	return h.domain.Dorayaki.DeleteDorayaki(req, dorayakiId)
}
