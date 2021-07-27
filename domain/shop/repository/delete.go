package repository

import (
	"github.com/ravielze/Labpro-BE/model/dao"
	"github.com/ravielze/oculi/logs"
	"github.com/ravielze/oculi/request"
)

func (r *repository) Delete(req request.Context, shopId uint64) error {
	item := dao.Shop{ID: shopId}
	if err := req.Transaction().
		Delete(&item).Error(); err != nil {
		if err := req.Transaction().Create(&item).Error(); err != nil {
			r.resource.Log.StandardError(logs.NewInfo(
				"Shop.Repository.Delete",
				logs.KeyValue("ID", shopId),
				logs.KeyValue("Error", err),
			))
			return err
		}
	}
	return nil
}
