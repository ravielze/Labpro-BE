package domain

import (
	dorayaki "github.com/ravielze/Labpro-BE/domain/dorayaki/service"
	shop "github.com/ravielze/Labpro-BE/domain/shop/service"
	"go.uber.org/dig"
)

type (
	Domain struct {
		dig.In

		Dorayaki dorayaki.Service
		Shop     shop.Service
	}
)
