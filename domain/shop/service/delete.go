package service

import (
	"github.com/ravielze/oculi/request"
)

func (s *service) Delete(req request.Context, shopId uint64) error {
	return s.repository.Delete(req, shopId)
}
