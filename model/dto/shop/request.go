package shop

import (
	"github.com/ravielze/Labpro-BE/model/dao"
	"github.com/ravielze/oculi/common/functions/typeutils"
	"github.com/ravielze/oculi/common/model/dto"
)

type (
	ShopCreateRequest struct {
		Name        string  `json:"name" binding:"required,max=256"`
		Description *string `json:"description" binding:"max=1024"`
		Street      string  `json:"street" binding:"required,min=4,max=256"`
		District    string  `json:"district" binding:"required,min=4,max=128"`
		City        string  `json:"city" binding:"required,min=4,max=128"`
		Province    string  `json:"province" binding:"required,min=4,max=128"`
	}

	ShopUpdateRequest struct {
		ID uint64 `json:"id" binding:"required,min=1"`
		ShopCreateRequest
	}
)

func (s ShopCreateRequest) ToDAO() dao.Shop {
	return dao.Shop{
		Name:            s.Name,
		Description:     typeutils.StringOrNil(s.Description),
		AddressStreet:   s.Street,
		AddressDistrict: s.District,
		AddressCity:     s.City,
		AddressProvince: s.Province,
	}
}

func (s ShopUpdateRequest) ToRequestMap() dto.Map {
	return dto.Map{
		"name":             s.Name,
		"description":      typeutils.StringOrNil(s.Description),
		"address_street":   s.Street,
		"address_district": s.District,
		"address_city":     s.City,
		"address_province": s.Province,
	}
}
