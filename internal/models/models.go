package models

import (
	"context"
	"time"
)

type User struct {
	ID        string    `json:"id"`
	Name      string    `json:"name"`
	Email     string    `json:"email"`
	CreatedAt time.Time `json:"created_at"`
}

type Product struct {
	ID          string    `json:"id"`
	Name        string    `json:"name"`
	Description string    `json:"description"`
	Price       float64   `json:"price"`
	CreatedAt   time.Time `json:"created_at"`
}

// Validator es una interfaz que pueden implementar los modelos
// para proporcionar validación de datos.
type Validator interface {
	Validate(ctx context.Context) map[string]string
}

// Implementación de Validate para User
func (u *User) Validate(ctx context.Context) map[string]string {
	problems := make(map[string]string)

	if u.Name == "" {
		problems["name"] = "Name is required"
	}

	if u.Email == "" {
		problems["email"] = "Email is required"
	}
	// Aquí podrías añadir más validaciones, como formato de email, etc.

	return problems
}

// Implementación de Validate para Product
func (p *Product) Validate(ctx context.Context) map[string]string {
	problems := make(map[string]string)

	if p.Name == "" {
		problems["name"] = "Name is required"
	}

	if p.Price < 0 {
		problems["price"] = "Price must be non-negative"
	}

	return problems
}
