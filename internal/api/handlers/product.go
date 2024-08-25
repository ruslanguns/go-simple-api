package handlers

import (
	"net/http"

	"github.com/go-chi/chi/v5"

	"github.com/ruslanguns/go-simple-api/internal/models"
	"github.com/ruslanguns/go-simple-api/internal/services"
	"github.com/ruslanguns/go-simple-api/pkg/encoding"
	"github.com/ruslanguns/go-simple-api/pkg/logger"
)

type ProductHandler struct {
	log     *logger.Logger
	service services.ProductService
}

func NewProductHandler(log *logger.Logger, service services.ProductService) *ProductHandler {
	return &ProductHandler{log: log, service: service}
}

func (h *ProductHandler) Routes() chi.Router {
	r := chi.NewRouter()
	r.Get("/", h.ListProducts)
	r.Post("/", h.CreateProduct)
	return r
}

func (h *ProductHandler) ListProducts(w http.ResponseWriter, r *http.Request) {
	products, err := h.service.ListProducts(r.Context())
	if err != nil {
		h.log.Error("Failed to list products: %v", err)
		if err := encoding.Encode(w, r, http.StatusInternalServerError, map[string]string{"error": "Failed to list products"}); err != nil {
			handleEncodeError(w, r, h.log, err)
		}
		return
	}
	if err := encoding.Encode(w, r, http.StatusOK, products); err != nil {
		handleEncodeError(w, r, h.log, err)
	}
}

func (h *ProductHandler) CreateProduct(w http.ResponseWriter, r *http.Request) {
	product, err := encoding.Decode[models.Product](r)
	if err != nil {
		if err := encoding.Encode(w, r, http.StatusBadRequest, map[string]string{"error": "Invalid request body"}); err != nil {
			handleEncodeError(w, r, h.log, err)
		}
		return
	}

	createdProduct, err := h.service.CreateProduct(r.Context(), product)
	if err != nil {
		h.log.Error("Failed to create product: %v", err)
		if err := encoding.Encode(w, r, http.StatusInternalServerError, map[string]string{"error": "Failed to create product"}); err != nil {
			handleEncodeError(w, r, h.log, err)
		}
		return
	}

	if err := encoding.Encode(w, r, http.StatusCreated, createdProduct); err != nil {
		handleEncodeError(w, r, h.log, err)
	}
}
