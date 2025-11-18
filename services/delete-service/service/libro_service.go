package service

import (
	"context"
	"delete-service/repository"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type LibroService struct {
	Repo *repository.LibroRepository
}

func (s *LibroService) EliminarLibro(ctx context.Context, id primitive.ObjectID) error {
	return s.Repo.EliminarLibro(ctx, id)
}
