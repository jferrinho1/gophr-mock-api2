package middleware

import (
	"net/http"

	h "github.com/jferrinho1/gophr-mock-api2/handlers"
)

func HandlePreflightRequests(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		method := r.Method

		if method == http.MethodOptions {
			handler := h.DummyErrorHandler{Status: 200}
			handler.ServeHTTP(w, r)
			
		} else {
			next.ServeHTTP(w, r)
		}
	})
}
