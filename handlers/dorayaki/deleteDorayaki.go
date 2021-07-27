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
	return h.domain.Dorayaki.DeleteDorayaki(req, dorayakiId)
}
