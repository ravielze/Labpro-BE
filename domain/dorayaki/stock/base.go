package stock

import (
	"github.com/ravielze/Labpro-BE/model/dao"
	"github.com/ravielze/Labpro-BE/resources"
	"github.com/ravielze/oculi/request"
)

type (
	Repository interface {
		GetOrCreate(req request.Context, shopId, dorayakiId uint64) (dao.Stock, error)
		Update(req request.Context, shopId, dorayakiId, stock uint64) error
		GetSoldDorayaki(req request.Context, shopId uint64) ([]dao.Stock, error)
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
