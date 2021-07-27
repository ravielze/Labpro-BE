package service

import (
	dto "github.com/ravielze/Labpro-BE/model/dto/shop"
	"github.com/ravielze/oculi/request"
)

func (s *service) Update(req request.Context, item dto.ShopUpdateRequest) error {
	return s.repository.Edit(req, item.ID, item.ToRequestMap())
}
