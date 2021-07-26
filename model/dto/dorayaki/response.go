package dorayaki

import (
	"github.com/ravielze/Labpro-BE/model/dao"
	"github.com/ravielze/oculi/common/functions/typeutils"
	"github.com/ravielze/oculi/common/model/dto"
)

type (
	DorayakiResponse struct {
		ID               uint64 `json:"id"`
		Name             string `json:"name"`
		Description      string `json:"description"`
		TasteName        string `json:"taste_name"`
		TasteDescription string `json:"taste_description"`
		Stock            uint64 `json:"stock"`
		dto.BaseModel
	}

	DorayakisResponse []DorayakiResponse
)

func NewResponse(d dao.Dorayaki, stock dao.Stock) DorayakiResponse {
	return DorayakiResponse{
		ID:               d.ID,
		Name:             d.Name,
		Description:      typeutils.String(d.Description, ""),
		TasteName:        d.TasteName,
		TasteDescription: typeutils.String(d.TasteDescription, ""),
		BaseModel:        dto.NewBaseModel(d.BaseModel),
		Stock:            stock.Stock,
	}
}

func NewArrayResponse(d []dao.Dorayaki, stock []dao.Stock) DorayakisResponse {
	stockSize := len(stock)
	dorayakiIDToStockMap := make(map[uint64]dao.Stock, stockSize)
	for _, x := range stock {
		dorayakiIDToStockMap[x.DorayakiID] = x
	}

	result := make([]DorayakiResponse, len(d))
	for i, x := range d {
		result[i] = NewResponse(x, dorayakiIDToStockMap[x.ID])
	}

	return DorayakisResponse(result)
}
