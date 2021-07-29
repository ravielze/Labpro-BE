package stock

import (
	"github.com/ravielze/Labpro-BE/model/dao"
	"github.com/ravielze/oculi/logs"
	"github.com/ravielze/oculi/request"
)

func (r *repository) GetOrCreate(req request.Context, shopId, dorayakiId uint64) (dao.Stock, error) {
	var result dao.Stock
	if err := req.
		Transaction().
		Preload("Dorayaki").
		Where("shop_id = ?", shopId).
		Where("dorayaki_id = ?", dorayakiId).
		FirstOrCreate(&result,
			dao.Stock{
				ShopID:     shopId,
				DorayakiID: dorayakiId,
				Stock:      0,
			}).
		Error(); err != nil {
		r.resource.Log.StandardError(logs.NewInfo(
			"Dorayaki.Stock.GetOrCreate",
			logs.KeyValue("ShopID", shopId),
			logs.KeyValue("DorayakiID", dorayakiId),
			logs.KeyValue("Error", err),
		))
		return dao.Stock{}, err
	}

	if err := req.Transaction().Preload("Dorayaki").
		Where("shop_id = ?", shopId).
		Where("dorayaki_id = ?", dorayakiId).
		Take(&result).
		Error(); err != nil {
		r.resource.Log.StandardError(logs.NewInfo(
			"Dorayaki.Stock.GetOrCreate",
			logs.KeyValue("ShopID", shopId),
			logs.KeyValue("DorayakiID", dorayakiId),
			logs.KeyValue("Error", err),
		))
		return dao.Stock{}, err
	}
	return result, nil
}
