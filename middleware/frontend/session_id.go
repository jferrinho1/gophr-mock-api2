package middleware

import (
	h "github.com/jferrinho1/gophr-mock-api2/handlers"
	"net/http"
)

func RequireSessionIdHeader(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		sessionId := r.Header.Get("SESSION-ID")

		if sessionId != "" {
			next.ServeHTTP(w, r)
		} else {
			// Write an error and stop the handler chain
			handler := h.DummyHandler{
				FilePath: "./schemas/frontend/session_id_error.json",
				Status:   401,
			}

			handler.ServeHTTP(w, r)
		}
	})
}
