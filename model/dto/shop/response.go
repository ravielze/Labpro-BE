package shop

import (
	"github.com/ravielze/Labpro-BE/model/dao"
	"github.com/ravielze/Labpro-BE/model/dto/dorayaki"
	"github.com/ravielze/oculi/common/functions/typeutils"
	"github.com/ravielze/oculi/common/model/dto"
)

type (
	ShopInfoResponse struct {
		ID          uint64 `json:"id"`
		Name        string `json:"name"`
		Description string `json:"description"`

		AddressStreet   string `json:"street"`
		AddressDistrict string `json:"district"`
		AddressCity     string `json:"city"`
		AddressProvince string `json:"province"`
		dto.BaseModel
	}

	ShopResponse struct {
		Shop      ShopInfoResponse           `json:"shop"`
		Dorayakis dorayaki.DorayakisResponse `json:"dorayaki"`
	}

	ShopsResponse []ShopInfoResponse
)

func newInfoResponse(s dao.Shop) ShopInfoResponse {
	return ShopInfoResponse{
		ID:              s.ID,
		Name:            s.Name,
		Description:     typeutils.String(s.Description, ""),
		AddressStreet:   s.AddressStreet,
		AddressDistrict: s.AddressDistrict,
		AddressCity:     s.AddressCity,
		AddressProvince: s.AddressProvince,
		BaseModel:       dto.NewBaseModel(s.BaseModel),
	}
}

func NewArrayResponse(s []dao.Shop) ShopsResponse {
	result := make([]ShopInfoResponse, len(s))
	for i, x := range s {
		result[i] = newInfoResponse(x)
	}
	return ShopsResponse(result)
}

func NewResponse(s dao.Shop, stock []dao.Stock) ShopResponse {
	return ShopResponse{
		Shop:      newInfoResponse(s),
		Dorayakis: dorayaki.NewArrayResponse(stock),
	}
}
