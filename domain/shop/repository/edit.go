package repository

import (
	"github.com/ravielze/Labpro-BE/model/dao"
	"github.com/ravielze/oculi/common/model/dto"
	"github.com/ravielze/oculi/logs"
	"github.com/ravielze/oculi/request"
)

func (r *repository) Edit(req request.Context, shopId uint64, item dto.Map) error {
	if err := req.Transaction().
		Model(dao.Shop{}).
		Where("id = ?", shopId).
		Updates(item.ToMap()).
		Error(); err != nil {
		r.resource.Log.StandardError(logs.NewInfo(
			"Shop.Repository.Edit",
			logs.KeyValue("ID", shopId),
			logs.KeyValue("RequestMap", item.ToMap()),
			logs.KeyValue("Error", err),
		))
		return err
	}
	return nil
}
