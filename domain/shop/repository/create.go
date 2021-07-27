package repository

import (
	"github.com/ravielze/Labpro-BE/model/dao"
	"github.com/ravielze/oculi/logs"
	"github.com/ravielze/oculi/request"
)

func (r *repository) Create(req request.Context, item dao.Shop) (dao.Shop, error) {
	if err := req.Transaction().Create(&item).Error(); err != nil {
		r.resource.Log.StandardError(logs.NewInfo(
			"Shop.Repository.Create",
			logs.KeyValue("Item", item),
			logs.KeyValue("Error", err),
		))
		return dao.Shop{}, err
	}
	return item, nil
}
