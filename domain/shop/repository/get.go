package repository

import (
	"github.com/ravielze/Labpro-BE/model/dao"
	"github.com/ravielze/oculi/logs"
	"github.com/ravielze/oculi/request"
)

func (r *repository) Get(req request.Context, shopId uint64) (dao.Shop, error) {
	item := dao.Shop{
		ID: shopId,
	}
	if err := req.Transaction().Take(&item).Error(); err != nil {
		r.resource.Log.StandardError(logs.NewInfo(
			"Shop.Repository.Get",
			logs.KeyValue("ID", shopId),
			logs.KeyValue("Error", err),
		))
		return dao.Shop{}, err
	}
	return item, nil
}
