package health

import (
	"github.com/ravielze/Labpro-BE/constants"
	"github.com/ravielze/oculi/request"
)

func (h *handler) Reset(ctx request.Context) error {
	data := *ctx.Data()
	if data["query.key"] != h.resource.Config.DatabaseResetKey {
		return constants.ErrResetUnauthorized
	}
	h.resource.DBManager.Reset()
	h.resource.DBManager.Install()
	return nil
}
