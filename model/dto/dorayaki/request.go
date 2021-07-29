package dorayaki

import (
	"github.com/ravielze/Labpro-BE/model/dao"
	"github.com/ravielze/oculi/common/functions/typeutils"
	"github.com/ravielze/oculi/common/model/dto"
)

type (
	DorayakiCreateRequest struct {
		Name             string  `json:"name" binding:"required,min=4,max=256"`
		Description      *string `json:"description" binding:"omitempty,max=1024"`
		TasteName        string  `json:"taste_name" binding:"required,min=4,max=128"`
		TasteDescription *string `json:"taste_description" binding:"omitempty,max=1024"`
	}

	DorayakiUpdateRequest struct {
		ID uint64 `json:"id" binding:"required,min=1"`
		DorayakiCreateRequest
	}

	StockUpdateRequest struct {
		DorayakiID uint64 `json:"id_dorayaki" binding:"required,min=1"`
		ShopID     uint64 `json:"id_shop" binding:"required,min=1"`
		Stock      uint64 `json:"stock" binding:"required,min=0"`
	}

	TransferDorayakiRequest struct {
		DorayakiID uint64 `json:"id_dorayaki" binding:"required,min=1"`
		FromShopID uint64 `json:"id_from_shop" binding:"required,min=1"`
		ToShopID   uint64 `json:"id_to_shop" binding:"required,min=1"`
		Stock      uint64 `json:"amount" binding:"required,min=1"`
	}
)

func (d DorayakiCreateRequest) ToDAO() dao.Dorayaki {
	return dao.Dorayaki{
		Name:             d.Name,
		Description:      typeutils.StringOrNil(d.Description),
		TasteName:        d.TasteName,
		TasteDescription: typeutils.StringOrNil(d.TasteDescription),
	}
}

func (d DorayakiUpdateRequest) ToRequestMap() dto.Map {
	return dto.Map{
		"name":              d.Name,
		"description":       typeutils.StringOrNil(d.Description),
		"taste_name":        d.TasteName,
		"taste_description": typeutils.StringOrNil(d.TasteDescription),
	}
}
