package service

import (
	"github.com/ravielze/Labpro-BE/domain/dorayaki/repository"
	"github.com/ravielze/Labpro-BE/domain/dorayaki/stock"
	"github.com/ravielze/Labpro-BE/model/dao"
	dto "github.com/ravielze/Labpro-BE/model/dto/dorayaki"
	"github.com/ravielze/Labpro-BE/resources"
	"github.com/ravielze/oculi/request"
)

type (
	Service interface {
		CreateDorayaki(req request.Context, item dto.DorayakiCreateRequest) (dao.Dorayaki, error)
		EditDorayaki(req request.Context, item dto.DorayakiUpdateRequest) error
		DeleteDorayaki(req request.Context, dorayakiId uint64) error
		GetDorayaki(req request.Context, dorayakiId uint64) (dao.Dorayaki, error)

		GetStock(req request.Context, dorayakiId, shopId uint64) (dao.Stock, error)
		UpdateStock(req request.Context, item dto.StockUpdateRequest) (dao.Stock, error)
		TransferStock(req request.Context, item dto.TransferDorayakiRequest) error
		GetAllSoldByShop(req request.Context, shopId uint64) ([]dao.Stock, error)
	}

	service struct {
		resource   resources.Resource
		repository repository.Repository
		stock      stock.Repository
	}
)

func New(r resources.Resource, repo repository.Repository, st stock.Repository) Service {
	return &service{
		resource:   r,
		repository: repo,
		stock:      st,
	}
}
