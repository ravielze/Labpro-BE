package handlers

import (
	"github.com/ravielze/Labpro-BE/handlers/dorayaki"
	"github.com/ravielze/Labpro-BE/handlers/health"
	"github.com/ravielze/Labpro-BE/handlers/shop"
	"go.uber.org/dig"
)

type (
	Handler struct {
		dig.In

		Health   health.Handler
		Shop     shop.Handler
		Dorayaki dorayaki.Handler
	}
)
