package domain

import (
	"github.com/ravielze/Labpro-BE/domain/dorayaki"
	"github.com/ravielze/Labpro-BE/domain/shop"
	"github.com/ravielze/oculi/di"
	"go.uber.org/dig"
)

func Register(c *dig.Container) error {
	return di.NewRegistrant(c).
		Register(dorayaki.Register).
		Register(shop.Register).
		End()
}
