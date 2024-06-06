package health

import "net/http"

func healthCheck(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("OK"))
}

func RegisterRoute(router *http.ServeMux) {
	router.HandleFunc("GET /health", healthCheck)
}
