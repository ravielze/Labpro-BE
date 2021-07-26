package rest

import (
	"github.com/ravielze/Labpro-BE/infrastructures/rest/health"
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

		Health health.Controller
	}
)
