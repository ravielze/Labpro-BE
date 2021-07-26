package dorayaki

import (
	"github.com/ravielze/Labpro-BE/model/dao"
	"github.com/ravielze/oculi/common/functions/typeutils"
	"github.com/ravielze/oculi/common/model/dto"
)

type (
	DorayakiCreateRequest struct {
		Name             string  `json:"name" binding:"required,min=4,max=256"`
		Description      *string `json:"description" binding:"max=1024"`
		TasteName        string  `json:"taste_name" binding:"required,min=4,max=128"`
		TasteDescription *string `json:"taste_description" binding:"max=1024"`
	}

	DorayakiUpdateRequest struct {
		ID uint64 `json:"id" binding:"required,min=1"`
		DorayakiCreateRequest
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
