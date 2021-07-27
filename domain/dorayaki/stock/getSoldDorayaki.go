package stock

import (
	"github.com/ravielze/Labpro-BE/model/dao"
	"github.com/ravielze/oculi/logs"
	"github.com/ravielze/oculi/request"
)

func (r *repository) GetSoldDorayaki(req request.Context, shopId uint64) ([]dao.Stock, error) {
	var result []dao.Stock
	if err := req.Transaction().
		Preload("Dorayaki").
		Where("shop_id = ?", shopId).
		Where("stock > ?", 0).
		Find(&result).
		Error(); err != nil {
		r.resource.Log.StandardError(logs.NewInfo(
			"Dorayaki.Stock.GetSoldDorayaki",
			logs.KeyValue("ShopID", shopId),
			logs.KeyValue("Error", err),
		))
		return nil, err
	}
	return result, nil
}
