package shop

import (
	"github.com/ravielze/Labpro-BE/domain/shop/repository"
	"github.com/ravielze/Labpro-BE/domain/shop/service"
	"github.com/ravielze/oculi/di"
	"go.uber.org/dig"
)

func Register(c *dig.Container) error {
	return di.NewRegistrant(c).
		Provide(repository.New).
		Provide(service.New).
		End()
}
