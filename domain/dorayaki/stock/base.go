package stock

import "github.com/ravielze/Labpro-BE/resources"

type (
	Repository interface {
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
