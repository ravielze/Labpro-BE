package rest

import (
	"github.com/ravielze/Labpro-BE/infrastructures/rest/dorayaki"
	"github.com/ravielze/Labpro-BE/infrastructures/rest/health"
	"github.com/ravielze/Labpro-BE/infrastructures/rest/shop"
	"github.com/ravielze/Labpro-BE/resources"
	"go.uber.org/dig"
)

type (
	Rest struct {
		dig.In

		Controller Controller
		Resource   resources.Resource
	}

	Controller struct {
		dig.In

		Health   health.Controller
		Shop     shop.Controller
		Dorayaki dorayaki.Controller
	}
)
