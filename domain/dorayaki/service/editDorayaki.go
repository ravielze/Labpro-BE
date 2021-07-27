package service

import (
	dto "github.com/ravielze/Labpro-BE/model/dto/dorayaki"
	"github.com/ravielze/oculi/request"
)

func (s *service) EditDorayaki(req request.Context, item dto.DorayakiUpdateRequest) error {
	return s.repository.Update(req, item.ID, item.ToRequestMap())
}
