package repository

import (
	"errors"

	"github.com/ravielze/Labpro-BE/constants"
	"github.com/ravielze/Labpro-BE/model/dao"
	consts "github.com/ravielze/oculi/constant/errors"
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
		if errors.Is(err, consts.ErrRecordNotFound) {
			return dao.Shop{}, constants.ErrShopNotFound
		}
		return dao.Shop{}, err
	}
	return item, nil
}
