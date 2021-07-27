package repository

import (
	"github.com/ravielze/Labpro-BE/model/dao"
	"github.com/ravielze/Labpro-BE/resources"
	"github.com/ravielze/oculi/common/model/dto"
	"github.com/ravielze/oculi/request"
)

type (
	Repository interface {
		Get(req request.Context, shopId uint64) (dao.Shop, error)
		GetAll(req request.Context, page int) ([]dao.Shop, int, error)
		Create(req request.Context, item dao.Shop) (dao.Shop, error)
		Edit(req request.Context, shopId uint64, item dto.Map) error
		Delete(req request.Context, shopId uint64) error
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
