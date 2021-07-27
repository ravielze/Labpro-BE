package service

import (
	"github.com/ravielze/Labpro-BE/model/dao"
	"github.com/ravielze/oculi/request"
)

func (s *service) GetAllSoldByShop(req request.Context, shopId uint64) ([]dao.Stock, error) {
	return s.stock.GetSoldDorayaki(req, shopId)
}
