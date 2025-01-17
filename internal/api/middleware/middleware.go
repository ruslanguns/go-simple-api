package middleware

import (
	"net/http"
	"time"

	"github.com/ruslanguns/go-simple-api/pkg/logger"
)

func Logging(log *logger.Logger) func(http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			start := time.Now()
			next.ServeHTTP(w, r)
			log.Info("%s %s %s", r.Method, r.RequestURI, time.Since(start))
		})
	}
}
