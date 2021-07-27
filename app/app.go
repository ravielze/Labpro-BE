package app

import (
	"github.com/ravielze/Labpro-BE/di"
	"github.com/ravielze/Labpro-BE/infrastructures"
	"github.com/ravielze/Labpro-BE/resources"
	"github.com/ravielze/oculi/app"

	webserver "github.com/ravielze/oculi/server/echo"
	"go.uber.org/dig"
)

func Run() {
	invoker := func(container *dig.Container) error {
		return container.Invoke(
			func(i infrastructures.Component, r resources.Resource) error {
				s := webserver.New(i, r)
				if r.Config.IsDevelopment() {
					s.DevelopmentMode()
				}
				return s.Run()
			},
		)
	}

	if err := app.Run(di.Container, invoker); err != nil {
		panic(err)
	}
}
