package shop

import (
	"github.com/ravielze/Labpro-BE/constants"
	"github.com/ravielze/oculi/common/functions"
	"github.com/ravielze/oculi/request"
)

func (h *handler) Delete(req request.Context) error {
	data := *req.Data()
	shopId := functions.Atoi(data["parameter.id"], 0)
	if shopId == 0 {
		return constants.ErrFailedToParseID
	}
	return h.domain.Shop.Delete(req, shopId)
}
