package service

import (
	"github.com/ravielze/Labpro-BE/domain/dorayaki/repository"
	"github.com/ravielze/Labpro-BE/resources"
)

type (
	Service interface {
	}

	service struct {
		resource   resources.Resource
		repository repository.Repository
	}
)

func New(r resources.Resource, repo repository.Repository) Service {
	return &service{
		resource:   r,
		repository: repo,
	}
}
