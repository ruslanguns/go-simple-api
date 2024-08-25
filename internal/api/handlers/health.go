package handlers

import (
	"net/http"

	"github.com/ruslanguns/go-simple-api/pkg/encoding"
	"github.com/ruslanguns/go-simple-api/pkg/logger"
)

func HandleHealthCheck(log *logger.Logger) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if err := encoding.Encode(w, r, http.StatusOK, map[string]string{"status": "OK"}); err != nil {
			handleEncodeError(w, r, log, err)
		}
	}
}
