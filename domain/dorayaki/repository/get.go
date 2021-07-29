package repository

import (
	"errors"

	"github.com/ravielze/Labpro-BE/constants"
	"github.com/ravielze/Labpro-BE/model/dao"
	consts "github.com/ravielze/oculi/constant/errors"
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
		if errors.Is(err, consts.ErrRecordNotFound) {
			return dao.Dorayaki{}, constants.ErrDorayakiNotFound
		}
		return dao.Dorayaki{}, err
	}
	return item, nil
}
