package repository

import (
	"github.com/ravielze/Labpro-BE/model/dao"
	"github.com/ravielze/oculi/common/model/dto"
	"github.com/ravielze/oculi/logs"
	"github.com/ravielze/oculi/request"
)

func (r *repository) Update(req request.Context, dorayakiId uint64, updateMap dto.Map) error {
	if err := req.Transaction().Model(dao.Dorayaki{}).
		Where("id = ?", dorayakiId).
		Updates(updateMap.ToMap()).Error(); err != nil {
		r.resource.Log.StandardError(logs.NewInfo(
			"Todo.Repository.Update",
			logs.KeyValue("ID", dorayakiId),
			logs.KeyValue("RequestMap", updateMap.ToMap()),
			logs.KeyValue("Error", err),
		))
		return err
	}
	return nil
}
