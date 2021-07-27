package service

import (
	"github.com/ravielze/Labpro-BE/model/dao"
	dto "github.com/ravielze/Labpro-BE/model/dto/dorayaki"
	"github.com/ravielze/oculi/request"
)

func (s *service) CreateDorayaki(req request.Context, item dto.DorayakiCreateRequest) (dao.Dorayaki, error) {
	return s.repository.Create(req, item.ToDAO())
}
