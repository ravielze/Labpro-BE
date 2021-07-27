package dorayaki

import (
	"github.com/ravielze/Labpro-BE/model/dao"
	"github.com/ravielze/oculi/common/functions/typeutils"
	"github.com/ravielze/oculi/common/model/dto"
)

type (
	DorayakiBaseResponse struct {
		ID               uint64 `json:"id"`
		Name             string `json:"name"`
		Description      string `json:"description"`
		TasteName        string `json:"taste_name"`
		TasteDescription string `json:"taste_description"`
		dto.BaseModel
	}
	DorayakiResponse struct {
		DorayakiBaseResponse
		Stock uint64 `json:"stock"`
	}

	DorayakisBaseResponse []DorayakiBaseResponse
	DorayakisResponse     []DorayakiResponse
)

func NewBaseResponse(d dao.Dorayaki) DorayakiBaseResponse {
	return DorayakiBaseResponse{
		ID:               d.ID,
		Name:             d.Name,
		Description:      typeutils.String(d.Description, ""),
		TasteName:        d.TasteName,
		TasteDescription: typeutils.String(d.TasteDescription, ""),
		BaseModel:        dto.NewBaseModel(d.BaseModel),
	}
}

func NewResponse(s dao.Stock) DorayakiResponse {
	return DorayakiResponse{
		Stock:                s.Stock,
		DorayakiBaseResponse: NewBaseResponse(s.Dorayaki),
	}
}

func NewArrayResponse(stock []dao.Stock) DorayakisResponse {
	result := make([]DorayakiResponse, len(stock))
	for i, x := range stock {
		result[i] = NewResponse(x)
	}

	return DorayakisResponse(result)
}

func NewArrayBaseResponse(d []dao.Dorayaki) DorayakisBaseResponse {
	result := make([]DorayakiBaseResponse, len(d))
	for i, x := range d {
		result[i] = NewBaseResponse(x)
	}

	return DorayakisBaseResponse(result)
}
