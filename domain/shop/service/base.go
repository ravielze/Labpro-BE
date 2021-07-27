package service

import (
	"github.com/ravielze/Labpro-BE/domain/shop/repository"
	"github.com/ravielze/Labpro-BE/model/dao"
	dto "github.com/ravielze/Labpro-BE/model/dto/shop"
	"github.com/ravielze/Labpro-BE/resources"
	"github.com/ravielze/oculi/request"
)

type (
	Service interface {
		Create(req request.Context, item dto.ShopCreateRequest) (dao.Shop, error)
		Delete(req request.Context, shopId uint64) error
		Update(req request.Context, item dto.ShopUpdateRequest) error
		Get(req request.Context, shopId uint64) (dao.Shop, error)
		GetAll(req request.Context, page int) ([]dao.Shop, int, error)
	}

	service struct {
		resource   resources.Resource
		repository repository.Repository
	}
)

func New(r resources.Resource, repo repository.Repository) Service {
	return &service{
		resource:   r,
		repository: repo,
	}
}
