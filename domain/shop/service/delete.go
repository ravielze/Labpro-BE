package service

import (
	"github.com/ravielze/oculi/request"
)

func (s *service) Delete(req request.Context, shopId uint64) error {
	_, err := s.repository.Get(req, shopId)
	if err != nil {
		return err
	}
	return s.repository.Delete(req, shopId)
}
