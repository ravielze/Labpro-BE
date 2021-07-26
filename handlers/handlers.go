package handlers

import (
	"github.com/ravielze/Labpro-BE/handlers/health"
	"go.uber.org/dig"
)

type (
	Handler struct {
		dig.In

		Health health.Handler
	}
)
