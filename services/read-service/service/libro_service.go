package service

import (
	"context"
	"read-service/model"
	"read-service/repository"
)

type LibroService struct {
	Repo repository.LibroRepositorio
}

func NewLibroService(repo repository.LibroRepositorio) *LibroService {
	return &LibroService{Repo: repo}
}

func (s *LibroService) ObtenerLibros(ctx context.Context) ([]model.Libro, error) {
	return s.Repo.ObtenerTodos(ctx)
}
