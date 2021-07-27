package repository

import (
	"github.com/ravielze/Labpro-BE/model/dao"
	"github.com/ravielze/oculi/logs"
	"github.com/ravielze/oculi/request"
)

func (r *repository) Get(req request.Context, dorayakiId uint64) (dao.Dorayaki, error) {
	item := dao.Dorayaki{ID: dorayakiId}
	if err := req.Transaction().Take(&item).Error(); err != nil {
		r.resource.Log.StandardError(logs.NewInfo(
			"Dorayaki.Repository.Get",
			logs.KeyValue("ID", dorayakiId),
			logs.KeyValue("Error", err),
		))
		return dao.Dorayaki{}, err
	}
	return item, nil
}
