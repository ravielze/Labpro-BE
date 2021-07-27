package service

import (
	"github.com/ravielze/Labpro-BE/model/dao"
	dto "github.com/ravielze/Labpro-BE/model/dto/dorayaki"
	"github.com/ravielze/oculi/request"
)

func (s *service) UpdateStock(req request.Context, item dto.StockUpdateRequest) (dao.Stock, error) {
	shopId := item.ShopID
	dorayakiId := item.DorayakiID
	stock := item.Stock

	// Try to init record if not exists
	if _, err := s.stock.
		GetOrCreate(req, shopId, dorayakiId); err != nil {
		return dao.Stock{}, err
	}

	if err := s.stock.Update(req, shopId, dorayakiId, stock); err != nil {
		return dao.Stock{}, err
	}

	return s.stock.GetOrCreate(req, shopId, dorayakiId)
}
