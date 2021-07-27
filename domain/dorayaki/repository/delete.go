package repository

import (
	"github.com/ravielze/Labpro-BE/model/dao"
	"github.com/ravielze/oculi/logs"
	"github.com/ravielze/oculi/request"
)

func (r *repository) Delete(req request.Context, dorayakiId uint64) error {
	item := dao.Dorayaki{ID: dorayakiId}
	if err := req.Transaction().Delete(&item).Error(); err != nil {
		r.resource.Log.StandardError(logs.NewInfo(
			"Dorayaki.Repository.Delete",
			logs.KeyValue("ID", dorayakiId),
			logs.KeyValue("Error", err),
		))
		return err
	}
	return nil
}
