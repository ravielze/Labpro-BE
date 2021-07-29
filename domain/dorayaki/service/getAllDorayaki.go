package service

import (
	"github.com/ravielze/Labpro-BE/model/dao"
	"github.com/ravielze/oculi/request"
)

func (s *service) GetAllDorayaki(req request.Context, page int) ([]dao.Dorayaki, int, error) {
	return s.repository.GetAll(req, page)
}
