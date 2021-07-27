package service

import (
	"github.com/ravielze/Labpro-BE/model/dao"
	dto "github.com/ravielze/Labpro-BE/model/dto/shop"
	"github.com/ravielze/oculi/request"
)

func (s *service) Create(req request.Context, item dto.ShopCreateRequest) (dao.Shop, error) {
	return s.repository.Create(req, item.ToDAO())
}
