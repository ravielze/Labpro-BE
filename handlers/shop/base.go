package shop

import (
	"github.com/ravielze/Labpro-BE/domain"
	dto "github.com/ravielze/Labpro-BE/model/dto/shop"
	"github.com/ravielze/Labpro-BE/resources"
	"github.com/ravielze/oculi/request"
)

type (
	handler struct {
		domain   domain.Domain
		resource resources.Resource
	}

	Handler interface {
		Create(req request.Context, item dto.ShopCreateRequest) (dto.ShopResponse, error)
		Delete(req request.Context) error
		Update(req request.Context, item dto.ShopUpdateRequest) error
		Get(req request.Context) (dto.ShopResponse, error)
		GetAll(req request.Context) (dto.ShopsResponse, error)
	}
)

func NewHandler(domain domain.Domain, resource resources.Resource) (Handler, error) {
	return &handler{
		domain:   domain,
		resource: resource,
	}, nil
}
