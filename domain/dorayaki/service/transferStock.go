package service

import (
	"github.com/ravielze/Labpro-BE/constants"
	dto "github.com/ravielze/Labpro-BE/model/dto/dorayaki"
	"github.com/ravielze/oculi/request"
)

func (s *service) TransferStock(req request.Context, item dto.TransferDorayakiRequest) error {
	fromShopId := item.FromShopID
	dorayakiId := item.DorayakiID
	toShopId := item.ToShopID
	stock := item.Stock

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
