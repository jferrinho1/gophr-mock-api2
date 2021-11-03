package middleware

import (
	h "github.com/jferrinho1/gophr-mock-api2/handlers"
	"net/http"
)

func RequireFrontendVersionHeader(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		version := r.Header.Get("FRONTEND-VERSION")

		if version == "2.0.20200626.1.5" {
			next.ServeHTTP(w, r)
		} else {
			// Write an error and stop the handler chain
			handler := h.DummyHandler{
				FilePath: "./schemas/frontend/frontend_version_error.json",
				Status:   422,
			}

			handler.ServeHTTP(w,r)
		}
	})
}
