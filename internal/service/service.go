package service

import (
	"app/internal/repository/repos"
)

type Deps struct {
	Repos *repos.Repositories
}

type Services struct {
	Example Example
}

func NewServices(deps Deps) *Services {
	exampleService := NewExampleService(deps.Repos.Example)

	return &Services{
		Example: *exampleService,
	}
}
