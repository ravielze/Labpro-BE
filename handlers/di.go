package handlers

import (
	"github.com/ravielze/Labpro-BE/handlers/health"
	"github.com/ravielze/oculi/di"
	"go.uber.org/dig"
)

func Register(container *dig.Container) error {
	return di.NewRegistrant(container).
		Provide(health.NewHandler).
		End()
}
