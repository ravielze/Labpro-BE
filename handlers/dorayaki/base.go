package dorayaki

import (
	"github.com/ravielze/Labpro-BE/domain"
	dto "github.com/ravielze/Labpro-BE/model/dto/dorayaki"
	"github.com/ravielze/Labpro-BE/resources"
	"github.com/ravielze/oculi/request"
)

type (
	handler struct {
		domain   domain.Domain
		resource resources.Resource
	}

	Handler interface {
		CreateDorayaki(req request.Context, item dto.DorayakiCreateRequest) (dto.DorayakiBaseResponse, error)
		EditDorayaki(req request.Context, item dto.DorayakiUpdateRequest) error
		DeleteDorayaki(req request.Context) error
		GetDorayaki(req request.Context) (dto.DorayakiBaseResponse, error)

		GetStock(req request.Context) (dto.DorayakiResponse, error)
		UpdateStock(req request.Context, item dto.StockUpdateRequest) (dto.DorayakiResponse, error)
		TransferStock(req request.Context, item dto.TransferDorayakiRequest) error
		GetAllSoldByShop(req request.Context) (dto.DorayakisResponse, error)
	}
)

func NewHandler(domain domain.Domain, resource resources.Resource) (Handler, error) {
	return &handler{
		domain:   domain,
		resource: resource,
	}, nil
}
