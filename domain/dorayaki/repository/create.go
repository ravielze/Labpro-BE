package repository

import (
	"github.com/ravielze/Labpro-BE/model/dao"
	"github.com/ravielze/oculi/logs"
	"github.com/ravielze/oculi/request"
)

func (r *repository) Create(req request.Context, item dao.Dorayaki) (dao.Dorayaki, error) {
	if err := req.Transaction().Create(&item).Error(); err != nil {
		r.resource.Log.StandardError(logs.NewInfo(
			"Dorayaki.Repository.Create",
			logs.KeyValue("Item", item),
			logs.KeyValue("Error", err),
		))
		return dao.Dorayaki{}, err
	}
	return item, nil
}
