package service

import (
	"app/internal/repository/repos"
)

type Example struct {
	exampleRepo repos.IExampleRepo
}

func NewExampleService(exampleRepo repos.IExampleRepo) *Example {
	return &Example{exampleRepo: exampleRepo}
}

func (s *Example) Create(name string) {
	s.exampleRepo.Create(name)
}
