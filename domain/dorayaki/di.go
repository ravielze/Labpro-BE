package dorayaki

import (
	"github.com/ravielze/Labpro-BE/domain/dorayaki/repository"
	"github.com/ravielze/Labpro-BE/domain/dorayaki/service"
	"github.com/ravielze/Labpro-BE/domain/dorayaki/stock"
	"github.com/ravielze/oculi/di"
	"go.uber.org/dig"
)

func Register(c *dig.Container) error {
	return di.NewRegistrant(c).
		Provide(repository.New).
		Provide(service.New).
		Provide(stock.New).
		End()
}
