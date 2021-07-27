package service

import (
	"github.com/ravielze/Labpro-BE/model/dao"
	"github.com/ravielze/oculi/request"
)

func (s *service) GetAll(req request.Context, page int) ([]dao.Shop, int, error) {
	return s.repository.GetAll(req, page)
}
