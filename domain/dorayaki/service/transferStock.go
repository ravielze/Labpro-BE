package service

import (
	"github.com/ravielze/Labpro-BE/constants"
	"github.com/ravielze/oculi/request"
)

func (s *service) TransferStock(req request.Context, dorayakiId, fromShopId, toShopId, stock uint64) error {
	fromStock, err := s.stock.GetOrCreate(req, fromShopId, dorayakiId)
	if err != nil {
		return err
	}

	toStock, err := s.stock.GetOrCreate(req, toShopId, dorayakiId)
	if err != nil {
		return err
	}

	if fromStock.Stock < stock {
		return constants.ErrStockNotEnough
	}

	if err := s.stock.
		Update(req, toShopId,
			dorayakiId, toStock.Stock+stock); err != nil {
		return err
	}
	return nil
}
