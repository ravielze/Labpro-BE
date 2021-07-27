package service

import (
	"github.com/ravielze/Labpro-BE/model/dao"
	"github.com/ravielze/oculi/request"
)

func (s *service) Get(req request.Context, shopId uint64) (dao.Shop, error) {
	return s.repository.Get(req, shopId)
}
