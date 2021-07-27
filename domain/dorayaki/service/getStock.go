package service

import (
	"github.com/ravielze/Labpro-BE/model/dao"
	"github.com/ravielze/oculi/request"
)

func (s *service) GetStock(req request.Context, dorayakiId, shopId uint64) (dao.Stock, error) {
	return s.stock.GetOrCreate(req, shopId, dorayakiId)
}
