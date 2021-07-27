package repository

import (
	"github.com/ravielze/Labpro-BE/model/dao"
	"github.com/ravielze/Labpro-BE/resources"
	"github.com/ravielze/oculi/common/model/dto"
	"github.com/ravielze/oculi/request"
)

type (
	Repository interface {
		Get(req request.Context, dorayakiId uint64) (dao.Dorayaki, error)
		Create(req request.Context, item dao.Dorayaki) (dao.Dorayaki, error)
		Update(req request.Context, dorayakiId uint64, updateMap dto.Map) error
		Delete(req request.Context, dorayakiId uint64) error
	}

	repository struct {
		resource resources.Resource
	}
)

func New(r resources.Resource) Repository {
	return &repository{
		resource: r,
	}
}
