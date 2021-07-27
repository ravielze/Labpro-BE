package handlers

import (
	"github.com/ravielze/Labpro-BE/handlers/dorayaki"
	"github.com/ravielze/Labpro-BE/handlers/health"
	"github.com/ravielze/Labpro-BE/handlers/shop"
	"github.com/ravielze/oculi/di"
	"go.uber.org/dig"
)

func Register(container *dig.Container) error {
	return di.NewRegistrant(container).
		Provide(health.NewHandler).
		Provide(shop.NewHandler).
		Provide(dorayaki.NewHandler).
		End()
}
