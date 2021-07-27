package service

import (
	"github.com/ravielze/Labpro-BE/model/dao"
	"github.com/ravielze/oculi/request"
)

func (s *service) UpdateStock(req request.Context, dorayakiId, shopId, stock uint64) (dao.Stock, error) {
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
