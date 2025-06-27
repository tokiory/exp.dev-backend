package handler

import "net/http"

func HealthHandler(opts HandlerOptions) http.HandlerFunc {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("OK"))
		w.WriteHeader(200)
	})
}
