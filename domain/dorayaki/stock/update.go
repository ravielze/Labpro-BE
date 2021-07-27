package stock

import (
	"github.com/ravielze/Labpro-BE/model/dao"
	"github.com/ravielze/oculi/logs"
	"github.com/ravielze/oculi/request"
)

func (r *repository) Update(req request.Context, shopId, dorayakiId, stock uint64) error {
	if err := req.Transaction().
		Model(dao.Stock{}).
		Where("shop_id = ?", shopId).
		Where("dorayaki_id = ?", dorayakiId).
		Update("stock", stock).Error(); err != nil {
		r.resource.Log.StandardError(logs.NewInfo(
			"Dorayaki.Stock.Update",
			logs.KeyValue("ShopID", shopId),
			logs.KeyValue("DorayakiID", dorayakiId),
			logs.KeyValue("Stock", stock),
			logs.KeyValue("Error", err),
		))
		return err
	}
	return nil
}
