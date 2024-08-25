package handlers

import (
	"net/http"

	"github.com/go-chi/chi/v5"

	"github.com/ruslanguns/go-simple-api/internal/models"
	"github.com/ruslanguns/go-simple-api/internal/services"
	"github.com/ruslanguns/go-simple-api/pkg/encoding"
	"github.com/ruslanguns/go-simple-api/pkg/logger"
)

type UserHandler struct {
	log     *logger.Logger
	service services.UserService
}

func NewUserHandler(log *logger.Logger, service services.UserService) *UserHandler {
	return &UserHandler{log: log, service: service}
}

func (h *UserHandler) Routes() chi.Router {
	r := chi.NewRouter()
	r.Get("/", h.ListUsers)
	r.Get("/{id}", h.GetUser)
	r.Post("/", h.CreateUser)
	return r
}

func (h *UserHandler) ListUsers(w http.ResponseWriter, r *http.Request) {
	users, err := h.service.ListUsers(r.Context())
	if err != nil {
		h.log.Error("Failed to list users: %v", err)
		if err := encoding.Encode(w, r, http.StatusInternalServerError, map[string]string{"error": "Failed to list users"}); err != nil {
			handleEncodeError(w, r, h.log, err)
		}
		return
	}
	if err := encoding.Encode(w, r, http.StatusOK, users); err != nil {
		handleEncodeError(w, r, h.log, err)
	}
}

func (h *UserHandler) GetUser(w http.ResponseWriter, r *http.Request) {
	userID := chi.URLParam(r, "id")
	user, err := h.service.GetUser(r.Context(), userID)
	if err != nil {
		h.log.Error("Failed to get user: %v", err)
		if err := encoding.Encode(w, r, http.StatusInternalServerError, map[string]string{"error": "Failed to get user"}); err != nil {
			handleEncodeError(w, r, h.log, err)
		}
		return
	}
	if err := encoding.Encode(w, r, http.StatusOK, user); err != nil {
		handleEncodeError(w, r, h.log, err)
	}
}

func (h *UserHandler) CreateUser(w http.ResponseWriter, r *http.Request) {
	user, err := encoding.Decode[models.User](r)
	if err != nil {
		if err := encoding.Encode(w, r, http.StatusBadRequest, map[string]string{"error": "Invalid request body"}); err != nil {
			handleEncodeError(w, r, h.log, err)
		}
		return
	}

	createdUser, err := h.service.CreateUser(r.Context(), user)
	if err != nil {
		h.log.Error("Failed to create user: %v", err)
		if err := encoding.Encode(w, r, http.StatusInternalServerError, map[string]string{"error": "Failed to create user"}); err != nil {
			handleEncodeError(w, r, h.log, err)
		}
		return
	}

	if err := encoding.Encode(w, r, http.StatusCreated, createdUser); err != nil {
		handleEncodeError(w, r, h.log, err)
	}
}
