package service

import (
	"github.com/ravielze/Labpro-BE/model/dao"
	"github.com/ravielze/oculi/request"
)

func (s *service) GetDorayaki(req request.Context, dorayakiId uint64) (dao.Dorayaki, error) {
	return s.repository.Get(req, dorayakiId)
}
