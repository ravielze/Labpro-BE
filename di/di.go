package di

import (
	"sync"

	"github.com/ravielze/Labpro-BE/config"
	"github.com/ravielze/Labpro-BE/domain"
	"github.com/ravielze/Labpro-BE/handlers"
	"github.com/ravielze/Labpro-BE/resources"
	"github.com/ravielze/oculi/di"
	"go.uber.org/dig"
)

var (
	container *dig.Container
	once      sync.Once
)

func Container() (*dig.Container, error) {
	items := []di.Registerable{
		config.Register,
		resources.Register,
		domain.Register,
		handlers.Register,
	}
	return di.Container(items)(&once, container)
}
