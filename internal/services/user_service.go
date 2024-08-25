package services

import (
	"context"

	"github.com/ruslanguns/go-simple-api/internal/models"
)

type UserService interface {
	ListUsers(ctx context.Context) ([]models.User, error)
	GetUser(ctx context.Context, userID string) (models.User, error)
	CreateUser(ctx context.Context, user models.User) (models.User, error)
}

type userService struct {
	// Aquí podrías inyectar dependencias como un repositorio
}

func NewUserService() UserService {
	return &userService{}
}

func (s *userService) ListUsers(ctx context.Context) ([]models.User, error) {
	// Implementación simulada
	return []models.User{
		{ID: "1", Name: "John Doe", Email: "john@example.com"},
		{ID: "2", Name: "Jane Doe", Email: "jane@example.com"},
	}, nil
}

func (s *userService) GetUser(ctx context.Context, userID string) (models.User, error) {
	return models.User{ID: userID, Name: "John Doe", Email: "john.doe@foo.com"}, nil
}

func (s *userService) CreateUser(ctx context.Context, user models.User) (models.User, error) {
	// Implementación simulada
	user.ID = "3" // En una implementación real, esto se generaría o se obtendría de la base de datos
	return user, nil
}
