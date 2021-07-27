package service

import (
	"github.com/ravielze/oculi/request"
)

func (s *service) DeleteDorayaki(req request.Context, dorayakiId uint64) error {
	return s.repository.Delete(req, dorayakiId)
}
