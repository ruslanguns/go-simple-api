package handlers

import (
	"net/http"

	"github.com/ruslanguns/go-simple-api/pkg/logger"
)

func handleEncodeError(w http.ResponseWriter, _ *http.Request, log *logger.Logger, err error) {
	log.Error("Failed to encode response: %v", err)
	http.Error(w, "Internal Server Error", http.StatusInternalServerError)
}
