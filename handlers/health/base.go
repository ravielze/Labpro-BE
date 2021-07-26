package health

import (
	"github.com/ravielze/Labpro-BE/resources"
	"github.com/ravielze/oculi/common/model/dto/health"
	"github.com/ravielze/oculi/request"
)

type (
	Handler interface {
		Check(ctx request.Context) health.CheckResponseDTO
		Reset(ctx request.Context) error
	}

	handler struct {
		resource resources.Resource
	}
)

func NewHandler(resource resources.Resource) (Handler, error) {
	return &handler{
		resource: resource,
	}, nil
}
