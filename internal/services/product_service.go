package services

import (
	"context"

	"github.com/ruslanguns/go-simple-api/internal/models"
)

type ProductService interface {
	ListProducts(ctx context.Context) ([]models.Product, error)
	CreateProduct(ctx context.Context, product models.Product) (models.Product, error)
}

type productService struct {
	// Aquí podrías inyectar dependencias como un repositorio
}

func NewProductService() ProductService {
	return &productService{}
}

func (s *productService) ListProducts(ctx context.Context) ([]models.Product, error) {
	// Implementación simulada
	return []models.Product{
		{ID: "1", Name: "Product 1", Description: "Description 1", Price: 9.99},
		{ID: "2", Name: "Product 2", Description: "Description 2", Price: 19.99},
	}, nil
}

func (s *productService) CreateProduct(ctx context.Context, product models.Product) (models.Product, error) {
	// Implementación simulada
	product.ID = "3" // En una implementación real, esto se generaría o se obtendría de la base de datos
	return product, nil
}
